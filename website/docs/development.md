---
title: Development
description: Build, test, and contribute to postura.
---

# Development

## Go CLI

The Go module targets Go 1.26.5. Run the test suite from the repository root:

```bash
go test ./...
```

Build and try the CLI with a local ruleset:

```bash
go build -o postura .
./postura rules list --rules examples/rules.yaml
```

## Documentation site

The Docusaurus application is in `website/` and requires Node.js 20 or later:

```bash
cd website
npm install
npm start
```

Build the static site and catch broken links:

```bash
npm run build
```

The site is configured for GitHub Pages at `https://perdiga.github.io/postura/`. Documentation examples should stay aligned with the Go CLI and the canonical YAML files in the repository's `examples/` directory.

## Contribution notes

- Keep collection read-only and rule evaluation deterministic for a fixed fact snapshot.
- Add or update tests when changing fact fields, rule behavior, or report contracts.
- When changing collector facts, update the schema and the fact reference together.
- Do not encode organization-specific policy as built-in rules; keep policy in user-supplied rulesets/configuration.