# postura — data collection steps

The client only runs `collect`, which writes a JSON file with the settings read from GitHub. The evaluation and report are produced on our side.

The client only needs the executable for their OS (for example `postura-windows-amd64.exe`). No `rules.yaml` or `config.yaml` is required.

## Client steps

### 1. Create the token (classic PAT)

GitHub → **Settings → Developer settings → Personal access tokens → Tokens (classic) → Generate new token (classic)**.

Give it a short expiry (for example 7 days) and select these scopes:

| Scope | Used for |
| --- | --- |
| `repo` | Repositories, branch protection, workflows, files |
| `admin:org` | Org settings, owners, roles, Actions policy, runner groups (read only, but some endpoints require this scope) |
| `security_events` | Dependabot alerts and security settings |
| `read:enterprise` | Enterprise settings (only if auditing the enterprise) |

- If the org uses **SAML SSO**, click **Configure SSO → Authorize** on the token for each org.
- The token owner must be an **org owner** (and **enterprise owner** if auditing the enterprise).

### 2. Open Command Prompt (cmd) in the executable's folder

```bat
cd C:\postura
set GITHUB_TOKEN=ghp_YOUR_TOKEN_HERE
```

> Use **cmd**, not PowerShell: PowerShell 5.1's `>` writes UTF-16 and the JSON becomes unreadable. Don't put spaces around `=` or quotes around the token.

### 3. Check permissions

```bat
postura-windows-amd64.exe permissions check --enterprise ENTERPRISE-SLUG --org ORG-NAME
```

Fix the token's scopes or SSO authorization if anything is reported as unreadable.

> If Windows shows "Windows protected your PC", click **More info → Run anyway** (the executable is not code-signed).

### 4. Collect the data

```bat
postura-windows-amd64.exe collect --enterprise ENTERPRISE-SLUG --org ORG-NAME > facts.json
```

- Several orgs: `--org org-a,org-b`
- Specific repositories only: `postura-windows-amd64.exe collect --org ORG-NAME ORG-NAME/repo1 ORG-NAME/repo2 > facts.json`
- No enterprise access: drop `--enterprise ...`
- Large orgs can take a few minutes.

### 5. Send `facts.json` and revoke the token

Send the file over a secure channel. It contains **no secrets or source code**, but it does include usernames, collaborators, security settings, and alert details. Afterwards, **revoke the token** under GitHub → Settings → Developer settings → Tokens.

## Our side

### Generate the report

```bash
postura eval facts.json --rules examples/rules.yaml --config config.yaml --format html --fail-on none > report.html
```

### Manual assessments (optional)

`facts.json` only contains API data, so the assessment questions are handled with the client. Generate the template without a token by listing the repositories explicitly (repo names are in the `name` field of each repo in `facts.json`):

```bash
postura config template --rules examples/rules.yaml --enterprise SLUG --org ORG ORG/repo1 ORG/repo2 --output config.yaml
```

- Fill in `config.yaml` with the client (for example in a review meeting): `status`, `evidence`, `reviewed_at` (`YYYY-MM-DD`), `reviewed_by`.
- Remove entries you won't answer; blank entries are rejected, and removed ones show as "unknown" in the report.
- Passing only `--org` without repositories makes the command look them up on GitHub, which needs a token with access to that org.
- The command won't overwrite an existing file.
