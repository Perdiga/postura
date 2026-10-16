---
title: Quickstart
description: Install postura and run your first GitHub audit.
sidebar_position: 1
---

# Quickstart

## Requirements

- Go 1.26.5 or later to install from source with `go install`.
- A GitHub token in `GITHUB_TOKEN` or `GH_TOKEN`, or an authenticated GitHub CLI (`gh auth token` is used as a fallback).
- A ruleset YAML file. postura intentionally ships without built-in policy; the repository's [`examples/rules.yaml`](https://github.com/Perdiga/postura/blob/main/examples/rules.yaml) is a starting baseline.

## Install

Install the latest version:

```bash
go install github.com/Perdiga/postura@latest
```

Or download a prebuilt Linux, macOS, or Windows binary for amd64 or arm64 from [GitHub Releases](https://github.com/Perdiga/postura/releases).

## Authenticate

Export a token before running an audit:

```bash
export GITHUB_TOKEN=YOUR_TOKEN
```

The token needs read access for the facts you want to collect. Missing access is reported as unknown, not as a pass. Use the [permissions reference](./reference/permissions.md) to check requirements and diagnose incomplete collection.

## Run an audit

```bash
postura run --org acme --rules examples/rules.yaml
```

With an organization and no repository arguments, postura audits the org and discovers its non-forked repositories. To include enterprise settings:

```bash
postura run --org acme --enterprise acme-inc --rules examples/rules.yaml
```

Use `--scope` to limit collection. For example, org-level checks without the repo sweep:

```bash
postura run --org acme --scope org --rules examples/rules.yaml
```

Run posture against selected repositories (the owner is inferred as the org):

```bash
postura run acme/web acme/api --rules examples/rules.yaml
```

## Set a CI threshold

By default, `run` exits non-zero when an error-severity rule fails. Make warnings gate too, and emit SARIF for code scanning:

```bash
postura run --org acme --fail-on warn --format sarif --rules examples/rules.yaml > postura.sarif
```

To also fail when collection is incomplete for a finding at or above the threshold, add `--fail-on-unknown`. See the [CI guide](./guides/ci.md).

## Next steps

- Learn the [command flags](./guides/commands.md).
- Adapt the [ruleset](./guides/rules.md) and [per-target configuration](./guides/configuration.md).
- Read about [facts and their scopes](./reference/facts.md).