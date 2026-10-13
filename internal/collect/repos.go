package collect

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/google/go-github/v84/github"
)

// repoWorkers bounds how many repos are collected concurrently. Each repo makes
// several API calls, so a small pool cuts wall-clock on large orgs while keeping
// secondary-rate-limit pressure modest (the client also retries on a throttle).
const repoWorkers = 8

// Committed-file paths whose presence drives repo facts. All are shallow, so a
// single recursive tree of the default branch (see repoFiles) resolves every one
// in one API call. The deprecated package.json "renovate" key is intentionally
// not detected.
var (
	codeownersFilePaths = []string{".github/CODEOWNERS", "CODEOWNERS", "docs/CODEOWNERS"}
	dependabotFilePaths = []string{".github/dependabot.yml", ".github/dependabot.yaml"}
	renovateFilePaths   = []string{
		"renovate.json", "renovate.json5",
		".renovaterc", ".renovaterc.json", ".renovaterc.json5",
		".github/renovate.json", ".github/renovate.json5",
	}
)

// orgWide holds the results of the sweeps read once for the whole org rather
// than per repo. Each result carries an OK flag that is false when the listing
// could not be completed, so the repo leaves the corresponding fact absent
// (unknown) instead of recording a false empty.
type orgWide struct {
	grants   map[string][]any
	grantsOK bool
	alerts   map[string]any
	alertsOK bool
}

// Repos collects per-repo facts. With no names it audits every non-forked repo
// in the org. Each returned object is fully normalized against RepoVars.
func Repos(ctx context.Context, c *github.Client, org string, names []string) ([]map[string]any, error) {
	repos, err := resolveRepos(ctx, c, org, names)
	if err != nil {
		return nil, err
	}
	// The org-wide sweeps: direct per-repo team grants, read from the team side
	// (org-role access does not appear there), and open Dependabot alerts. Both
	// cover every repo in one paginated listing rather than a call per repo.
	shared := orgWide{}
	shared.grants, shared.grantsOK = directRepoGrants(ctx, c, org)
	// The alert sweep is worth its pages only when auditing the whole org: for an
	// explicit subset it would page the org's entire alert set to use a handful of
	// it, so each repo reads its own instead. Repos falls back per repo whenever
	// the sweep is missing, which also covers a token that cannot read the org
	// endpoint but can read the repo one.
	if len(names) == 0 {
		shared.alerts, shared.alertsOK = dependabotAlerts(ctx, c, org)
	}
	// Collect each repo's facts concurrently, writing into its own slot so the
	// output order matches the (already deterministic) repo list.
	out := make([]map[string]any, len(repos))
	sem := make(chan struct{}, repoWorkers)
	var wg sync.WaitGroup
	for i, r := range repos {
		wg.Add(1)
		sem <- struct{}{}
		go func(i int, r *github.Repository) {
			defer wg.Done()
			defer func() { <-sem }()
			out[i] = repoFacts(ctx, c, org, r, shared)
		}(i, r)
	}
	wg.Wait()
	return out, nil
}

func resolveRepos(ctx context.Context, c *github.Client, org string, names []string) ([]*github.Repository, error) {
	if len(names) > 0 {
		var out []*github.Repository
		for _, n := range names {
			r, _, err := c.Repositories.Get(ctx, org, n)
			if err != nil {
				return nil, err
			}
			out = append(out, r)
		}
		return out, nil
	}
	opt := &github.RepositoryListByOrgOptions{Type: "all", ListOptions: github.ListOptions{PerPage: 100}}
	var out []*github.Repository
	for {
		repos, resp, err := c.Repositories.ListByOrg(ctx, org, opt)
		if err != nil {
			return nil, err
		}
		for _, r := range repos {
			if r.GetFork() {
				continue
			}
			out = append(out, r)
		}
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return out, nil
}

// RepositoryNames returns all non-forked repository names in an organization.
// It is intentionally metadata-only for commands that need target names but not
// repository posture facts.
func RepositoryNames(ctx context.Context, c *github.Client, org string) ([]string, error) {
	repos, err := resolveRepos(ctx, c, org, nil)
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(repos))
	for _, repo := range repos {
		names = append(names, repo.GetName())
	}
	sort.Strings(names)
	return names, nil
}

