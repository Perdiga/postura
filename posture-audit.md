# GitHub audit

_1359 fail, 121 unknown across org + 41 repos_

### ORG-1 — Two-factor authentication is enforced org-wide (error)

✅ all 1 pass

### ORG-21 — Workflows cannot approve pull requests (error)

✅ all 1 pass

### ORG-23 — Runner groups cannot run workflows from public repositories (error)

✅ all 1 pass

### REPO-1 — Default branch requires a pull request before merging (error)

_37 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | protection not readable (missing scope?) |

**Fix:** Repo Settings -> Branches/Rules -> require a pull request before merging on the default branch (add push bots to bypass_actors as actor_type=Integration)

### REPO-15 — Required code-owner review is not silently voided by a broken CODEOWNERS (error)

_0 fail, 1 unknown of 2 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | protection not readable (missing scope?) |

**Fix:** Code-owner review is required but GitHub cannot resolve every CODEOWNERS entry (or there is no CODEOWNERS at all); the paths those lines should own are merging without owner approval. Fix the flagged owners, add the file, or drop the requirement.

### REPO-18 — Workflow jobs avoid implicit or broad write GITHUB_TOKEN permissions (error)

_21 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ⚠️ unknown | workflow_permission_issue_count not readable (missing scope?) |

**Fix:** Set workflow- or job-level permissions explicitly; grant only read/none scopes and add write permissions only to the specific job that needs them

### REPO-19 — Default branch requires at least one status check (error)

_39 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | protection not readable (missing scope?) |

**Fix:** Branch protection / ruleset -> require at least one status check before merging

### REPO-2 — At least one approving review is required (error)

_40 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | protection not readable (missing scope?) |

**Fix:** Branch protection / ruleset -> require approvals (>= 1) on the default branch

### REPO-3 — Secret scanning is enabled (public repos) (error)

✅ all 2 pass

### REPO-9 — Push protection is enabled (public repos) (error)

✅ all 2 pass

### ORG-10 — Organization-role grants of broad repo access are intentional and minimal (warn)

✅ all 1 pass

### ORG-11 — Disaster recovery and backup procedures are tested (warn)

_0 fail, 1 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ⚠️ unknown | assessment not established by the collector |

**Fix:** Document backup coverage, recovery objectives, and a recent successful restore exercise

### ORG-14 — CI/CD performance is monitored and reviewed (warn)

_0 fail, 1 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ⚠️ unknown | assessment not established by the collector |

**Fix:** Record CI/CD performance signals, ownership, and review cadence

### ORG-16 — Runner isolation and exposure are reviewed (warn)

_0 fail, 1 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ⚠️ unknown | assessment not established by the collector |

**Fix:** Record runner ownership, isolation, network boundaries, and protections against untrusted workflows

### ORG-18 — GitHub Actions are restricted to selected repositories (warn)

_1 fail, 0 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ❌ fail |  |

**Fix:** Org Settings -> Actions -> General -> Repository access: select only repositories that need Actions

### ORG-19 — GitHub Actions are restricted to trusted actions (warn)

_1 fail, 0 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ❌ fail |  |

**Fix:** Org Settings -> Actions -> General -> Policies: allow GitHub/verified actions and explicitly approved actions only

### ORG-2 — Owner/admin count is proportionate to org size (warn)

_1 fail, 0 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ❌ fail |  |

**Fix:** Org Settings -> People: review who holds Owner; downgrade to Member where possible

### ORG-20 — Default workflow token permission is read-only (warn)

✅ all 1 pass

### ORG-22 — Organization Actions secrets are rotated within policy (warn)

✅ all 1 pass

### ORG-24 — Runner groups are limited to selected repositories (warn)

_1 fail, 0 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ❌ fail |  |

**Fix:** Org Settings -> Actions -> Runner groups: grant access only to repositories that require each group

### ORG-25 — Cloud deployments use OIDC instead of long-lived credentials (warn)

_0 fail, 1 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ⚠️ unknown | assessment not established by the collector |

**Fix:** Document cloud identity federation coverage and remove long-lived deployment credentials where OIDC is supported

### ORG-26 — Members cannot create public repositories (warn)

_1 fail, 0 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ❌ fail |  |

**Fix:** Org Settings -> Member privileges -> Repository creation: disable public repository creation for members

