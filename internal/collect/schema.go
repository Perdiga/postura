// Package collect gathers facts from the GitHub API into fully-normalized JSON
// objects. It is deterministic and rule-unaware: it only fetches facts. Every
// documented field is always present (null/false/0/[] when unknown) so CEL
// rules need no defensive guards — except fields the token could not read,
// which stay absent so a rule reading them reports unknown, not a false pass.
package collect

// Canonical fact-field names per scope. The engine declares exactly these as
// CEL variables, so a rule referencing an unknown field fails to compile (the
// lint), and collect guarantees each is populated on every object.
var (
	EnterpriseVars = []string{
		"enterprise",
		"saml_sso_url",
		"default_repository_permission",
		"ip_allow_list_enabled",
		"allow_private_repository_forking",
		"members_can_change_repository_visibility",
		"members_can_invite_collaborators",
		"two_factor_required_setting",
		"members_can_create_public_repositories",
		"members_can_delete_repository",
		"notification_delivery_restriction_enabled",
		"advanced_security_enabled_for_new_repositories",
		"secret_scanning_enabled_for_new_repositories",
		"secret_scanning_push_protection_enabled_for_new_repositories",
		"outside_collaborator_count",
		"outside_collaborators",
	}

	OrgVars = []string{
		"org",
		"two_factor_requirement_enabled",
		"default_repository_permission",
		"members_can_create_repositories",
		"members_can_create_public_repositories",
		"advanced_security_enabled_for_new_repositories",
		"members_can_fork_private_repositories",
		"secret_scanning_enabled_for_new_repositories",
		"secret_scanning_push_protection_enabled_for_new_repositories",
		"dependabot_alerts_enabled_for_new_repositories",
		"admins", "admin_count",
		"members", "member_count",
		"outside_collaborators", "outside_collaborator_count",
		"teams",
		"organization_roles",
		"invitations",
		"actions_enabled_repositories", "actions_allowed_actions",
		"actions_default_workflow_permissions", "actions_can_approve_pull_request_reviews",
		"oldest_actions_secret_age_days",
		"runner_groups",
	}

	RepoVars = []string{
		"name", "owner", "archived", "visibility", "allow_forking", "last_pushed_days", "oldest_actions_secret_age_days",
		"collaborator_count", "admin_collaborator_count",
		"environments",
		"secret_scanning", "secret_scanning_push_protection",
		"vulnerability_alerts", "private_vulnerability_reporting", "codeowners", "codeowners_error_count", "license",
		"dependabot_config", "renovate_config", "dependabot_alerts",
		"unpinned_action_ref_count", "workflow_permission_issue_count", "repository_files",
		"workflow_pr_trigger_count", "codeql_action_count", "dependency_review_action_count",
		"cache_action_count", "concurrency_workflow_count", "matrix_job_count",
		"narrow_trigger_count", "reusable_workflow_count", "release_automation_count",
		"stale_issue_automation_count",
		"teams", "ruleset_count", "protection",
	}
)