func repoFacts(ctx context.Context, c *github.Client, org string, r *github.Repository, shared orgWide) map[string]any {
	name := r.GetName()
	f := map[string]any{
		"name":       org + "/" + name,
		"owner":      org,
		"archived":   r.GetArchived(),
		"visibility": r.GetVisibility(),
	}
	setBool(f, "allow_forking", r.AllowForking)
	if pushedAt := r.GetPushedAt().Time; !pushedAt.IsZero() {
		f["last_pushed_days"] = ageDays(time.Now(), pushedAt)
	}

	// Code security & analysis: readable on public repos / with admin. When the
	// block is absent we leave the keys out so REPO-3/REPO-9 report unknown rather
	// than a false "disabled".
	if sa := r.GetSecurityAndAnalysis(); sa != nil {
		if ss := sa.GetSecretScanning(); ss != nil {
			f["secret_scanning"] = ss.GetStatus()
		}
		if pp := sa.GetSecretScanningPushProtection(); pp != nil {
			f["secret_scanning_push_protection"] = pp.GetStatus()
		}
	}

	// License: SPDX id or null.
	f["license"] = nil
	if lic := r.GetLicense(); lic != nil {
		f["license"] = strOrNull(lic.GetSPDXID())
	}

	// Dependabot / vulnerability alerts. Left absent (→ unknown) when unreadable
	// (missing scope / no admin), never defaulted to a value that would read as a
	// false "disabled" for a rule over this fact.
	if enabled, _, err := c.Repositories.GetVulnerabilityAlerts(ctx, org, name); err == nil {
		f["vulnerability_alerts"] = enabled
	}
	if enabled, _, err := c.Repositories.IsPrivateReportingEnabled(ctx, org, name); err == nil {
		f["private_vulnerability_reporting"] = enabled
	}
	if age, ok := oldestActionsSecretAge(ctx, c, "repos/"+org+"/"+name+"/actions/secrets"); ok {
		f["oldest_actions_secret_age_days"] = age
	}
	if collaborators, admins, ok := repoCollaboratorAdminCounts(ctx, c, org, name); ok {
		f["collaborator_count"] = collaborators
		f["admin_collaborator_count"] = admins
	}
	if environments, ok := repoEnvironments(ctx, c, org, name); ok {
		f["environments"] = environments
	}

	// CODEOWNERS plus automated dependency-update config (version updates), all
	// resolved from one default-branch tree read. Each fact is omitted on
	// uncertainty so the rule reports unknown rather than a false "absent".
	files := repoFiles(ctx, c, org, name, r.GetDefaultBranch())
	if files.codeownersKnown {
		f["codeowners"] = files.codeowners
	}
	if n, ok := codeownersErrorCount(ctx, c, org, name, files); ok {
		f["codeowners_error_count"] = n
	}
	if files.dependabotKnown {
		f["dependabot_config"] = files.dependabot
	}
	if files.renovateKnown {
		f["renovate_config"] = files.renovate
	}
	if files.repositoryFilesKnown {
		f["repository_files"] = files.repositoryFiles
	}
	if counts, ok := workflowSecurityCountsForRepo(ctx, c, org, name, files); ok {
		f["unpinned_action_ref_count"] = counts.unpinnedActionRefCount
		f["workflow_permission_issue_count"] = counts.permissionIssueCount
		f["workflow_pr_trigger_count"] = counts.pullRequestTriggerCount
		f["codeql_action_count"] = counts.codeQLActionCount
		f["dependency_review_action_count"] = counts.dependencyReviewCount
		f["cache_action_count"] = counts.cacheActionCount
		f["concurrency_workflow_count"] = counts.concurrencyWorkflowCount
		f["matrix_job_count"] = counts.matrixJobCount
		f["narrow_trigger_count"] = counts.narrowTriggerCount
		f["reusable_workflow_count"] = counts.reusableWorkflowCount
		f["release_automation_count"] = counts.releaseAutomationCount
		f["stale_issue_automation_count"] = counts.staleIssueAutomationCount
	}

	// Teams holding a direct per-repo grant on this repo. A repo absent from the
	// org's grant map definitively has none (empty list). Left absent (unknown)
	// only when the org grants could not be enumerated, so a rule over `teams`
	// never sees a false empty.
	if shared.grantsOK {
		if g, ok := shared.grants[strings.ToLower(org+"/"+name)]; ok {
			f["teams"] = g
		} else {
			f["teams"] = []any{}
		}
	}

	// Open Dependabot alerts, folded into per-severity counts and ages by the
	// org-wide sweep when there was one, else read for this repo alone. Absent
	// when neither listing could be completed, so a repo never reports a clean
	// "no alerts" that only means nobody was allowed to look. A repo known to
	// have alerts disabled is skipped outright: the sweep never mentions it (a
	// zero-fill would be a false clean, and would differ from a repo-subset
	// run), and its own listing is a guaranteed 403 not worth the request.
	if f["vulnerability_alerts"] != false {
		if shared.alertsOK {
			f["dependabot_alerts"] = alertsFor(shared.alerts, org, name)
		} else if a, ok := repoAlerts(ctx, c, org, name); ok {
			f["dependabot_alerts"] = a
		}
	}

	// Branch protection and rulesets, both omitted together when either read was
	// refused: an unreadable branch is not an unprotected one.
	if prot, rulesetCount, ok := protection(ctx, c, org, name, r.GetDefaultBranch()); ok {
		f["protection"] = prot
		f["ruleset_count"] = rulesetCount
	}

	return f
}

