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

SARIF levels map from rule severity: `error` becomes `error`, `warn` becomes `warning`, and `info` becomes `note`.

## Exit behavior

`run` defaults to `--fail-on error`; `eval` defaults to `--fail-on none`. Select `error`, `warn`, `info`, or `none`. `--fail-on-unknown` includes unknown findings at or above the selected threshold in the gate. A rule `fail` and a collection permission diagnostic are different outcomes; use `permissions check` to investigate the latter.