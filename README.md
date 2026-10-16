<div align="center">

<h1>postura</h1>

<p>
  <strong>A deterministic CLI that audits your GitHub enterprise, orgs, and repos<br>against a security baseline you own — rules are data, the bar is per-target.</strong>
</p>

<p>
  <a href="https://github.com/Perdiga/postura/actions/workflows/test.yml"><img src="https://github.com/Perdiga/postura/actions/workflows/test.yml/badge.svg" alt="Test"></a>
  <a href="https://github.com/Perdiga/postura/releases"><img src="https://img.shields.io/github/v/release/Perdiga/postura?sort=semver" alt="Release"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License: MIT"></a>
</p>

</div>

---

This is an independently maintained fork of [jackchuka/postura](https://github.com/jackchuka/postura).
Installations, releases, and documentation links below refer to [Perdiga/postura](https://github.com/Perdiga/postura).

**postura** audits a GitHub organization's **security posture** — branch protection,
2FA enforcement, secret scanning, outside collaborators, org-role grants, and more —
against a **per-target-configurable** policy baseline. Rules are data
([CEL](https://github.com/google/cel-go) expressions in YAML); the same ruleset can hold
a different bar for each enterprise, org, or repo. It is **read-only** — it never mutates
anything on GitHub — and **deterministic**, so it gates cleanly in CI.

Full guides and reference documentation are in the [Docusaurus site](website/docs/index.md)
([run it locally](website/)).

```console
$ postura run --org acme --enterprise acme-inc --rules examples/rules.yaml
```

```markdown
# GitHub audit

_3 fail, 1 unknown across 12 repos + org + enterprise_

### ORG-1 — Two-factor authentication is enforced org-wide (error)

_1 fail, 0 unknown of 1 checked_

| Target | Status  | Notes                                   |
| ------ | ------- | --------------------------------------- |
| acme   | ❌ fail | two_factor_requirement_enabled is false |

**Fix:** Org Settings → Authentication security: require two-factor authentication

### REPO-1 — Default branch requires a pull request before merging (error)

_2 fail, 0 unknown of 12 checked_

| Target   | Status  | Notes                                  |
| -------- | ------- | -------------------------------------- |
| acme/web | ❌ fail | required_approving_review_count is 0   |
| acme/api | ❌ fail | no branch protection on default branch |

**Fix:** Settings → Branches: require a pull request before merging

### ENT-1 — SAML single sign-on is configured for the enterprise (error)

_0 fail, 1 unknown of 1 checked_

| Target   | Status     | Notes                                             |
| -------- | ---------- | ------------------------------------------------- |
| acme-inc | ⚠️ unknown | needs an enterprise-owner token (read:enterprise) |
```

## Why postura

Tools like [legitify](https://github.com/Legit-Labs/legitify) apply a _global_ policy —
you can disable a check everywhere, but not hold a different bar for different targets.
[OSSF Scorecard](https://github.com/ossf/scorecard) does take per-repo config, but it's
_decentralized_ (a `.github/scorecard.yml` each repo owns) and scoped to OSS/supply-chain
repo health — not org/enterprise settings like 2FA enforcement, outside collaborators, or
org-role grants.

postura's difference is **centralized, layered per-target policy and assessment config**:
one ruleset a security team operates, where `acme` may allow 8 owners while `acme-labs`
allows 3, a legacy repo can be exempted from a rule, and enterprise + org + repo all
audit from the same policy. Human assessments can also be recorded centrally per target
with evidence, reviewer, and review date — no repo-owned policy file is required.

API-collected GitHub facts remain deterministic and CI-gateable. Human assessments are
attestations, not independently verified facts; posture marks missing or stale entries
unknown and includes their evidence in the report. Thresholds, allowlists, and assessment
freshness are configurable per target.

|                       | postura                                                        |
| --------------------- | -------------------------------------------------------------- |
| **Read-only**         | Never writes to GitHub — pure audit                            |
| **Rules are data**    | CEL expressions in YAML, compiled at load (typos fail fast)    |
| **Per-target bar**    | Same ruleset, different thresholds per enterprise / org / repo |
| **Honest about gaps** | A fact it can't read → `unknown`, never a false pass           |
| **CI-native**         | `--fail-on` gate + SARIF output for code scanning              |
| **Three scopes**      | Enterprise account settings, org settings, per-repo posture    |

## Install

```bash
go install github.com/Perdiga/postura@latest
```

Or grab a prebuilt binary from [Releases](https://github.com/Perdiga/postura/releases)
(Linux / macOS / Windows, amd64 + arm64).

Each release archive includes the binary and the example `examples/rules.yaml`
and `examples/config.yaml` files. The examples are also available as separate
`rules.yaml` and `config.yaml` release downloads. Customize them for your
organization before running an audit.

### Run in a Dev Container

Open this repository in VS Code and run **Dev Containers: Reopen in Container** from
the Command Palette. The container uses the Go version in `go.mod` and downloads
dependencies when it is created. Set `GITHUB_TOKEN` or `GH_TOKEN` in your local
environment before opening the container, then run:

```bash
go run . run --org acme --rules examples/rules.yaml
```

### Authentication

postura needs a GitHub token via `GITHUB_TOKEN` / `GH_TOKEN`, or a logged-in
[`gh`](https://cli.github.com/) (`gh auth token` is used as a fallback).

| Scope of audit              | Token requirement                          |
| --------------------------- | ------------------------------------------ |
| Repo settings               | repo read; repository admin for collaborators and branch protection |
| Actions secret metadata     | Actions read at repo/org scope             |
| Private vulnerability reporting | public repos only; fine-grained PAT needs repository Metadata read |
| Environments                | repo read                                  |
| Org settings, owners, roles | `admin:org` read                           |
| Org Actions settings/runners | org Actions read / `admin:org`             |
| Enterprise account settings | enterprise-owner token (`read:enterprise`) |
| Dependabot alerts           | `security_events`                          |

Without the right scope, the relevant facts report as **unknown** rather than a false pass.
The org-wide Dependabot alert sweep additionally needs org-owner or security-manager
standing; without it, alerts are read per repo instead.

## Usage

```text
postura run --rules R [owner/repo ...]    # collect + evaluate + report (the common case)
postura collect [owner/repo ...]          # write facts as JSON (cache / offline eval)
postura eval --rules R <facts.json>       # evaluate a facts file — pure, no network
postura rules list --rules R              # list every rule
postura rules explain --rules R <id>      # show a rule's CEL, config knobs, and fix hint
postura config template --rules R --org acme --output config.yaml # generate assessment inputs
postura permissions check --org acme     # test collector access for the org
```

`--rules` is required for every command except `collect`, which only gathers facts:
`collect` writes facts; `eval` and `run` apply rules to them. postura ships **no
built-in ruleset** — you always pass your own. A generic GitHub security baseline you can
copy and tune lives in [`examples/rules.yaml`](examples/rules.yaml).

With no repo arguments, every non-forked repo in the org is audited. Pass `owner/repo`
tokens to scope to specific repos. `--org` is repeatable, so several orgs can be audited
in one run.

### Flags

Shared by `run` and `eval`:

| Flag        | Default                          | Meaning                                                                          |
| ----------- | -------------------------------- | -------------------------------------------------------------------------------- |
| `--config`  | —                                | policy-config YAML with per-target overrides                                     |
| `--rules`   | — (**required**)                 | ruleset YAML (see [`examples/rules.yaml`](examples/rules.yaml))                  |
| `--format`  | `md`                             | `md` \| `json` \| `sarif` \| `html`                                              |
| `--fail-on` | `error` (`run`), `none` (`eval`) | exit non-zero on a fail at/above this severity (`error`\|`warn`\|`info`\|`none`) |

Shared by `run` and `collect` (the commands that talk to GitHub):

| Flag           | Default        | Meaning                                                                                                                 |
| -------------- | -------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `--org`        | —              | org(s) to audit — repeatable/comma-separated, or inferred from `owner/repo` args; optional for an enterprise-only audit |
| `--enterprise` | —              | enterprise slug to also audit (needs an enterprise-owner token)                                                         |
| `--scope`      | all applicable | restrict to scopes: `enterprise` \| `org` \| `repos` (repeatable, e.g. `--scope org,repos`)                             |

`eval` takes its org/enterprise from the facts file, so it has none of the GitHub-facing
flags above. SARIF output is uploadable to GitHub code scanning and usable as audit evidence.
HTML output is a self-contained maturity report for enterprise/org owners, scored against the
[GitHub Well-Architected](https://learn.github.com/well-architected) pillars.

HTML maturity can tolerate warnings without changing their severity or finding
status. Add this top-level block to the file passed with `--config`:

```yaml
maturity:
  max_warnings_per_tier: 2
  ignore_warnings: false
```

This allows up to two failing/unknown warnings **per target, per maturity tier**.
Failing/unknown errors always block progression; info never blocks. The default
allowance is zero (existing behavior). To ignore all warning blockers, set
`ignore_warnings: true` and leave `max_warnings_per_tier` at zero. Scores, findings,
and `--fail-on` exit gates are unchanged. The policy applies to enterprise,
organization, and repository maturity in the report.

### Examples

```bash
# Enterprise + org + every repo (default scope = all applicable)
postura run --org acme --enterprise acme-inc --rules examples/rules.yaml

# Generate blank assessment inputs for an enterprise, org, and every non-forked repo
postura config template --rules examples/rules.yaml --enterprise acme-inc --org acme --output config.yaml

# Check whether the token can read postura's org/repo/enterprise fact sources
postura permissions check --org acme --enterprise acme-inc

# Limit generated repository assessments to specific repos
postura config template --rules examples/rules.yaml --org acme acme/api acme/web --output config.yaml

# Several orgs in one run (repeatable or comma-separated --org)
postura run --org acme --org acme-labs --rules examples/rules.yaml
postura run --org acme,acme-labs --rules examples/rules.yaml

# Enterprise only — no --org needed (scope defaults to enterprise)
postura run --enterprise acme-inc --rules examples/rules.yaml

# Org-level checks only (skip the per-repo sweep)
postura run --org acme --scope org --rules examples/rules.yaml

# Specific repos (org inferred from the args), audited alongside the enterprise
postura run --enterprise acme-inc acme/web acme/api --rules examples/rules.yaml

# Collect enterprise facts to JSON, then evaluate offline
postura collect --enterprise acme-inc --scope enterprise > facts.json
postura eval --rules examples/rules.yaml facts.json

# CI gate: fail on any warn-or-worse, emit SARIF for code scanning
postura run --org acme --enterprise acme-inc --scope enterprise,org \
  --fail-on warn --format sarif --rules examples/rules.yaml > postura.sarif
```

`--scope` is repeatable and comma-separated (`--scope org,repos` == `--scope org --scope
repos`). With no `--scope`, postura audits everything nameable: org + its repos when an
org is given, plus the enterprise when `--enterprise` is given. Each scope needs its input
— so `--scope enterprise` without `--enterprise` (or `--scope org` with no org) is an error.

An enterprise is audited only for its **account-level settings**; it does not enumerate the
orgs beneath it. Audit those by naming them with `--org` — repos auto-expand from an org,
but orgs are named entities.

## The ruleset

You supply the ruleset with `--rules`; postura embeds none. Each rule's `pass` is a
[CEL](https://github.com/google/cel-go) expression over the scope's fact fields (bound as
top-level variables) plus `cfg` (the rule's merged config):

```yaml
- id: ORG-2
  scope: org
  severity: warn
  title: Owner/admin count is proportionate to org size
  config:
    max_admins: 5 # overridable per enterprise/org/repo
    max_admin_ratio: 0.25
  pass: >-
    admin_count <= cfg.max_admins &&
    (member_count == 0 || double(admin_count) / double(member_count) <= cfg.max_admin_ratio)
  fix_hint: "Org Settings → People: review who holds Owner"
```

- **`severity`** — `error` \| `warn` \| `info`. Drives report ordering, the `--fail-on`
  gate, and the SARIF level (`error` → error, `warn` → warning, `info` → note).
- **`applies_when`** — scope a rule to matching targets (e.g. public repos only).
- **`enabled: false`** — off by default; opt in per target via config.
- **`assessment: true`** — the rule consumes a human-reviewed entry from `assessments`;
  `postura config template` uses this marker to generate target-specific blank inputs.
- **`question`** — for assessment rules, the question the reviewer should answer; copied
  into generated templates alongside the rule's `title`.

The whole ruleset is CEL-compiled at load, so a typo'd field name (an undeclared variable)
fails fast there. When a fact a rule needs can't be read (e.g. missing token scope), the
collector omits it and postura reports **unknown** (⚠️) for that target automatically — no
rule annotation needed.

## Per-target configuration

A policy-config file (`--config`) overrides rule settings, merged most-specific-wins:
**rule default → enterprise → org → repo**.

```yaml
# config.yaml
orgs:
  acme:
    ORG-2: { max_admins: 8 } # raise the bar
    ORG-3: { allowed_collaborators: [contractor-x] } # allowlist → deterministic
  acme-labs:
    REPO-8: { enabled: true } # opt into an off-by-default rule
repos:
  acme/legacy-thing:
    REPO-1: { enabled: false } # exempt one repo
```

Within a rule's block, `enabled` (bool) and `severity` (string) tune the rule; every other
key is merged into its `cfg`.

Human assessments use a separate `assessments` section keyed by scope, target, and rule
ID. Set `assessments.as_of` (`YYYY-MM-DD`) and `assessments.max_age_days` to make
freshness deterministic. Each entry requires a status (`compliant`, `partial`,
`non_compliant`, or `not_applicable`), evidence, `reviewed_at` (`YYYY-MM-DD`), and
`reviewed_by`. Missing or stale inputs report `unknown`; `not_applicable` skips that
check. Findings include the supplied evidence and review metadata. See the commented examples in
[`examples/config.yaml`](examples/config.yaml).

`postura config template` generates blank entries for every assessment-marked rule in the
selected targets. With `--org acme` and no repository arguments, it discovers all
non-forked repositories; pass `owner/repo` arguments to narrow the repository set. The
command needs a GitHub token only when it discovers repositories, writes `config.yaml` by
default, and refuses to overwrite an existing file. Each entry is pre-filled with the rule's
`title` and `question` to guide the reviewer; these are informational and do not affect
scoring. Blank fields must be completed before
the config can be used. The generated file contains human assessments only, not
machine-collected GitHub settings or per-target policy overrides.

`postura permissions check` runs read-only collectors and reports which capability groups
could not be read, with the relevant token permissions to check. It does not require the
GitHub CLI: provide `GITHUB_TOKEN` or `GH_TOKEN` in the environment. An `unknown`/missing
fact can also mean a GitHub feature is unavailable or ineligible; policy `fail` results
are separate from permission diagnostics.

## Facts available to rules

Rules read these fields (bound as top-level CEL variables), grouped by scope. The source of
truth is [`internal/collect/schema.go`](internal/collect/schema.go); a test fails if this
list drifts from it. A field that couldn't be read is omitted (→ the rule reports unknown).

<details>
<summary><strong>Enterprise scope</strong></summary>

| Field                              | Meaning                                                                     |
| ---------------------------------- | --------------------------------------------------------------------------- |
| `enterprise`                       | enterprise slug (the audited target)                                        |
| `saml_sso_url`                     | configured SAML SSO URL, or null                                            |
| `default_repository_permission`    | enterprise base repo permission (`NONE`/`READ`/`WRITE`/`ADMIN`/`NO_POLICY`) |
| `ip_allow_list_enabled`            | IP allow list setting (`ENABLED`/`DISABLED`)                                |
| `allow_private_repository_forking` | private/internal forking policy (`ENABLED`/`DISABLED`/`NO_POLICY`)          |
| `members_can_change_repository_visibility` | enterprise policy for repository visibility changes                  |
| `members_can_invite_collaborators`  | enterprise policy for inviting outside collaborators                    |
| `two_factor_required_setting`      | enterprise-wide 2FA policy (`ENABLED`/`DISABLED`)                         |
| `members_can_create_public_repositories` | whether enterprise members can create public repositories            |
| `members_can_delete_repository`    | enterprise repository deletion/transfer policy                            |
| `notification_delivery_restriction_enabled` | whether notifications are limited to verified domains              |
| `advanced_security_enabled_for_new_repositories` | whether GHAS is enabled by default for new repositories at org/enterprise scope |
| `outside_collaborators`            | outside-collaborator logins                                                 |
| `outside_collaborator_count`       | number of outside collaborators                                             |

</details>

<details>
<summary><strong>Org scope</strong></summary>

| Field                                                          | Meaning                                                     |
| -------------------------------------------------------------- | ----------------------------------------------------------- |
| `org`                                                          | org login (the audited target)                              |
| `two_factor_requirement_enabled`                               | whether 2FA is enforced org-wide                            |
| `default_repository_permission`                                | base member repo permission (`none`/`read`/`write`/`admin`) |
| `members_can_create_repositories`                              | whether members may create repos                            |
| `members_can_create_public_repositories`                       | whether members may create public repos                     |
| `members_can_fork_private_repositories`                        | whether members may fork private/internal repos             |
| `secret_scanning_enabled_for_new_repositories`                 | new-repo default: secret scanning                           |
| `secret_scanning_push_protection_enabled_for_new_repositories` | new-repo default: push protection                           |
| `dependabot_alerts_enabled_for_new_repositories`               | new-repo default: Dependabot alerts                         |
| `advanced_security_enabled_for_new_repositories`               | new-repo default: GitHub Advanced Security                  |
| `admins`                                                       | org owner logins                                            |
| `admin_count`                                                  | number of org owners                                        |
| `members`                                                      | org member logins                                           |
| `member_count`                                                 | number of members                                           |
| `outside_collaborators`                                        | outside-collaborator logins                                 |
| `outside_collaborator_count`                                   | number of outside collaborators                             |
| `teams`                                                        | org teams, each `{slug, privacy}`                           |
| `organization_roles`                                           | org-role assignments, each `{role, teams, users}`           |
| `invitations`                                                  | pending invitations, each `{login, age_days}`               |
| `actions_enabled_repositories`                                 | Actions repository scope (`all`, `selected`, or `none`)      |
| `actions_allowed_actions`                                      | Actions allow policy (`all`, `local_only`, or `selected`)    |
| `actions_default_workflow_permissions`                         | default workflow token permission (`read` or `write`)        |
| `actions_can_approve_pull_request_reviews`                     | whether workflows may approve pull requests                 |
| `runner_groups`                                                  | org runner groups with visibility and public-repository access settings |

</details>

<details>
<summary><strong>Repo scope</strong></summary>

| Field                             | Meaning                                                                                                                                        |
| --------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `name`                            | `owner/repo`                                                                                                                                   |
| `owner`                           | owning org                                                                                                                                     |
| `archived`                        | whether the repo is archived                                                                                                                   |
| `visibility`                      | `public` / `private` / `internal`                                                                                                              |
| `allow_forking`                   | whether this repository permits forks                                                                                                         |
| `last_pushed_days`                | whole days since the last push, computed at collect time                                                                                       |
| `oldest_actions_secret_age_days`  | age of the oldest Actions secret; available at org and repo scope, metadata only                                                                |
| `collaborator_count`               | number of repository collaborators with readable permission details                                                                            |
| `admin_collaborator_count`         | number of individual repository collaborators with admin permission                                                                            |
| `environments`                    | deployment environments with protection-rule count and required-reviewer status                                                                |
| `secret_scanning`                 | secret-scanning status (`enabled`/`disabled`)                                                                                                  |
| `secret_scanning_push_protection` | push-protection status (`enabled`/`disabled`)                                                                                                  |
| `vulnerability_alerts`            | whether Dependabot alerts are enabled                                                                                                          |
| `private_vulnerability_reporting`  | whether the repository accepts private vulnerability reports                                                                                  |
| `dependabot_alerts`               | open Dependabot alerts: `{total, critical, high, medium, low}`, each severity `{count, oldest_open_days}` |
| `dependabot_config`               | whether a Dependabot version-updates config exists (`.github/dependabot.yml`)                                                                  |
| `renovate_config`                 | whether a Renovate config exists (`renovate.json`, `.renovaterc*`, or under `.github/`)                                                        |
| `unpinned_action_ref_count`       | number of remote Actions and reusable workflows not pinned to a full commit SHA; absent if workflow files could not be read or parsed          |
| `workflow_permission_issue_count` | number of jobs with implicit, malformed, or write-level `GITHUB_TOKEN` permissions; absent if workflow files could not be read or parsed       |
| `repository_files`               | known presence booleans for README, SECURITY.md, contribution guides, templates, Copilot instructions, and common SAST config files             |
| `workflow_pr_trigger_count`      | number of workflows with a `pull_request` or `pull_request_target` trigger                                                                       |
| `codeql_action_count`             | CodeQL action references in workflows                                                                                                            |
| `dependency_review_action_count`  | dependency-review action references in workflows                                                                                                 |
| `cache_action_count`              | `actions/cache` references in workflows                                                                                                          |
| `concurrency_workflow_count`     | workflows defining a concurrency group                                                                                                           |
| `matrix_job_count`                 | workflow jobs defining a strategy matrix                                                                                                         |
| `narrow_trigger_count`            | workflows with branch/path filters on an event                                                                                                   |
| `reusable_workflow_count`         | reusable workflow references                                                                                                                     |
| `release_automation_count`        | recognized release automation action references                                                                                                 |
| `stale_issue_automation_count`    | recognized stale-issue automation action references                                                                                              |
| `codeowners`                      | whether a CODEOWNERS file exists                                                                                                               |
| `codeowners_error_count`          | errors GitHub's CODEOWNERS linter reports (unknown owners, bad syntax); `0` when no file exists                                                |
| `license`                         | SPDX license id, or null                                                                                                                       |
| `teams`                           | teams with a direct grant, each `{slug, permission}`                                                                                           |
| `ruleset_count`                   | active branch rulesets on the default branch                                                                                                   |
| `protection`                      | classic + ruleset facts: PR reviews, status checks, stale-review/conversation rules, signed commits, linear history, force-push/deletion controls, push/dismissal restrictions, and bypass actor types |

</details>

> **Note:** `dependabot_alerts` is absent when Dependabot alerts are known to be
> turned off for the repo, and when no alert listing could be read (a token
> without `security_events`), so a rule over it reports unknown rather than a
> false "no alerts". The org sweep reports only repos that have alerts enabled,
> so a rule over this fact should still gate on `vulnerability_alerts == true`
> for the case where that flag itself is unreadable — see REPO-16 in
> [`examples/rules.yaml`](examples/rules.yaml). Ages are computed at collect time
> and frozen into the facts, so re-running `eval` on cached facts stays
> reproducible; they count from the alert's `created_at`, which GitHub never
> resets, so a reopened alert carries its full original age. A severity bucket
> holding an alert whose age GitHub does not report omits `oldest_open_days`
> entirely, since an undatable alert is of unknown age, not a fresh one.

> **Note:** `dependabot_config` / `renovate_config` detect only *committed* config.
> A repo onboarded to Renovate via a shared org preset (in the org's `.github`
> repo) or the Renovate App's default, with no repo-level file, is not detected.

## How it works

```text
            ┌───────────┐      ┌──────────────┐      ┌───────────────┐
 GitHub ──► │  collect  │ ───► │     eval     │ ───► │    report     │
 REST +     │  (facts)  │ json │ (CEL × cfg)  │      │ md/json/sarif │
 GraphQL    └───────────┘      └──────────────┘      └───────────────┘
                read-only        deterministic          + exit code
```

`run` chains all three. Splitting `collect` (the only networked step) from `eval` lets you
cache facts and re-evaluate offline against different rulesets. Branch protection is read
from **both** classic protection and repository rulesets, folding the stricter outcome.

## License

MIT — see [LICENSE](LICENSE).