func repoCollaboratorAdminCounts(ctx context.Context, c *github.Client, org, repo string) (int, int, bool) {
	collaborators, admins := 0, 0
	page := 1
	for {
		var body []struct {
			Permissions struct {
				Admin *bool `json:"admin"`
			} `json:"permissions"`
		}
		resp, err := getJSON(ctx, c, "repos/"+org+"/"+repo+"/collaborators?affiliation=all&per_page=100&page="+itoa(int64(page)), &body)
		if err != nil {
			return 0, 0, false
		}
		for _, collaborator := range body {
			if collaborator.Permissions.Admin == nil {
				return 0, 0, false
			}
			collaborators++
			if *collaborator.Permissions.Admin {
				admins++
			}
		}
		if resp == nil || resp.NextPage == 0 {
			return collaborators, admins, true
		}
		if resp.NextPage == page {
			return 0, 0, false
		}
		page = resp.NextPage
	}
}

// fileFacts holds the presence of committed config files, each paired with a
// Known flag that is false when absence could not be proven (the caller then
// omits the fact so the rule reports unknown, never a false "absent").
type fileFacts struct {
	codeowners, codeownersKnown bool
	dependabot, dependabotKnown bool
	renovate, renovateKnown     bool
	workflowFiles               map[string]string
	workflowFilesKnown          bool
	repositoryFiles             map[string]bool
	repositoryFilesKnown        bool
}

