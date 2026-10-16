---
title: Reports and exit codes
description: Interpret findings, formats, severity gates, and unknown results.
---

# Reports and exit codes

Each evaluated rule and target produces a finding. Outcomes are `pass`, `fail`, `unknown`, or a skipped/non-applicable check. Findings include the target, rule ID and title, severity, notes, and remediation hint; human-assessment findings also include submitted evidence and review metadata.

## Unknown is not a pass

When a required fact is absent because posture could not read it, the result is unknown. This keeps an under-scoped token or unavailable GitHub feature from silently satisfying a security rule. Use [`permissions check`](./permissions.md) to distinguish common access gaps; feature availability can also affect results.

## Output formats

| Format | Use |
| --- | --- |
| `md` | Human-readable report; default for `run`. |
| `json` | Structured findings for scripts and integrations. |
| `sarif` | Static Analysis Results Interchange Format for code scanning and audit evidence. |
| `html` | Self-contained maturity report for enterprise and org owners. |

SARIF levels map from rule severity: `error` becomes `error`, `warn` becomes `warning`, and `info` becomes `note`.

## HTML maturity report

```bash
postura run --enterprise acme-ent --org acme --format html --fail-on none \
  --rules examples/rules.yaml > postura.html
```

The HTML report scores posture against the [GitHub Well-Architected](https://learn.github.com/well-architected) pillars (Productivity, Collaboration, Application Security, Governance, Architecture) using each rule's `principle_domain` and `maturity_level`. It contains:

- **Executive summary** — overall pass rate, the enterprise's maturity level, and the distribution of orgs and repositories across levels.
- **Pillar scorecard** — pass rate per pillar for each scope. Pillars without automated checks link to the framework for self-assessment.
- **Improvement roadmap** — failing and unknown checks grouped by maturity tier, including tolerated warnings and informational gaps, ordered by severity and number of affected targets.
- **Per-target tables** — each enterprise, org, and repository with its maturity level, per-pillar scores, and expandable failing checks.

Maturity levels follow GitHub's start → mature → advance progression. A target's level is the lowest blocked tier under the configured maturity policy:

| Level | Meaning |
| --- | --- |
| **Start** | Blocked by `start` checks. |
| **Mature** | No `start` blockers; blocked by `mature` checks. |
| **Advance** | No `start` or `mature` blockers; `advance` checks are the remaining roadmap. |

Failing or unknown `error` checks always gate a level (after any per-target severity override in `--config`). By default, failing or unknown `warn` checks also gate a level; `info` checks never do. Rules without `maturity_level` do not affect it.

Use the top-level `maturity` block in `--config` to allow up to `max_warnings_per_tier` failing/unknown warnings per target in each tier, or set `ignore_warnings: true` with an allowance of zero to ignore warning blockers entirely. See [maturity warning allowance](../guides/configuration.md#maturity-warning-allowance).

All findings still count toward scores and appear in the roadmap. Tolerating a warning does not make it a pass. The results-completeness figure remains the share of checks postura could read; unknown results remain unknown. Maturity policy does not change CLI exit gates.

## Exit behavior

`run` defaults to `--fail-on error`; `eval` defaults to `--fail-on none`. Select `error`, `warn`, `info`, or `none`. `--fail-on-unknown` includes unknown findings at or above the selected threshold in the gate. A rule `fail` and a collection permission diagnostic are different outcomes; use `permissions check` to investigate the latter.