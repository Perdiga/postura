---
title: Commands and flags
description: Complete command reference for the postura CLI.
---

# Commands and flags

Every command accepts `--help`. `--rules` is required for `run`, `eval`, `rules list`, and `rules explain`. `collect` gathers facts without evaluating a ruleset.

## `run`

Collect facts from GitHub, evaluate a ruleset, write findings, and choose the process exit code:

```bash
postura run --org acme --enterprise acme-inc --rules rules.yaml
```

Syntax: `postura run --rules RULES [owner/repo ...]`. Without repository arguments, each named org expands to all non-forked repositories. `--org` is repeatable or comma-separated; when repository arguments are supplied, their owner is inferred as the org.

| Flag | Default | Meaning |
| --- | --- | --- |
| `--rules` | required | Ruleset YAML. |
| `--config` | none | Policy config with target overrides and/or assessments. |
| `--org` | inferred or none | Organization(s) to audit; repeatable or comma-separated. |
| `--enterprise` | none | Enterprise slug to audit. Requires enterprise scope access. |
| `--scope` | all applicable | `enterprise`, `org`, and/or `repos`; repeatable or comma-separated. |
| `--format` | `md` | `md`, `json`, or `sarif`. |
| `--fail-on` | `error` | Exit non-zero for a fail at or above `error`, `warn`, `info`, or `none`. |
| `--fail-on-unknown` | `false` | Also fail for unknown findings at or above `--fail-on`. |

Each selected scope needs a target: enterprise scope requires `--enterprise`; org or repo scopes require an org or repository target. Enterprise audits cover account-level settings only and do not enumerate member orgs.

## `collect`

Write collected facts as JSON for caching or offline evaluation:

```bash
postura collect --org acme > facts.json
postura collect --enterprise acme-inc --scope enterprise > enterprise.json
```

It accepts `--org`, `--enterprise`, and `--scope` with the same target semantics as `run`. Collection is the only networked stage.

## `eval`

Evaluate a saved facts file without contacting GitHub:

```bash
postura eval --rules rules.yaml facts.json
```

`eval` accepts `--config`, `--format`, and `--fail-on`. Its default `--fail-on` is `none`; target identity comes from the facts file, so it does not take `--org`, `--enterprise`, or `--scope`.

## `rules`

List rules or inspect one rule's expressions, settings, and remediation hint:

```bash
postura rules list --rules rules.yaml
postura rules explain --rules rules.yaml REPO-1
```

## `config template`

Generate blank inputs for every rule marked `assessment: true` in the selected targets:

```bash
postura config template --rules rules.yaml --org acme --output assessments.yaml
postura config template --rules rules.yaml --enterprise acme-inc --org acme --as-of 2026-10-13
```

Flags: `--rules` is required; `--output` defaults to `config.yaml` (`-` writes to stdout); `--enterprise` and repeatable `--org` select targets; `--as-of` defaults to today's UTC date; `--max-age-days` defaults to 180. If an org is selected without repo arguments, it discovers non-forked repositories and therefore needs a token. The command refuses to overwrite an existing output file.

The output contains human assessment inputs only. It does not generate policy overrides or collected GitHub facts. Complete the required fields or remove unused entries before passing it as `--config`.

## `permissions check`

Probe collection capability groups and print missing facts with access guidance:

```bash
postura permissions check --org acme --enterprise acme-inc
```

Supports `--org`, `--enterprise`, and `--scope`, plus repository arguments. A non-zero result means one or more capability groups were incomplete; a missing fact can also indicate an unavailable or ineligible GitHub feature.