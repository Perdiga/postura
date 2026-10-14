package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/jackchuka/postura/internal/collect"
	"github.com/spf13/cobra"
)

type permissionCheck struct {
	scope      string
	target     string
	capability string
	missing    []string
	guidance   string
	problem    string
}

var (
	permissionsEnterprise string
	permissionsOrg        []string
	permissionsScope      []string
)

var permissionsCmd = &cobra.Command{
	Use:   "permissions",
	Short: "Check GitHub access needed by postura",
}

var permissionsCheckCmd = &cobra.Command{
	Use:   "check [owner/repo ...]",
	Short: "Test API access for posture collection capabilities",
	Long:  "Runs read-only collection probes and reports which settings postura could not read. Missing data can also mean an endpoint or feature is unavailable; policy failures are not permission failures.",
	RunE:  runPermissionsCheck,
}

func init() {
	f := permissionsCheckCmd.Flags()
	f.StringSliceVar(&permissionsOrg, "org", nil, "organization(s) to check")
	f.StringVar(&permissionsEnterprise, "enterprise", "", "enterprise slug to check")
	f.StringSliceVar(&permissionsScope, "scope", nil, "limit checks to enterprise|org|repos (repeatable)")
	permissionsCmd.AddCommand(permissionsCheckCmd)
	rootCmd.AddCommand(permissionsCmd)
}

func runPermissionsCheck(cmd *cobra.Command, args []string) error {
	specs, err := orgsFromArgs(permissionsOrg, args)
	if err != nil {
		return err
	}
	scopes, err := parseScopes(permissionsScope, permissionsEnterprise != "", len(specs) > 0)
	if err != nil {
		return err
	}
	token, err := resolveToken()
	if err != nil {
		return err
	}
	ctx := context.Background()
	client := collect.NewClient(ctx, token)
	checks := make([]permissionCheck, 0)

	if scopes.enterprise {
		facts, err := collect.Enterprise(ctx, client, permissionsEnterprise)
		if err != nil {
			checks = append(checks, permissionCheck{
				scope: "enterprise", target: permissionsEnterprise, capability: "enterprise collection",
				problem: err.Error(), guidance: "Use an enterprise-owner token with read:enterprise access.",
			})
		} else {
			checks = append(checks, enterprisePermissionChecks(facts)...)
		}
	}
	for _, spec := range specs {
		if scopes.org {
			facts, err := collect.Org(ctx, client, spec.name, cmd.ErrOrStderr())
			if err != nil {
				checks = append(checks, permissionCheck{
					scope: "org", target: spec.name, capability: "organization collection",
					problem: err.Error(), guidance: "Check organization membership and grant org read access (classic PAT: admin:org; fine-grained PAT: Organization members read).",
				})
			} else {
				checks = append(checks, orgPermissionChecks(facts)...)
			}
		}
		if scopes.repos {
			facts, err := collect.Repos(ctx, client, spec.name, spec.repos)
			if err != nil {
				checks = append(checks, permissionCheck{
					scope: "repos", target: spec.name, capability: "repository listing",
					problem: err.Error(), guidance: "Grant repository Metadata read access and verify the token can list the selected repositories.",
				})
			} else {
				for _, repo := range facts {
					checks = append(checks, repoPermissionChecks(repo)...)
				}
			}
		}
	}
	return writePermissionChecks(cmd.OutOrStdout(), checks)
}

func enterprisePermissionChecks(facts map[string]any) []permissionCheck {
	target := stringFact(facts, "enterprise")
	checks := []permissionCheck{
		probe("enterprise", target, "enterprise owner settings", facts,
			[]string{"saml_sso_url", "default_repository_permission", "ip_allow_list_enabled", "allow_private_repository_forking", "members_can_change_repository_visibility", "members_can_invite_collaborators", "two_factor_required_setting", "members_can_create_public_repositories", "members_can_delete_repository", "notification_delivery_restriction_enabled", "outside_collaborators", "outside_collaborator_count"},
			"Use an enterprise-owner token with read:enterprise access."),
		probe("enterprise", target, "enterprise GHAS defaults", facts,
			[]string{"advanced_security_enabled_for_new_repositories", "secret_scanning_enabled_for_new_repositories", "secret_scanning_push_protection_enabled_for_new_repositories"},
			"Use an enterprise-owner token with read:enterprise access; GHAS settings may require an eligible enterprise plan."),
	}
	return checks
}