### ORG-27 — New repositories get GitHub Advanced Security by default (warn)

_1 fail, 0 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ❌ fail |  |

**Fix:** Org Settings -> Code security -> automatically enable Advanced Security for new repositories

### ORG-28 — Stale members and administrators are reviewed and removed (warn)

_0 fail, 1 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ⚠️ unknown | assessment not established by the collector |

**Fix:** Review inactive member/admin access on a regular cadence and remove or downgrade accounts no longer needed

### ORG-29 — Organization SSO is configured and enforced (warn)

_0 fail, 1 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ⚠️ unknown | assessment not established by the collector |

**Fix:** Record the organization SAML/SSO configuration and enforcement status

### ORG-3 — Outside collaborators are known and current (warn)

_1 fail, 0 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ❌ fail |  |

**Fix:** Org Settings -> Outside collaborators: remove anyone no longer active, or allowlist expected logins in --config

### ORG-4 — Default member repository permission is not write/admin (warn)

✅ all 1 pass

### ORG-6 — Members cannot create repositories (warn)

_1 fail, 0 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ❌ fail |  |

**Fix:** Org Settings -> Member privileges -> Repository creation: disable member repo creation

### ORG-7 — New repositories get secret scanning, push protection, and Dependabot by default (warn)

_1 fail, 0 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ❌ fail |  |

**Fix:** Org Settings -> Code security -> enable "Automatically enable for new repositories" for Secret scanning, Push protection, and Dependabot alerts

### ORG-8 — Members cannot fork private/internal repositories (warn)

✅ all 1 pass

### REPO-11 — No team holds admin/maintain via a direct per-repo grant (push cap) (warn)

_10 fail, 0 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |

**Fix:** Cap direct (github_team_repository) team grants at 'push'; downgrade any team holding admin/maintain. Org-role grants are audited by ORG-10.

### REPO-12 — PR-bypass is restricted to automation (bots/Apps) (warn)

_1 fail, 1 unknown of 4 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | protection not readable (missing scope?) |

**Fix:** On any ruleset that requires a PR, keep bypass to bot/App actors only (actor_type Integration); remove Team/RepositoryRole/OrganizationAdmin/user bypass.

### REPO-14 — CODEOWNERS file has no unknown owners or syntax errors (warn)

_1 fail, 0 unknown of 2 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |

**Fix:** Open the CODEOWNERS file on GitHub to see the flagged lines; an unknown owner usually means that user or team lacks write access on the repo

### REPO-16 — Open critical/high Dependabot alerts are within the remediation SLA (warn)

_9 fail, 0 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |

**Fix:** Security -> Dependabot alerts: upgrade the flagged dependencies, or dismiss the alerts that don't apply with a reason

### REPO-17 — GitHub Actions are pinned to immutable commit SHAs (warn)

_21 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ⚠️ unknown | unpinned_action_ref_count not readable (missing scope?) |

**Fix:** In each workflow, replace mutable action tags or branches with the full 40-character commit SHA; use Dependabot to keep pinned actions updated

### REPO-20 — Required status checks apply to the latest commit (warn)

_40 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | protection not readable (missing scope?) |

**Fix:** Branch protection / ruleset -> require branches to be up to date before merging

### REPO-21 — New commits dismiss stale approvals (warn)

_40 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | protection not readable (missing scope?) |

**Fix:** Branch protection / ruleset -> dismiss stale pull request approvals when new commits are pushed

### REPO-22 — Pull request conversations must be resolved before merging (warn)

_40 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | protection not readable (missing scope?) |

**Fix:** Branch protection / ruleset -> require conversation resolution before merging

### REPO-23 — Default branch requires signed commits (warn)

_40 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | protection not readable (missing scope?) |

**Fix:** Branch protection / ruleset -> require signed commits on the default branch

### REPO-25 — Force pushes are disabled on the default branch (warn)

_35 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | protection not readable (missing scope?) |

**Fix:** Branch protection / ruleset -> block force pushes on the default branch

### REPO-26 — Deletion of the default branch is disabled (warn)

_36 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | protection not readable (missing scope?) |

**Fix:** Branch protection / ruleset -> block deletion of the default branch

### REPO-29 — Ruleset bypass is limited to automation (warn)

_1 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | protection not readable (missing scope?) |

