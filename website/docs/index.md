---
slug: /
title: Documentation
description: Guides and reference for posture, the read-only GitHub security auditor.
---

import Link from '@docusaurus/Link';

# posture documentation

posture audits GitHub enterprise settings, organizations, and repositories against a security baseline you own. Rules are YAML data evaluated with CEL; policy thresholds can differ by target. Collection is read-only, and facts can be evaluated offline for repeatable CI checks.

<div className="home-grid">
  <Link to="/quickstart"><strong>Start here</strong><span>Install posture, authenticate, and run your first audit.</span></Link>
  <Link to="/guides/rules"><strong>Write rules</strong><span>Define scopes, severities, CEL expressions, and config knobs.</span></Link>
  <Link to="/guides/configuration"><strong>Tune policy</strong><span>Layer target overrides and record human assessments.</span></Link>
  <Link to="/guides/commands"><strong>Use the CLI</strong><span>Run, collect, evaluate, inspect rules, and generate templates.</span></Link>
  <Link to="/reference/facts"><strong>Fact reference</strong><span>See every fact exposed to enterprise, org, and repo rules.</span></Link>
  <Link to="/reference/permissions"><strong>Access and reports</strong><span>Understand token access, unknown results, and output formats.</span></Link>
</div>

## The audit pipeline

```text
GitHub REST + GraphQL -> collect facts -> evaluate CEL rules -> Markdown / JSON / SARIF
                          read-only       deterministic          exit-code gate
```

`postura run` performs all three stages. `collect` writes a JSON snapshot, and `eval` applies rules to that snapshot without network access. postura ships no default ruleset: begin with the [example baseline](https://github.com/Perdiga/postura/blob/main/examples/rules.yaml) and make the policy yours.

## Project links

- [Quickstart](./quickstart.md)
- [Command guide](./guides/commands.md)
- [Rules and CEL](./guides/rules.md)
- [Policy configuration and assessments](./guides/configuration.md)
- [Fact fields](./reference/facts.md)
- [Example ruleset](https://github.com/Perdiga/postura/blob/main/examples/rules.yaml)
- [Example config](https://github.com/Perdiga/postura/blob/main/examples/config.yaml)