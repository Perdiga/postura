---
title: Rules and CEL
description: Create and tune posture rulesets.
---

# Rules and CEL

A ruleset is YAML data supplied with `--rules`. postura does not embed a baseline. The full, runnable baseline is maintained in [`examples/rules.yaml`](https://github.com/jackchuka/postura/blob/main/examples/rules.yaml).

## Rule shape

```yaml
- id: ORG-2
  scope: org
  severity: warn
  title: Owner/admin count is proportionate to org size
  config:
    max_admins: 5
    max_admin_ratio: 0.25
  pass: >-
    admin_count <= cfg.max_admins &&
    (member_count == 0 || double(admin_count) / double(member_count) <= cfg.max_admin_ratio)
  fix_hint: "Review org owners and downgrade unnecessary owners"
```

Rules have a unique `id`, a `scope` (`enterprise`, `org`, or `repo`), a `severity` (`error`, `warn`, or `info`), a title, and a CEL `pass` expression. Optional fields include:

- `config`: defaults exposed to CEL as `cfg`; values can be overridden per target.
- `applies_when`: CEL expression that limits the rule to matching targets, such as public repos.
- `enabled: false`: opt-in rule, disabled unless enabled in config.
- `assessment: true`: declares that this rule consumes a human assessment from config.
- `fix_hint`: remediation text included with findings and rule inspection.

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