**Fix:** Remove user/team/role bypass actors from default-branch rulesets or explicitly allowlist required automation

### REPO-31 — Repository publishes a security vulnerability reporting policy (warn)

_40 fail, 0 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |

**Fix:** Add SECURITY.md at the repository root, .github/, or docs/ with supported versions and a private reporting path

### REPO-36 — Repository configures CodeQL or another SAST tool (warn)

_30 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ⚠️ unknown | codeql_action_count not readable (missing scope?) |

**Fix:** Configure CodeQL or an approved SAST tool such as Semgrep or SonarQube

### REPO-37 — Dependency review runs in GitHub Actions (warn)

_39 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ⚠️ unknown | dependency_review_action_count not readable (missing scope?) |

**Fix:** Add actions/dependency-review-action to a required pull request workflow

### REPO-38 — CI workflows run on pull requests (warn)

_24 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ⚠️ unknown | workflow_pr_trigger_count not readable (missing scope?) |

**Fix:** Add a pull_request trigger to the repository's CI/test workflow

### REPO-4 — Dependabot / vulnerability alerts are enabled (warn)

✅ all 41 pass

### REPO-47 — Repository Actions secrets are rotated within policy (warn)

_6 fail, 0 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |

**Fix:** Rotate repository Actions secrets older than the configured limit; only secret metadata is collected

### REPO-48 — Individual repository admin access is limited (warn)

_41 fail, 0 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |

**Fix:** Review repository collaborators and downgrade unnecessary individual admins; tune the count and ratio per repo in --config

### REPO-49 — Production environments require reviewers (warn)

_41 fail, 0 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |

**Fix:** Add required reviewers to the production GitHub Environment; configure the environment name per repo in --config

### REPO-50 — Deployment rollback and recovery procedures are tested (warn)

_0 fail, 41 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/test | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ⚠️ unknown | assessment not established by the collector |

**Fix:** Document the rollback/recovery mechanism and evidence from a recent exercise

### REPO-52 — Private/internal repository forking is disabled (warn)

✅ all 39 pass

### ORG-12 — Teams use a documented work-tracking system (info)

_0 fail, 1 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ⚠️ unknown | assessment not established by the collector |

**Fix:** Record the work-tracking system and how engineering work is linked to repositories and changes

### ORG-13 — InnerSource contribution practices are established (info)

_0 fail, 1 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ⚠️ unknown | assessment not established by the collector |

**Fix:** Document cross-team contribution guidance, ownership, and review expectations

### ORG-15 — Repository onboarding is automated (info)

_0 fail, 1 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ⚠️ unknown | assessment not established by the collector |

**Fix:** Document the automated repository onboarding process and its security defaults

### ORG-17 — GitHub Discussions have an intentional purpose and owner (info)

_0 fail, 1 unknown of 1 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground | ⚠️ unknown | assessment not established by the collector |

**Fix:** Record whether Discussions are enabled and their purpose, or mark the assessment not applicable

### ORG-5 — No member invitations pending longer than the allowed age (info)

✅ all 1 pass

### ORG-9 — Org teams are private (privacy = closed) (info)

✅ all 1 pass

### REPO-10 — Default branch requires code owner review (info)

_39 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | protection not readable (missing scope?) |

**Fix:** Branch protection / ruleset on the default branch -> Require review from Code Owners

### REPO-13 — Automated dependency updates are configured (Dependabot or Renovate) (info)

_38 fail, 0 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |

**Fix:** Add .github/dependabot.yml (version updates) or a Renovate config (renovate.json)

### REPO-24 — Default branch requires linear history (info)

_40 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | protection not readable (missing scope?) |

**Fix:** Branch protection / ruleset -> require linear history on the default branch

### REPO-27 — Review dismissal is restricted (info)

_40 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | protection not readable (missing scope?) |

**Fix:** Branch protection -> restrict who can dismiss pull request reviews

### REPO-28 — Direct pushes to the default branch are restricted (info)

_37 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | protection not readable (missing scope?) |

**Fix:** Branch protection -> restrict who can push to the default branch

### REPO-30 — Branch naming conventions are documented and enforced (info)

_0 fail, 41 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/test | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ⚠️ unknown | assessment not established by the collector |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ⚠️ unknown | assessment not established by the collector |

**Fix:** Document the branch naming convention and how repository rulesets or CI enforce it

