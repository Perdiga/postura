package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Perdiga/postura/internal/rules"
)

func TestRunConfigTemplateEnterpriseOnly(t *testing.T) {
	oldRules, oldOutput, oldEnterprise := configTemplateRules, configTemplateOutput, configTemplateEnterprise
	oldOrgs, oldAsOf, oldMaxAge := configTemplateOrgs, configTemplateAsOf, configTemplateMaxAge
	t.Cleanup(func() {
		configTemplateRules, configTemplateOutput, configTemplateEnterprise = oldRules, oldOutput, oldEnterprise
		configTemplateOrgs, configTemplateAsOf, configTemplateMaxAge = oldOrgs, oldAsOf, oldMaxAge
	})

	configTemplateRules = filepath.Join("..", "examples", "rules.yaml")
	configTemplateOutput = filepath.Join(t.TempDir(), "config.yaml")
	configTemplateEnterprise = "acme-inc"
	configTemplateOrgs = nil
	configTemplateAsOf = "2026-10-13"
	configTemplateMaxAge = 180

	var stdout, stderr bytes.Buffer
	configTemplateCmd.SetOut(&stdout)
	configTemplateCmd.SetErr(&stderr)
	t.Cleanup(func() {
		configTemplateCmd.SetOut(os.Stdout)
		configTemplateCmd.SetErr(os.Stderr)
	})
	if err := runConfigTemplate(configTemplateCmd, nil); err != nil {
		t.Fatal(err)
	}
	if stdout.Len() != 0 || !strings.Contains(stderr.String(), configTemplateOutput) {
		t.Fatalf("expected path on stderr only, stdout=%q stderr=%q", stdout.String(), stderr.String())
	}
	data, err := os.ReadFile(configTemplateOutput)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "ENT-6:") || strings.Contains(string(data), "ORG-11:") {
		t.Fatalf("enterprise-only template has wrong rule entries:\n%s", data)
	}
	parsed, err := rules.LoadConfig(data)
	if err == nil {
		t.Fatal("blank template should require completion before config validation")
	}
	if parsed != nil {
		t.Fatal("invalid blank template returned a config")
	}
	info, err := os.Stat(configTemplateOutput)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm()&0o077 != 0 {
		t.Errorf("generated config permissions = %o, want private file permissions", info.Mode().Perm())
	}
	if err := runConfigTemplate(configTemplateCmd, nil); err == nil {
		t.Fatal("template generator should refuse to overwrite an existing file")
	}
}

func TestRunConfigTemplateExplicitRepoNeedsNoToken(t *testing.T) {
	oldRules, oldOutput, oldEnterprise := configTemplateRules, configTemplateOutput, configTemplateEnterprise
	oldOrgs, oldAsOf, oldMaxAge := configTemplateOrgs, configTemplateAsOf, configTemplateMaxAge
	t.Cleanup(func() {
		configTemplateRules, configTemplateOutput, configTemplateEnterprise = oldRules, oldOutput, oldEnterprise
		configTemplateOrgs, configTemplateAsOf, configTemplateMaxAge = oldOrgs, oldAsOf, oldMaxAge
	})
	t.Setenv("GITHUB_TOKEN", "")
	t.Setenv("GH_TOKEN", "")
	t.Setenv("PATH", t.TempDir())

	configTemplateRules = filepath.Join("..", "examples", "rules.yaml")
	configTemplateOutput = filepath.Join(t.TempDir(), "repo-config.yaml")
	configTemplateEnterprise = ""
	configTemplateOrgs = nil
	configTemplateAsOf = "2026-10-13"
	configTemplateMaxAge = 180
	if err := runConfigTemplate(configTemplateCmd, []string{"acme/api"}); err != nil {
		t.Fatalf("explicit repo template should not need a token: %v", err)
	}
	data, err := os.ReadFile(configTemplateOutput)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "acme/api:") || !strings.Contains(string(data), "REPO-30:") {
		t.Fatalf("repo assessment entries missing:\n%s", data)
	}
}
