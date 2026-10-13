---
title: CI integration
description: Gate GitHub posture checks and publish SARIF results.
---

# CI integration

`run` is deterministic for a given collected state, ruleset, and config. It can gate on findings and emit SARIF for GitHub code scanning or other SARIF consumers.

## Gate on severity

```bash
postura run \
  --org acme \
  --fail-on warn \
  --rules rules.yaml
```

`--fail-on` accepts `error`, `warn`, `info`, or `none`. The `run` default is `error`; `eval` defaults to `none`. A failure at or above the chosen severity exits non-zero.

Unknown means posture could not make a reliable determination, often because the token could not read a fact. To treat unknown findings at or above the selected severity as gate failures, add `--fail-on-unknown`. Permission/feature availability can also be checked separately with `postura permissions check`.

## Emit SARIF

```bash
postura run \
  --org acme \
  --fail-on warn \
  --format sarif \
  --rules rules.yaml > postura.sarif
```

Supported formats are Markdown (`md`), JSON (`json`), and SARIF (`sarif`). Rule severities map to SARIF levels: `error` to `error`, `warn` to `warning`, and `info` to `note`.

## Separate collection from evaluation

For controlled collection or offline review, save facts first and evaluate them later:

```bash
postura collect --org acme > facts.json
postura eval --rules rules.yaml --fail-on warn facts.json
```

The fact snapshot freezes computed ages at collection time. Re-evaluating the same facts does not refresh those values, preserving reproducibility. Store fact snapshots securely: they include organization and repository security metadata.

## Scope carefully

Use `--scope org,repos` or repeat `--scope` to limit what is collected. Enterprise scope checks enterprise account settings only; it does not discover organizations beneath the enterprise. Name orgs explicitly with `--org` when auditing them.