### REPO-32 — Repository has contribution guidance (info)

_40 fail, 0 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |

**Fix:** Add CONTRIBUTING.md at the repository root, .github/, or docs/

### REPO-33 — Repository has a pull request template (info)

_41 fail, 0 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |

**Fix:** Add a pull request template under .github/ or at the repository root

### REPO-34 — Repository has issue templates (info)

_40 fail, 0 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |

**Fix:** Add issue templates under .github/ISSUE_TEMPLATE/

### REPO-35 — Repository has Copilot instructions (info)

_38 fail, 0 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |

**Fix:** Add .github/copilot-instructions.md or scoped instructions under .github/instructions/

### REPO-39 — Repository uses reusable workflows (info)

_35 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ⚠️ unknown | reusable_workflow_count not readable (missing scope?) |

**Fix:** Consider centralizing repeated CI logic in reviewed reusable workflows

### REPO-40 — Workflow dependencies are cached (info)

_39 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ⚠️ unknown | cache_action_count not readable (missing scope?) |

**Fix:** Add a dependency cache such as actions/cache or a language setup action's cache option

### REPO-41 — Workflows use concurrency controls (info)

_38 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ⚠️ unknown | concurrency_workflow_count not readable (missing scope?) |

**Fix:** Add a concurrency group with cancel-in-progress where overlapping runs are safe

### REPO-42 — Workflow triggers use branch or path filters (info)

_24 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ⚠️ unknown | narrow_trigger_count not readable (missing scope?) |

**Fix:** Narrow workflow triggers with paths/branches filters where appropriate

### REPO-43 — Workflows use matrix strategies where useful (info)

_33 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ⚠️ unknown | matrix_job_count not readable (missing scope?) |

**Fix:** Use a strategy matrix to test supported language/runtime versions in parallel where applicable

### REPO-44 — Releases are automated (info)

_40 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ⚠️ unknown | release_automation_count not readable (missing scope?) |

**Fix:** Consider release-please, semantic-release, Changesets, or another reviewed release automation

### REPO-45 — Stale issues and pull requests are managed automatically (info)

_40 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ⚠️ unknown | stale_issue_automation_count not readable (missing scope?) |

**Fix:** Configure an issue lifecycle policy such as actions/stale if it fits the repository's support model

### REPO-46 — Repository has recent development activity (info)

_41 fail, 0 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-secret-filter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |

**Fix:** Review whether the repository is still needed; archive it or document why its current inactivity is expected

### REPO-5 — CODEOWNERS file is present (info)

_39 fail, 0 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-dependency-review | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-extending-codeql-at-scale-workflow | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/codeql-codebuild-integration-2 | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-default | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/assessement-orquestrator | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |

**Fix:** Add a CODEOWNERS file in .github/, repo root, or docs/

### REPO-51 — Repository accepts private vulnerability reports (info)

✅ all 2 pass

### REPO-6 — Public repository declares a license (info)

_2 fail, 0 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |

**Fix:** Add a LICENSE file so the repository's license is detected

### REPO-7 — Default branch is governed by at least one active ruleset (info)

_34 fail, 1 unknown of 41 checked_

| Target | Status | Notes |
|--------|--------|-------|
| ModusCreate-Perdigao-GHAS-Playground/GHAS-Training-Sessions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-with-third-party-tool | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-vscode-extensions | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-powerbi-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-code-scanning-integrated-with-linter | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/audit-toolkit | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-codeql-workflow-status | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-allow-github-ip | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/Org-Workflows | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/testing_repo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script_add_comment_to_alert | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/backduck-sample | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-monorepo | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/template-bfg | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-codeql-with-bazel-with-repo-detection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-delete-all-github-analysis | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/python-secure-template | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/script-repo-report | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/checkov | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/bazel-sonarqube | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/vscode | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-iac | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/hands-on-enforce-code-scanning | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-bypass-push-protection | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql-bug | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/copilot-jira-integration | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/.github-private | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/squadgram | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/multisolution-codeql | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/terraform-modules-usage | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/test-ruleset | ❌ fail |  |
| ModusCreate-Perdigao-GHAS-Playground/poc-dependabot | ⚠️ unknown | ruleset_count not readable (missing scope?) |

**Fix:** Repo Settings -> Rules: add a branch ruleset targeting the default branch

