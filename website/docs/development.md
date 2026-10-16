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

## Publishing a release

The **Release** workflow tests the tagged source, then uses GoReleaser to build
Linux, macOS, and Windows binaries for amd64 and arm64 and publish them as assets
on a GitHub Release in `Perdiga/postura`. GoReleaser also generates checksums and
a changelog. The workflow uses the repository's automatic `GITHUB_TOKEN`; no
personal access token is needed.

Each release includes six platform archives: Linux and macOS `.tar.gz` files and
Windows `.zip` files, each for amd64 and arm64. Every archive contains the
platform binary (`postura.exe` on Windows), `LICENSE`, `README.md`, and
`examples/rules.yaml` and `examples/config.yaml`. The examples are also attached
as standalone `rules.yaml` and `config.yaml` release assets so they can be
downloaded without a platform archive. Copy and customize the examples for your
organization before using them.

After committing and pushing the changes you want to release, create and push a
new version tag pointing to that commit:

```bash
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0
```

Use the next unused version for your fork. Tags must follow the `vMAJOR.MINOR.PATCH`
format, optionally with a prerelease suffix such as `v0.1.0-rc.1`.

You can also start the workflow under **Actions > Release > Run workflow**.
Select an existing version tag in the branch/tag selector; selecting a branch
is rejected. The workflow does not create tags.

Download published archives from
[GitHub Releases](https://github.com/Perdiga/postura/releases). For untagged
development builds, use the separate **Build binaries** workflow and download
its run artifacts instead.

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