// repoFiles resolves every committed-file fact from a single recursive tree of
// the default branch, instead of probing each path with its own Contents call.
//
// Absence semantics: a matched blob proves presence (true, known). A complete
// (non-truncated) tree with no match proves absence (false, known). An empty
// repository — no default branch, or the tree endpoint's 404/409 — is likewise a
// definitive absence. Every other error, and a truncated tree in which a path was
// not seen, leaves absence unproven (Known=false → unknown). The tree caps at
// 100k entries; all the paths here are shallow, so truncation only bites
// pathologically large repos.
func repoFiles(ctx context.Context, c *github.Client, org, repo, branch string) fileFacts {
	if branch == "" {
		return allAbsent() // no default branch => no commits => nothing present
	}
	// go-github places the ref unescaped in the URL path, so a default branch
	// whose name contains a slash (e.g. "release/v1") won't resolve; every file
	// fact for that repo then degrades to unknown. Rare, and it fails safe.
	tree, resp, err := c.Git.GetTree(ctx, org, repo, branch, true)
	if err != nil {
		// An empty repository has no tree (404, or 409 "Git Repository is empty"):
		// definitive absence. Any other error leaves every fact unknown.
		if resp != nil && (resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusConflict) {
			return allAbsent()
		}
		return fileFacts{}
	}

	present := make(map[string]bool, len(tree.Entries))
	workflowFiles := map[string]string{}
	for _, e := range tree.Entries {
		if e.GetType() == "blob" {
			present[e.GetPath()] = true
			if isWorkflowPath(e.GetPath()) {
				workflowFiles[e.GetPath()] = e.GetSHA()
			}
		}
	}

	truncated := tree.GetTruncated()
	// fact returns (present, known): a hit is always trustworthy; a miss is only
	// trustworthy when the tree was complete.
	fact := func(paths []string) (bool, bool) {
		for _, p := range paths {
			if present[p] {
				return true, true
			}
		}
		return false, !truncated
	}

	var f fileFacts
	f.codeowners, f.codeownersKnown = fact(codeownersFilePaths)
	f.dependabot, f.dependabotKnown = fact(dependabotFilePaths)
	f.renovate, f.renovateKnown = fact(renovateFilePaths)
	if !truncated {
		f.workflowFilesKnown = true
		f.workflowFiles = workflowFiles
		for _, sha := range workflowFiles {
			if sha == "" {
				f.workflowFilesKnown = false
				break
			}
		}
		f.repositoryFilesKnown = true
		f.repositoryFiles = repositoryFileFacts(present)
	}
	return f
}

// allAbsent is the definitive "no files present" result (every fact known-false).
func allAbsent() fileFacts {
	return fileFacts{
		codeownersKnown: true, dependabotKnown: true, renovateKnown: true,
		workflowFiles: map[string]string{}, workflowFilesKnown: true,
		repositoryFiles: repositoryFileFacts(map[string]bool{}), repositoryFilesKnown: true,
	}
}

func repositoryFileFacts(present map[string]bool) map[string]bool {
	has := func(paths ...string) bool {
		for _, path := range paths {
			if present[path] {
				return true
			}
		}
		return false
	}
	copilotInstructions := has(".github/copilot-instructions.md")
	for path := range present {
		lower := strings.ToLower(path)
		if strings.HasPrefix(lower, ".github/instructions/") && strings.HasSuffix(lower, ".instructions.md") {
			copilotInstructions = true
			break
		}
	}
	hasTemplate := func(prefix string, extensions ...string) bool {
		for path := range present {
			if !strings.HasPrefix(path, prefix) {
				continue
			}
			lower := strings.ToLower(path)
			for _, extension := range extensions {
				if strings.HasSuffix(lower, extension) {
					return true
				}
			}
		}
		return false
	}
	return map[string]bool{
		"readme":                has("README.md", "README", "readme.md"),
		"security_policy":       has("SECURITY.md", ".github/SECURITY.md", "docs/SECURITY.md"),
		"contributing_guide":    has("CONTRIBUTING.md", ".github/CONTRIBUTING.md", "docs/CONTRIBUTING.md"),
		"pull_request_template": has("PULL_REQUEST_TEMPLATE.md", ".github/PULL_REQUEST_TEMPLATE.md") || hasTemplate(".github/PULL_REQUEST_TEMPLATE/", ".md"),
		"issue_templates":       has(".github/ISSUE_TEMPLATE.yml", ".github/ISSUE_TEMPLATE.yaml") || hasTemplate(".github/ISSUE_TEMPLATE/", ".md", ".yml", ".yaml"),
		"copilot_instructions":  copilotInstructions,
		"sast_configuration":    has(".semgrep.yml", ".semgrep.yaml", "semgrep.yml", "semgrep.yaml", "sonar-project.properties", "codeql-config.yml", ".github/codeql/codeql-config.yml"),
	}
}

