---
title: GitHub permissions
description: Token access required to collect posture facts.
---

# GitHub permissions

Use `GITHUB_TOKEN` or `GH_TOKEN`; if neither is set, posture falls back to a logged-in GitHub CLI through `gh auth token`. Grant read access only for the scopes and features you intend to audit. The exact fine-grained permission names can vary with endpoint and GitHub plan; use `postura permissions check` against your targets to see which capability groups are unreadable.

| Audit capability | Access to check |
| --- | --- |
| Repository settings and metadata | Repository Metadata read; repository admin read for branch protection, collaborators, and some security settings. |
| Repository content and workflow analysis | Repository Contents read; classic PATs generally need `repo` for private repos. |
| Actions secret metadata and environments | Actions read at repo or org scope. Only metadata/age is collected, never secret values. |
| Org settings, membership, owners, and roles | Classic PAT: `admin:org` read. Fine-grained PAT: Organization members read and the relevant Administration/custom-role read permission. |
| Org Actions policy and runner groups | Organization Actions/Administration and self-hosted runner read access; classic PATs generally use `admin:org`. |
| Enterprise account settings | Enterprise-owner token with `read:enterprise`. |
| Dependabot alerts | Dependabot alerts read; classic PATs require `security_events` and appropriate repository/org access. |
| Private vulnerability reporting for public repos | Public repository access; fine-grained PAT needs repository Metadata read. |

Without required access, affected facts are omitted and rules depending on them become unknown, never a false pass. Some missing fields can also mean a GitHub feature is unavailable or ineligible for the plan. The org-wide Dependabot sweep additionally needs org-owner or security-manager standing; otherwise alerts are read per repository.

## Diagnose access

```bash
postura permissions check --org acme --enterprise acme-inc
```

The command runs read-only probes and groups unreadable fields by capability. A permission-check failure is not a policy finding: it says posture could not verify data, not that the target violated a rule.

## Security of collected data

Postura is read-only and does not retrieve Actions secret values. Fact snapshots can still contain sensitive metadata such as collaborator names, security settings, and alert details. Protect saved JSON files and CI artifacts according to your organization's data-handling policy.