---
title: Fact fields
description: CEL facts collected at enterprise, organization, and repository scope.
---

# Fact fields

The collector exposes these names as top-level CEL variables. The authoritative schema is [`internal/collect/schema.go`](https://github.com/jackchuka/postura/blob/main/internal/collect/schema.go). A field the token cannot read is omitted, so a rule that needs it reports unknown. Some facts are only available when a GitHub feature is enabled or eligible.

## Enterprise

| Field | Meaning |
| --- | --- |
| `enterprise` | Enterprise slug. |
| `saml_sso_url` | SAML SSO URL, or null. |
| `default_repository_permission` | Enterprise base repo permission: `NONE`, `READ`, `WRITE`, `ADMIN`, or `NO_POLICY`. |
| `ip_allow_list_enabled` | IP allow-list setting. |
| `allow_private_repository_forking` | Private/internal forking policy. |
| `members_can_change_repository_visibility` | Enterprise visibility-change policy. |
| `members_can_invite_collaborators` | Outside-collaborator invitation policy. |
| `two_factor_required_setting` | Enterprise-wide 2FA policy. |
| `members_can_create_public_repositories` | Whether members can create public repos. |
| `members_can_delete_repository` | Repository deletion/transfer policy. |
| `notification_delivery_restriction_enabled` | Whether notifications are limited to verified domains. |
| `advanced_security_enabled_for_new_repositories` | GHAS default for new repositories. |
| `secret_scanning_enabled_for_new_repositories` | Secret scanning default for new repositories. |
| `secret_scanning_push_protection_enabled_for_new_repositories` | Push protection default for new repositories. |
| `outside_collaborators` | Outside-collaborator logins. |
| `outside_collaborator_count` | Number of outside collaborators. |

## Organization

| Field | Meaning |
| --- | --- |
| `org` | Org login. |
| `two_factor_requirement_enabled` | Whether org-wide 2FA is enforced. |
| `default_repository_permission` | Base member permission: `none`, `read`, `write`, or `admin`. |
| `members_can_create_repositories` | Whether members may create repos. |
| `members_can_create_public_repositories` | Whether members may create public repos. |
| `members_can_fork_private_repositories` | Whether members may fork private/internal repos. |
| `advanced_security_enabled_for_new_repositories` | GHAS default for new repositories. |
| `secret_scanning_enabled_for_new_repositories` | Secret scanning default for new repositories. |
| `secret_scanning_push_protection_enabled_for_new_repositories` | Push protection default for new repositories. |
| `dependabot_alerts_enabled_for_new_repositories` | Dependabot alerts default for new repositories. |
| `admins`, `admin_count` | Org owner logins and count. |
| `members`, `member_count` | Org member logins and count. |
| `outside_collaborators`, `outside_collaborator_count` | Outside-collaborator logins and count. |
| `teams` | Org teams, each with `slug` and `privacy`. |
| `organization_roles` | Role assignments, each with `role`, `teams`, and `users`. |
| `invitations` | Pending invitations, each with `login` and `age_days`. |
| `actions_enabled_repositories` | Actions repo scope: `all`, `selected`, or `none`. |
| `actions_allowed_actions` | Actions policy: `all`, `local_only`, or `selected`. |
| `actions_default_workflow_permissions` | Default workflow token permission: `read` or `write`. |
| `actions_can_approve_pull_request_reviews` | Whether workflows can approve pull requests. |
| `oldest_actions_secret_age_days` | Age of oldest Actions secret; metadata only. |
| `runner_groups` | Org runner groups, visibility, and public-repo access settings. |

## Repository

| Field | Meaning |
| --- | --- |
| `name`, `owner` | `owner/repo` name and owning org. |
| `archived`, `visibility`, `allow_forking` | Archive status, visibility (`public`, `private`, `internal`), and fork policy. |
| `last_pushed_days` | Whole days since last push, computed during collection. |
| `oldest_actions_secret_age_days` | Age of the oldest Actions secret; metadata only. |
| `collaborator_count`, `admin_collaborator_count` | Readable collaborator and individual-admin counts. |
| `environments` | Deployment environments, protection-rule counts, and required-reviewer status. |
| `secret_scanning`, `secret_scanning_push_protection` | Secret scanning and push protection status. |
| `vulnerability_alerts` | Whether Dependabot alerts are enabled. |
| `private_vulnerability_reporting` | Whether private vulnerability reports are accepted. |
| `dependabot_alerts` | Open-alert totals by severity, including oldest open age when available. |
| `dependabot_config`, `renovate_config` | Whether committed Dependabot or Renovate config is detected. |
| `unpinned_action_ref_count` | Remote Actions/reusable workflows not pinned to a full commit SHA. |
| `workflow_permission_issue_count` | Jobs with implicit, malformed, or write-level `GITHUB_TOKEN` permissions. |
| `repository_files` | Known presence of README, SECURITY, contribution/template, Copilot, and SAST files. |
| `workflow_pr_trigger_count` | Workflows triggered by `pull_request` or `pull_request_target`. |
| `codeql_action_count`, `dependency_review_action_count`, `cache_action_count` | Counts of those workflow action references. |
| `concurrency_workflow_count`, `matrix_job_count` | Workflows defining concurrency and jobs defining a matrix. |
| `narrow_trigger_count` | Workflows with branch/path filters on an event. |
| `reusable_workflow_count` | Reusable workflow references. |
| `release_automation_count`, `stale_issue_automation_count` | Recognized release and stale-issue automation references. |
| `codeowners`, `codeowners_error_count` | CODEOWNERS presence and GitHub-reported lint errors. |
| `license` | SPDX license ID, or null. |
| `teams` | Teams with a direct grant, each with `slug` and `permission`. |
| `ruleset_count` | Active branch rulesets on the default branch. |
| `protection` | Classic and ruleset branch-protection facts, including reviews, checks, signed commits, linear history, push/deletion controls, and bypass actor types. |

## Important availability details

- Dependabot alert details can be absent when alerts are disabled or unreadable. Gate alert rules on `vulnerability_alerts == true`; otherwise a missing collection must not look like zero alerts.
- Alert ages are calculated at collect time and frozen in the fact snapshot. Age is based on the original alert creation date; an undated alert makes its age unknown.
- Dependabot and Renovate config detection checks committed repository files only. Shared presets or app defaults without a repo-level file are not detected.
- Branch protection combines classic protection and repository rulesets, using the stricter outcome.

The rule compiler and collector schema are tested together. See [writing rules](../guides/rules.md) for expression examples.