func isWorkflowPath(path string) bool {
	if !strings.HasPrefix(path, ".github/workflows/") || strings.Count(path, "/") != 2 {
		return false
	}
	ext := strings.ToLower(path[strings.LastIndexByte(path, '.')+1:])
	return ext == "yml" || ext == "yaml"
}

type workflowSecurityCounts struct {
	unpinnedActionRefCount    int
	permissionIssueCount      int
	pullRequestTriggerCount   int
	codeQLActionCount         int
	dependencyReviewCount     int
	cacheActionCount          int
	concurrencyWorkflowCount  int
	matrixJobCount            int
	narrowTriggerCount        int
	reusableWorkflowCount     int
	releaseAutomationCount    int
	staleIssueAutomationCount int
}

func workflowSecurityCountsForRepo(ctx context.Context, c *github.Client, org, repo string, files fileFacts) (workflowSecurityCounts, bool) {
	if !files.workflowFilesKnown {
		return workflowSecurityCounts{}, false
	}
	paths := make([]string, 0, len(files.workflowFiles))
	for path := range files.workflowFiles {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	counts := workflowSecurityCounts{}
	for _, path := range paths {
		sha := files.workflowFiles[path]
		var blob struct {
			Content  string `json:"content"`
			Encoding string `json:"encoding"`
		}
		if _, err := getJSON(ctx, c, "repos/"+org+"/"+repo+"/git/blobs/"+sha, &blob); err != nil || blob.Encoding != "base64" {
			return workflowSecurityCounts{}, false
		}
		data, err := base64.StdEncoding.DecodeString(blob.Content)
		if err != nil {
			return workflowSecurityCounts{}, false
		}
		refs, err := unpinnedActionRefs(data)
		if err != nil {
			return workflowSecurityCounts{}, false
		}
		issues, err := workflowPermissionIssues(data)
		if err != nil {
			return workflowSecurityCounts{}, false
		}
		static, err := inspectWorkflow(data)
		if err != nil {
			return workflowSecurityCounts{}, false
		}
		counts.unpinnedActionRefCount += refs
		counts.permissionIssueCount += issues
		if static.pullRequestTrigger {
			counts.pullRequestTriggerCount++
		}
		counts.codeQLActionCount += static.codeQLActionCount
		counts.dependencyReviewCount += static.dependencyReviewCount
		counts.cacheActionCount += static.cacheActionCount
		if static.concurrency {
			counts.concurrencyWorkflowCount++
		}
		counts.matrixJobCount += static.matrixJobs
		if static.narrowTrigger {
			counts.narrowTriggerCount++
		}
		counts.reusableWorkflowCount += static.reusableWorkflowCount
		counts.releaseAutomationCount += static.releaseAutomationCount
		counts.staleIssueAutomationCount += static.staleIssueAutomationCount
	}
	return counts, true
}

// codeownersErrorCount asks GitHub's own CODEOWNERS linter how many syntax and
// unknown-owner errors the default branch's file has, returning the count and
// whether it is trustworthy. An unknown owner is the failure that matters: GitHub
// silently ignores the bad line, so a repo requiring code-owner review can end up
// enforcing nothing on those paths.
//
// The call is only worth making for a repo proven to have the file: a proven
// absence has vacuously zero errors, and an unproven presence cannot be scored at
// all. Any non-200 leaves the count unknown rather than a false clean bill —
// including the 404 GitHub returns when it sees no CODEOWNERS, which contradicts
// the tree read that got us here.
func codeownersErrorCount(ctx context.Context, c *github.Client, org, repo string, files fileFacts) (int, bool) {
	if !files.codeownersKnown {
		return 0, false
	}
	if !files.codeowners {
		return 0, true
	}
	errs, _, err := c.Repositories.GetCodeownersErrors(ctx, org, repo, nil)
	if err != nil {
		return 0, false
	}
	return len(errs.Errors), true
}

// directRepoGrants maps "owner/repo" to the teams that hold a direct
// (github_team_repository) grant on it, each with that grant's permission. It is
// built from the team->repos endpoint, which lists only explicit per-repo grants
// — access a team receives through an organization role is not returned there
// (that lives in organization_roles). ok is false if any team's grants could not
// be read, so the caller leaves the teams fact absent rather than risk a false
// "no direct grant".
func directRepoGrants(ctx context.Context, c *github.Client, org string) (grants map[string][]any, ok bool) {
	slugs, err := orgTeamSlugs(ctx, c, org)
	if err != nil {
		return nil, false
	}
	grants = map[string][]any{}
	for _, slug := range slugs {
		opt := &github.ListOptions{PerPage: 100}
		for {
			repos, resp, err := c.Teams.ListTeamReposBySlug(ctx, org, slug, opt)
			if err != nil {
				return nil, false
			}
			for _, r := range repos {
				// Key on a lowercased "owner/repo": the lookup side derives the key
				// from the user-supplied --org string, whose casing may differ from
				// GitHub's canonical full name. A mismatch would yield a false empty
				// teams list (a silent REPO-11 pass), so normalize both sides.
				full := strings.ToLower(r.GetFullName())
				grants[full] = append(grants[full], map[string]any{
					"slug":       slug,
					"permission": highestPerm(r.GetPermissions()),
				})
			}
			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}
	}
	return grants, true
}

func orgTeamSlugs(ctx context.Context, c *github.Client, org string) ([]string, error) {
	opt := &github.ListOptions{PerPage: 100}
	var out []string
	for {
		teams, resp, err := c.Teams.ListTeams(ctx, org, opt)
		if err != nil {
			return nil, err
		}
		for _, t := range teams {
			out = append(out, t.GetSlug())
		}
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return out, nil
}

// highestPerm collapses GitHub's cumulative permission set to the single role
// name used elsewhere (admin > maintain > push > triage > pull), matching the
// legacy `permission` vocabulary REPO-11 reads.
func highestPerm(p *github.RepositoryPermissions) string {
	switch {
	case p.GetAdmin():
		return "admin"
	case p.GetMaintain():
		return "maintain"
	case p.GetPush():
		return "push"
	case p.GetTriage():
		return "triage"
	default:
		return "pull"
	}
}

// protection folds classic branch protection and repository rulesets into one
// normalized object. Many orgs protect the default branch with rulesets,
// which the classic protection endpoint does not return, so both are read and
// the stricter outcome wins for each field. Returns the protection facts, the
// count of active default-branch-targeting rulesets, and whether both surfaces
// were actually readable.
//
// The zero value here says "nothing is required", which is exactly what a wide-open
// branch looks like — so a read the token isn't allowed to make must not fall
// through to it. A token without admin on the repo would otherwise report every
// protection rule as a confident failure across the whole org.
func protection(ctx context.Context, c *github.Client, org, repo, branch string) (map[string]any, int, bool) {
	p := map[string]any{
		"required_pull_request_reviews":   false,
		"required_approving_review_count": 0,
		"require_code_owner_reviews":      false,
		"required_status_checks":          false,
		"branches_up_to_date":             false,
		"dismiss_stale_reviews":           false,
		"require_conversation_resolution": false,
		"require_signed_commits":          false,
		"require_linear_history":          false,
		"allow_force_pushes":              true,
		"allow_deletions":                 true,
		"restrict_pushes":                 false,
		"restrict_dismissals":             false,
		"pr_bypass_actor_types":           []any{},
		"ruleset_bypass_actor_types":      []any{},
	}
	// An empty repository has no default branch, so there is nothing to protect.
	if branch == "" {
		return p, 0, true
	}

	// Classic branch protection. The endpoint answers 404 both for a branch that
	// has none and for a token without admin on the repo — go-github tells the two
	// apart by the response message and normalizes the first to a sentinel. Only
	// that one is proof; every other error means the read failed.
	bp, _, err := c.Repositories.GetBranchProtection(ctx, org, repo, branch)
	switch {
	case err == nil:
		p["allow_force_pushes"] = false
		p["allow_deletions"] = false
		if rev := bp.GetRequiredPullRequestReviews(); rev != nil {
			p["required_pull_request_reviews"] = true
			p["restrict_pushes"] = true
			p["required_approving_review_count"] = rev.RequiredApprovingReviewCount
			p["require_code_owner_reviews"] = rev.RequireCodeOwnerReviews
			p["dismiss_stale_reviews"] = rev.DismissStaleReviews
			p["restrict_dismissals"] = rev.DismissalRestrictions != nil
			if allowances := rev.BypassPullRequestAllowances; allowances != nil {
				actors := map[string]bool{}
				if len(allowances.Users) > 0 {
					actors["User"] = true
				}
				if len(allowances.Teams) > 0 {
					actors["Team"] = true
				}
				if len(allowances.Apps) > 0 {
					actors["Integration"] = true
				}
				p["pr_bypass_actor_types"] = sortedAnyKeys(actors)
			}
		}
		if checks := bp.RequiredStatusChecks; checks != nil {
			p["required_status_checks"] = (checks.Contexts != nil && len(*checks.Contexts) > 0) || (checks.Checks != nil && len(*checks.Checks) > 0)
			p["branches_up_to_date"] = checks.Strict
		}
		if setting := bp.RequiredConversationResolution; setting != nil {
			p["require_conversation_resolution"] = setting.Enabled
		}
		if setting := bp.RequiredSignatures; setting != nil {
			p["require_signed_commits"] = setting.Enabled
		}
		if setting := bp.RequireLinearHistory; setting != nil {
			p["require_linear_history"] = setting.Enabled
		}
		if setting := bp.AllowForcePushes; setting != nil {
			p["allow_force_pushes"] = setting.Enabled
		}
		if setting := bp.AllowDeletions; setting != nil {
			p["allow_deletions"] = setting.Enabled
		}
		if bp.Restrictions != nil {
			p["restrict_pushes"] = true
		}
	case errors.Is(err, github.ErrBranchNotProtected):
	default:
		return nil, 0, false
	}

	count, ok := foldRulesets(ctx, c, org, repo, branch, p)
	if !ok {
		return nil, 0, false
	}
	return p, count, true
}

type rulesetSummary struct {
	ID          int64  `json:"id"`
	Target      string `json:"target"`
	Enforcement string `json:"enforcement"`
}

type rulesetDetail struct {
	Conditions struct {
		RefName struct {
			Include []string `json:"include"`
		} `json:"ref_name"`
	} `json:"conditions"`
	Rules []struct {
		Type       string `json:"type"`
		Parameters struct {
			RequiredApprovingReviewCount   int  `json:"required_approving_review_count"`
			RequireCodeOwnerReview         bool `json:"require_code_owner_review"`
			DismissStaleReviewsOnPush      bool `json:"dismiss_stale_reviews_on_push"`
			RequiredReviewThreadResolution bool `json:"required_review_thread_resolution"`
			StrictRequiredStatusChecks     bool `json:"strict_required_status_checks_policy"`
			RequiredStatusChecks           []struct {
				Context string `json:"context"`
			} `json:"required_status_checks"`
		} `json:"parameters"`
	} `json:"rules"`
	BypassActors []struct {
		ActorType string `json:"actor_type"`
	} `json:"bypass_actors"`
}

// foldRulesets reads repository rulesets via the stable REST endpoints (go-github's
// typed ruleset model churns across versions), folds any pull_request rule that
// targets the default branch into the protection facts (stricter wins), and
// collects bypass-actor types from PR-requiring rulesets for REPO-12. Returns the
// number of active, default-branch-targeting rulesets and whether every ruleset
// was readable — a repo with none answers 200 with an empty list, so any error
// is a read failure, and skipping one could miss the rule that protects the branch.
func foldRulesets(ctx context.Context, c *github.Client, org, repo, branch string, p map[string]any) (int, bool) {
	var summaries []rulesetSummary
	if _, err := getJSON(ctx, c, "repos/"+org+"/"+repo+"/rulesets?includes_parents=true", &summaries); err != nil {
		return 0, false
	}
	bypass := map[string]bool{}
	prBypass := map[string]bool{}
	if actors, ok := p["pr_bypass_actor_types"].([]any); ok {
		for _, actor := range actors {
			if actorType, ok := actor.(string); ok {
				prBypass[actorType] = true
			}
		}
	}
	count := 0
	for _, s := range summaries {
		if s.Enforcement != "active" || s.Target != "branch" {
			continue
		}
		var d rulesetDetail
		if _, err := getJSON(ctx, c, "repos/"+org+"/"+repo+"/rulesets/"+itoa(s.ID), &d); err != nil {
			return 0, false
		}
		if !targetsBranch(d.Conditions.RefName.Include, branch) {
			continue
		}
		count++
		for _, a := range d.BypassActors {
			bypass[a.ActorType] = true
		}
		for _, r := range d.Rules {
			switch r.Type {
			case "pull_request":
				p["required_pull_request_reviews"] = true
				p["restrict_pushes"] = true
				if r.Parameters.RequiredApprovingReviewCount > toInt(p["required_approving_review_count"]) {
					p["required_approving_review_count"] = r.Parameters.RequiredApprovingReviewCount
				}
				if r.Parameters.RequireCodeOwnerReview {
					p["require_code_owner_reviews"] = true
				}
				if r.Parameters.DismissStaleReviewsOnPush {
					p["dismiss_stale_reviews"] = true
				}
				if r.Parameters.RequiredReviewThreadResolution {
					p["require_conversation_resolution"] = true
				}
				for _, a := range d.BypassActors {
					prBypass[a.ActorType] = true
				}
			case "required_status_checks":
				if len(r.Parameters.RequiredStatusChecks) > 0 {
					p["required_status_checks"] = true
				}
				if r.Parameters.StrictRequiredStatusChecks {
					p["branches_up_to_date"] = true
				}
			case "required_signatures":
				p["require_signed_commits"] = true
			case "required_linear_history":
				p["require_linear_history"] = true
			case "non_fast_forward":
				p["allow_force_pushes"] = false
			case "deletion":
				p["allow_deletions"] = false
			}
		}
	}
	p["ruleset_bypass_actor_types"] = sortedAnyKeys(bypass)
	p["pr_bypass_actor_types"] = sortedAnyKeys(prBypass)
	return count, true
}

func sortedAnyKeys(values map[string]bool) []any {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	out := make([]any, len(keys))
	for i, key := range keys {
		out[i] = key
	}
	return out
}

// targetsBranch reports whether a ruleset's ref_name include list applies to the
// default branch (~DEFAULT_BRANCH / ~ALL wildcards or an explicit ref).
func targetsBranch(include []string, branch string) bool {
	for _, ref := range include {
		switch ref {
		case "~DEFAULT_BRANCH", "~ALL", "refs/heads/" + branch:
			return true
		}
	}
	return false
}

func toInt(v any) int {
	if n, ok := v.(int); ok {
		return n
	}
	return 0
}
