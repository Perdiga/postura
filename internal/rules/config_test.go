package rules

import (
	"os"
	"testing"
)

func boolPtr(b bool) *bool { return &b }

func TestLoadConfigMaturity(t *testing.T) {
	for _, tt := range []struct {
		name    string
		data    string
		want    MaturityPolicy
		wantErr bool
	}{
		{name: "default"},
		{name: "allowance", data: "maturity:\n  max_warnings_per_tier: 2\n", want: MaturityPolicy{MaxWarningsPerTier: 2}},
		{name: "ignore", data: "maturity:\n  ignore_warnings: true\n", want: MaturityPolicy{IgnoreWarnings: true}},
		{name: "negative", data: "maturity:\n  max_warnings_per_tier: -1\n", wantErr: true},
		{name: "fractional", data: "maturity:\n  max_warnings_per_tier: 1.5\n", wantErr: true},
		{name: "invalid boolean", data: "maturity:\n  ignore_warnings: invalid\n", wantErr: true},
		{name: "unknown setting", data: "maturity:\n  max_warning_per_tier: 2\n", wantErr: true},
		{name: "not mapping", data: "maturity: true\n", wantErr: true},
		{name: "duplicate setting", data: "maturity:\n  ignore_warnings: false\n  ignore_warnings: true\n", wantErr: true},
		{name: "conflicting", data: "maturity:\n  max_warnings_per_tier: 2\n  ignore_warnings: true\n", wantErr: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := LoadConfig([]byte(tt.data))
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected invalid maturity config error")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if cfg.Maturity != tt.want {
				t.Fatalf("maturity = %+v, want %+v", cfg.Maturity, tt.want)
			}
		})
	}
}

func TestResolvePrecedence(t *testing.T) {
	r := Rule{ID: "ORG-2", Severity: "warn", Config: map[string]any{"max_admins": 5, "max_admin_ratio": 0.25}}
	c := &Config{
		Enterprises: map[string]map[string]map[string]any{
			"acme": {"ORG-2": {"max_admins": 6}},
		},
		Orgs: map[string]map[string]map[string]any{
			"acme-inc": {"ORG-2": {"max_admins": 8, "severity": "error"}},
		},
	}

	// org override beats enterprise beats rule default; untouched knob keeps default.
	got := c.Resolve(r, "acme", "acme-inc", "")
	if got.Cfg["max_admins"] != 8 {
		t.Errorf("max_admins: got %v, want 8", got.Cfg["max_admins"])
	}
	if got.Cfg["max_admin_ratio"] != 0.25 {
		t.Errorf("max_admin_ratio: got %v, want 0.25 (default)", got.Cfg["max_admin_ratio"])
	}
	if got.Severity != "error" {
		t.Errorf("severity: got %q, want error", got.Severity)
	}

	// enterprise-only override applies when no org layer present.
	got = c.Resolve(r, "acme", "other-org", "")
	if got.Cfg["max_admins"] != 6 {
		t.Errorf("enterprise layer: got %v, want 6", got.Cfg["max_admins"])
	}
}

func TestResolveEnabled(t *testing.T) {
	off := Rule{ID: "REPO-8", Severity: "info", Enabled: boolPtr(false), Pass: "true"}
	// off by default
	if c := (&Config{}); c.Resolve(off, "", "", "").Enabled {
		t.Error("REPO-8 should be disabled by default")
	}
	// re-enabled per repo
	c := &Config{Repos: map[string]map[string]map[string]any{"o/r": {"REPO-8": {"enabled": true}}}}
	if !c.Resolve(off, "", "", "o/r").Enabled {
		t.Error("REPO-8 should be enabled for o/r")
	}
}

func TestLoadConfigAssessments(t *testing.T) {
	data := []byte(`assessments:
  as_of: "2026-10-13"
  max_age_days: 180
  orgs:
    acme:
      ORG-H1:
        status: compliant
        evidence: "Restore test documented"
        reviewed_at: "2026-10-01"
        reviewed_by: "security-team"
`)
	c, err := LoadConfig(data)
	if err != nil {
		t.Fatal(err)
	}
	r := Rule{ID: "ORG-H1", Severity: "warn"}
	got := c.Resolve(r, "", "acme", "")
	assessment, ok := got.Cfg["assessment"].(map[string]any)
	if !ok {
		t.Fatalf("assessment config = %T, want map[string]any", got.Cfg["assessment"])
	}
	if assessment["status"] != "compliant" || assessment["reviewed_by"] != "security-team" {
		t.Errorf("resolved assessment = %+v", assessment)
	}
	if assessment["stale"] != false {
		t.Errorf("recent assessment stale = %v, want false", assessment["stale"])
	}
}

func TestLoadConfigRejectsInvalidAssessment(t *testing.T) {
	cases := map[string]string{
		"invalid status":   `status: approved`,
		"missing evidence": "status: compliant\nreviewed_at: '2026-10-01'\nreviewed_by: team",
		"invalid date":     "status: compliant\nevidence: report\nreviewed_at: yesterday\nreviewed_by: team",
	}
	for name, fields := range cases {
		t.Run(name, func(t *testing.T) {
			data := []byte("assessments:\n  orgs:\n    acme:\n      ORG-H1:\n        " + fields + "\n")
			if _, err := LoadConfig(data); err == nil {
				t.Fatal("invalid assessment should fail config loading")
			}
		})
	}
}

func TestAssessmentTargetSpecificity(t *testing.T) {
	c := &Config{Assessments: AssessmentConfig{
		Orgs: map[string]map[string]AssessmentInput{
			"acme": {"REPO-H1": {Status: "partial", Evidence: "org evidence", ReviewedAt: "2026-10-01", ReviewedBy: "org"}},
		},
		Repos: map[string]map[string]AssessmentInput{
			"acme/payments": {"REPO-H1": {Status: "compliant", Evidence: "repo evidence", ReviewedAt: "2026-10-02", ReviewedBy: "repo"}},
		},
	}}
	got := c.Resolve(Rule{ID: "REPO-H1"}, "", "acme", "acme/payments")
	assessment := got.Cfg["assessment"].(map[string]any)
	if assessment["status"] != "compliant" || assessment["reviewed_by"] != "repo" {
		t.Errorf("repo assessment should override org assessment, got %+v", assessment)
	}
}

func TestAssessmentFreshness(t *testing.T) {
	data := []byte(`assessments:
  as_of: "2026-10-13"
  max_age_days: 30
  orgs:
    acme:
      ORG-H1:
        status: compliant
        evidence: "Restore plan tested"
        reviewed_at: "2026-08-01"
        reviewed_by: "security"
`)
	c, err := LoadConfig(data)
	if err != nil {
		t.Fatal(err)
	}
	got := c.Resolve(Rule{ID: "ORG-H1"}, "", "acme", "")
	assessment := got.Cfg["assessment"].(map[string]any)
	if assessment["stale"] != true {
		t.Fatalf("old assessment stale = %v, want true", assessment["stale"])
	}
}

func TestExampleConfigLoads(t *testing.T) {
	data, err := os.ReadFile("../../examples/config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := LoadConfig(data); err != nil {
		t.Fatalf("load examples/config.yaml: %v", err)
	}
}
