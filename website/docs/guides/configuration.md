---
title: Policy configuration and assessments
description: Configure target-specific policy and human-reviewed evidence.
---

# Policy configuration and assessments

Pass a policy config with `--config`. It has two independent parts: target-specific rule overrides and optional human assessments. A complete example lives in [`examples/config.yaml`](https://github.com/jackchuka/postura/blob/main/examples/config.yaml).

## Target overrides

Overrides resolve most-specific-wins: rule defaults, then enterprise, org, and repo. Each key is resolved independently.

```yaml
enterprises:
  acme-inc:
    ENT-3: { enabled: true }

orgs:
  acme:
    ORG-2: { max_admins: 8 }
    ORG-3: { allowed_collaborators: [external-auditor] }

repos:
  acme/legacy-mirror:
    REPO-1: { enabled: false }
```

Within a rule block, `enabled` (boolean) and `severity` (string) tune the rule. Every other key is merged into that rule's `cfg` map. The rule must define a matching config key for a CEL expression to consume it.

## Human assessments

Assessments are useful for controls that cannot be verified from GitHub API facts. Mark the rule with `assessment: true`, then provide a target-specific entry under `assessments`:

```yaml
assessments:
  as_of: "2026-10-13"
  max_age_days: 180
  orgs:
    acme:
      ORG-11:
        title: "Disaster recovery and backup procedures are tested" # optional guidance
        question: "When did you last restore GitHub data from backup?" # optional guidance
        status: compliant
        evidence: "A restore exercise completed on the production backup set."
        reviewed_at: "2026-10-01"
        reviewed_by: "security-team"
```

Allowed status values are `compliant`, `partial`, `non_compliant`, and `not_applicable`. Every entry requires non-empty evidence and reviewer, plus `reviewed_at` in `YYYY-MM-DD` format. `title` and `question` are optional reviewer guidance and do not affect scoring. `assessments.as_of` and positive `max_age_days` are required when entries exist. Dates after `as_of` are rejected. Missing or stale assessments are unknown; `not_applicable` skips the check. Evidence and review metadata are included in findings.

Assessments are attestations, not API-verified facts. The freshness calculation uses the configured `as_of` date so evaluation is deterministic.

## Generate assessment inputs

Let posture create entries for all assessment-marked rules:

```bash
postura config template --rules examples/rules.yaml --org acme --output assessments.yaml
```

With an org and no repository arguments, the command discovers non-forked repositories. Pass `acme/api` and `acme/web` to limit repository entries, or use `--enterprise` for enterprise assessments. The generator does not overwrite an existing file, and blank values must be filled or removed before evaluation. Each entry is pre-filled with the rule's `title` and `question` so the reviewer knows what to answer in `evidence`. It generates assessments only, not policy overrides.