func orgPermissionChecks(facts map[string]any) []permissionCheck {
	target := stringFact(facts, "org")
	return []permissionCheck{
		probe("org", target, "organization settings and membership", facts,
			[]string{"two_factor_requirement_enabled", "default_repository_permission", "members_can_create_repositories", "members_can_create_public_repositories", "members_can_fork_private_repositories", "secret_scanning_enabled_for_new_repositories", "secret_scanning_push_protection_enabled_for_new_repositories", "dependabot_alerts_enabled_for_new_repositories", "admins", "admin_count", "members", "member_count", "outside_collaborators", "outside_collaborator_count", "teams", "invitations"},
			"Classic PAT: admin:org. Fine-grained PAT: Organization members read and organization Administration read where required."),
		probe("org", target, "organization role assignments", facts,
			[]string{"organization_roles"},
			"Grant org role read access (classic PAT: admin:org; fine-grained PAT: Organization custom roles read)."),
		probe("org", target, "organization Actions policy", facts,
			[]string{"actions_enabled_repositories", "actions_allowed_actions", "actions_default_workflow_permissions", "actions_can_approve_pull_request_reviews"},
			"Grant organization Actions/Administration read access (classic PAT: admin:org)."),
		probe("org", target, "organization Actions secret metadata", facts,
			[]string{"oldest_actions_secret_age_days"},
			"Grant Actions secrets read access for the organization (classic PAT: admin:org)."),
		probe("org", target, "self-hosted runner groups", facts,
			[]string{"runner_groups"},
			"Grant organization Self-hosted runners read access (classic PAT: admin:org)."),
	}
}

func repoPermissionChecks(facts map[string]any) []permissionCheck {
	target := stringFact(facts, "name")
	checks := []permissionCheck{
		probe("repo", target, "repository contents and workflows", facts,
			[]string{"repository_files", "unpinned_action_ref_count", "workflow_permission_issue_count", "workflow_pr_trigger_count", "codeql_action_count", "dependency_review_action_count", "cache_action_count", "concurrency_workflow_count", "matrix_job_count", "narrow_trigger_count", "reusable_workflow_count", "release_automation_count", "stale_issue_automation_count", "codeowners", "codeowners_error_count", "dependabot_config", "renovate_config"},
			"Grant repository Contents read access (classic PAT: repo for private repositories)."),
		probe("repo", target, "branch protection and rulesets", facts,
			[]string{"protection", "ruleset_count"},
			"Grant repository Administration read access and ensure the token has access to the default branch."),
		probe("repo", target, "repository collaborators and team grants", facts,
			[]string{"collaborator_count", "admin_collaborator_count", "teams"},
			"Use a repository administrator token; org-wide team grants also require organization team read access."),
		probe("repo", target, "security settings and private reporting", facts,
			[]string{"secret_scanning", "secret_scanning_push_protection", "vulnerability_alerts"},
			"Use repository Administration read access, or organization owner/security-manager access for security_and_analysis settings; private repos also need repo access."),
		probe("repo", target, "repository Actions secrets and environments", facts,
			[]string{"oldest_actions_secret_age_days", "environments"},
			"Grant Actions read access for the repository."),
	}
	if facts["visibility"] == "public" && facts["archived"] != true {
		checks = append(checks, probe("repo", target, "private vulnerability reporting (public repository)", facts,
			[]string{"private_vulnerability_reporting"},
			"For a fine-grained PAT, grant repository Metadata read. For a classic PAT, use public_repo or repo; the token owner must be able to access the repository."))
	}
	if facts["vulnerability_alerts"] == true {
		checks = append(checks, probe("repo", target, "Dependabot alert details", facts,
			[]string{"dependabot_alerts"},
			"Grant Dependabot alerts read access; classic PATs need security_events and suitable repository/org access."))
	}
	return checks
}

func probe(scope, target, capability string, facts map[string]any, required []string, guidance string) permissionCheck {
	check := permissionCheck{scope: scope, target: target, capability: capability, guidance: guidance}
	for _, key := range required {
		if _, ok := facts[key]; !ok {
			check.missing = append(check.missing, key)
		}
	}
	return check
}

func stringFact(facts map[string]any, key string) string {
	if value, ok := facts[key].(string); ok {
		return value
	}
	return "unknown target"
}

func writePermissionChecks(out io.Writer, checks []permissionCheck) error {
	missing := 0
	for _, check := range checks {
		if len(check.missing) == 0 && check.problem == "" {
			_, _ = fmt.Fprintf(out, "[OK] %s %s: %s\n", check.scope, check.target, check.capability)
			continue
		}
		missing++
		_, _ = fmt.Fprintf(out, "[INCOMPLETE] %s %s: %s\n", check.scope, check.target, check.capability)
		if check.problem != "" {
			_, _ = fmt.Fprintf(out, "  collection error: %s\n", check.problem)
		}
		if len(check.missing) > 0 {
			_, _ = fmt.Fprintf(out, "  unreadable facts: %s\n", join(check.missing, ", "))
		}
		_, _ = fmt.Fprintf(out, "  access to try: %s\n", check.guidance)
		_, _ = fmt.Fprintln(out, "  Note: missing facts can also indicate an unsupported or ineligible GitHub feature, not only a token permission.")
	}
	if missing == 0 {
		_, _ = fmt.Fprintf(out, "\nPermission check complete: all %d capability groups were readable.\n", len(checks))
	} else {
		_, _ = fmt.Fprintf(out, "\nPermission check incomplete: %d of %d capability groups had unreadable data. Policy failures are not permission failures.\n", missing, len(checks))
		return fmt.Errorf("permission check incomplete")
	}
	return nil
}

func join(values []string, separator string) string {
	if len(values) == 0 {
		return ""
	}
	out := values[0]
	for _, value := range values[1:] {
		out += separator + value
	}
	return out
}
