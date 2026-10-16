---
title: Rules and CEL
description: Create and tune posture rulesets.
---

# Rules and CEL

A ruleset is YAML data supplied with `--rules`. postura does not embed a baseline. The full, runnable baseline is maintained in [`examples/rules.yaml`](https://github.com/Perdiga/postura/blob/main/examples/rules.yaml).

## Rule shape

```yaml
- id: ORG-2
  scope: org
  severity: warn
  title: Owner/admin count is proportionate to org size
  principle_domain: Governance
  maturity_level: mature
  config:
    max_admins: 5
    max_admin_ratio: 0.25
  pass: >-
    admin_count <= cfg.max_admins &&
    (member_count == 0 || double(admin_count) / double(member_count) <= cfg.max_admin_ratio)
  fix_hint: "Review org owners and downgrade unnecessary owners"
```

Rules have a unique `id`, a `scope` (`enterprise`, `org`, or `repo`), a `severity` (`error`, `warn`, or `info`), a title, and a CEL `pass` expression. Optional fields include:

- `principle_domain`: one of the five GitHub Well-Architected pillars: `Productivity`, `Collaboration`, `Application Security`, `Governance`, or `Architecture`.
- `maturity_level`: GitHub's maturity stage for the practice: `start`, `mature`, or `advance`.
- `config`: defaults exposed to CEL as `cfg`; values can be overridden per target.
- `applies_when`: CEL expression that limits the rule to matching targets, such as public repos.
- `enabled: false`: opt-in rule, disabled unless enabled in config.
- `assessment: true`: declares that this rule consumes a human assessment from config.
- `question`: for assessment rules, the question a reviewer should answer; `config template` copies it, with the title, into each generated entry.
- `fix_hint`: remediation text included with findings and rule inspection.

## Domain and maturity mapping

The shipped ruleset assigns every rule to the pillar it most directly supports. The five domain names follow the [GitHub Well-Architected pillars](https://learn.github.com/well-architected/overview/layers). A rule may support other pillars too; the tag is a primary classification, not an exclusive relationship.

Maturity levels follow GitHub's start → mature → advance progression and indicate the sophistication of the practice being checked:

| Level | Interpretation |
| --- | --- |
| `start` | A basic protective setting or repeatable practice is in place. |
| `mature` | The practice is standardized, reviewed, or enforced across targets. |
| `advance` | The practice uses risk-based controls, continuous monitoring, or tested response/automation. |

These values do not affect CEL evaluation, severity, or pass/fail status; the [HTML report](../reference/reports.md#html-maturity-report) uses them to score maturity.

## CEL expressions

Fact fields are bound as top-level variables, alongside the `cfg` map. For example, a repo rule can evaluate `visibility == "public"`, while a rule with a `max_admins` knob can use `admin_count <= cfg.max_admins`. Expressions are compiled when the ruleset loads; unknown fields and invalid CEL fail before evaluation.

`pass` must evaluate to true for compliance. If a required fact was not readable, the finding is unknown rather than a pass. Rules can use `applies_when` to skip targets where the check does not apply.

## Rule resolution

Built-in values are defaults, not immutable policy. A config file can override `enabled`, `severity`, and any `cfg` value at enterprise, org, or repo level. See [configuration](./configuration.md) for precedence and examples.

## Inspect and validate

```bash
postura rules list --rules rules.yaml
postura rules explain --rules rules.yaml ORG-2
postura eval --rules rules.yaml facts.json
```

Keep rules in version control and validate them with `eval` against representative fact snapshots before using `run` as a CI gate. The [fact reference](../reference/facts.md) lists the CEL variables available at each scope.