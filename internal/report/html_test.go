package report

import (
	"strings"
	"testing"
	"time"

	"github.com/Perdiga/postura/internal/engine"
	"github.com/Perdiga/postura/internal/rules"
)

func TestMaturityLevelIsCumulative(t *testing.T) {
	f := func(level string, s engine.Status) engine.Finding {
		return engine.Finding{MaturityLevel: level, Status: s}
	}
	cases := []struct {
		name string
		fs   []engine.Finding
		want int
	}{
		{"start fail", []engine.Finding{f("start", engine.StatusFail), f("mature", engine.StatusPass)}, 0},
		{"start unknown", []engine.Finding{f("start", engine.StatusUnknown)}, 0},
		{"mature gap", []engine.Finding{f("start", engine.StatusPass), f("mature", engine.StatusFail), f("advance", engine.StatusPass)}, 1},
		{"advance gap", []engine.Finding{f("start", engine.StatusPass), f("mature", engine.StatusPass), f("advance", engine.StatusFail)}, 2},
		{"all pass", []engine.Finding{f("start", engine.StatusPass), f("advance", engine.StatusPass)}, 2},
		{"untagged ignored", []engine.Finding{f("", engine.StatusFail)}, 2},
		{"info does not gate", []engine.Finding{{MaturityLevel: "start", Severity: "info", Status: engine.StatusFail}, f("mature", engine.StatusFail)}, 1},
	}
	for _, c := range cases {
		if got := maturityLevel(c.fs, rules.MaturityPolicy{}); got != c.want {
			t.Errorf("%s: got %d, want %d", c.name, got, c.want)
		}
	}
}

func TestMaturityWarningAllowance(t *testing.T) {
	finding := func(tier, severity string, status engine.Status) engine.Finding {
		return engine.Finding{MaturityLevel: tier, Severity: severity, Status: status}
	}
	warn := finding("start", "warn", engine.StatusFail)
	unknown := finding("start", "warn", engine.StatusUnknown)
	matureWarn := finding("mature", "warn", engine.StatusFail)
	cases := []struct {
		name   string
		fs     []engine.Finding
		policy rules.MaturityPolicy
		want   int
	}{
		{"default blocks warning", []engine.Finding{warn}, rules.MaturityPolicy{}, 0},
		{"at allowance", []engine.Finding{warn, unknown}, rules.MaturityPolicy{MaxWarningsPerTier: 2}, 2},
		{"above allowance", []engine.Finding{warn, unknown, warn}, rules.MaturityPolicy{MaxWarningsPerTier: 2}, 0},
		{"unknown warnings count", []engine.Finding{unknown}, rules.MaturityPolicy{}, 0},
		{"allowances reset per tier", []engine.Finding{warn, matureWarn}, rules.MaturityPolicy{MaxWarningsPerTier: 1}, 2},
		{"mature exceeds allowance", []engine.Finding{warn, matureWarn, matureWarn}, rules.MaturityPolicy{MaxWarningsPerTier: 1}, 1},
		{"ignore warnings", []engine.Finding{warn, unknown, matureWarn}, rules.MaturityPolicy{IgnoreWarnings: true}, 2},
		{"errors block with allowance", []engine.Finding{finding("start", "error", engine.StatusFail)}, rules.MaturityPolicy{MaxWarningsPerTier: 100}, 0},
		{"errors block ignoring warnings", []engine.Finding{warn, finding("mature", "error", engine.StatusFail)}, rules.MaturityPolicy{IgnoreWarnings: true}, 1},
		{"unknown errors block", []engine.Finding{finding("start", "error", engine.StatusUnknown)}, rules.MaturityPolicy{IgnoreWarnings: true}, 0},
		{"info never blocks", []engine.Finding{finding("start", "info", engine.StatusUnknown)}, rules.MaturityPolicy{}, 2},
		{"passing warnings do not consume allowance", []engine.Finding{warn, finding("start", "warn", engine.StatusPass)}, rules.MaturityPolicy{MaxWarningsPerTier: 1}, 2},
		{"untagged warning ignored", []engine.Finding{finding("", "warn", engine.StatusFail)}, rules.MaturityPolicy{}, 2},
		{"advance remains top level", []engine.Finding{finding("advance", "error", engine.StatusFail)}, rules.MaturityPolicy{}, 2},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := maturityLevel(tt.fs, tt.policy); got != tt.want {
				t.Fatalf("maturityLevel = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestHTMLMaturityPolicy(t *testing.T) {
	fs := []engine.Finding{
		{ID: "R1", Scope: "repo", Target: "acme/a", Severity: "warn", MaturityLevel: "start", Status: engine.StatusFail},
		{ID: "R1", Scope: "repo", Target: "acme/b", Severity: "warn", MaturityLevel: "start", Status: engine.StatusUnknown},
	}
	strict := buildHTMLReport(fs, "2 repos", time.Time{}, rules.MaturityPolicy{})
	policy := rules.MaturityPolicy{MaxWarningsPerTier: 1}
	allowed := buildHTMLReport(fs, "2 repos", time.Time{}, policy)
	if strict.Overall != allowed.Overall || allowed.Overall.Fail != 1 || allowed.Overall.Unknown != 1 {
		t.Fatal("warning policy must not change scores or unknown counts")
	}
	for _, target := range strict.Scopes[0].Targets {
		if target.Level != 0 {
			t.Fatal("default policy must block each target")
		}
	}
	for _, target := range allowed.Scopes[0].Targets {
		if target.Level != 2 {
			t.Fatal("warning allowance must apply independently per target")
		}
	}
	if len(allowed.Roadmap) != 1 || len(allowed.Roadmap[0].Gaps) != 1 {
		t.Fatal("tolerated warnings must remain in the roadmap")
	}
	for _, tt := range []struct {
		policy rules.MaturityPolicy
		text   string
	}{
		{policy, "Up to 1 failing or unknown"},
		{rules.MaturityPolicy{IgnoreWarnings: true}, "Warning checks do not block maturity."},
	} {
		out, err := HTML(fs, "2 repos", time.Time{}, tt.policy)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(out), tt.text) {
			t.Fatalf("HTML missing policy explanation %q", tt.text)
		}
	}
	if _, err := HTML(fs, "2 repos", time.Time{}, rules.MaturityPolicy{MaxWarningsPerTier: -1}); err == nil {
		t.Fatal("HTML must reject invalid maturity policy")
	}
}

func TestHTMLReport(t *testing.T) {
	fs := []engine.Finding{
		{ID: "ENT-1", Scope: "enterprise", Severity: "error", Title: "SSO", Target: "acme-ent", Status: engine.StatusPass, PrincipleDomain: "Governance", MaturityLevel: "start"},
		{ID: "ORG-1", Scope: "org", Severity: "error", Title: "2FA", Target: "acme", Status: engine.StatusFail, PrincipleDomain: "Governance", MaturityLevel: "start", FixHint: "require 2FA"},
		{ID: "REPO-1", Scope: "repo", Severity: "warn", Title: "<script>alert(1)</script>", Target: "acme/a", Status: engine.StatusUnknown, PrincipleDomain: "Application Security", MaturityLevel: "mature"},
		{ID: "REPO-1", Scope: "repo", Severity: "warn", Title: "<script>alert(1)</script>", Target: "acme/b", Status: engine.StatusPass, PrincipleDomain: "Application Security", MaturityLevel: "mature"},
	}
	out, err := HTML(fs, "enterprise + org + 2 repos", time.Date(2026, 1, 2, 3, 4, 0, 0, time.UTC), rules.MaturityPolicy{})
	if err != nil {
		t.Fatal(err)
	}
	html := string(out)
	for _, want := range []string{
		"2026-01-02 03:04 UTC",
		"https://learn.github.com/well-architected/governance/overview",
		"Start checks — fix to reach Mature",
		"Results completeness",
		"1 of 4 checks could not be read",
		"Mature checks — fix to reach Advance",
		"require 2FA",
		"Advance",                // enterprise passes everything
		"No automated checks",    // productivity/architecture have no rules
		"&lt;script&gt;alert(1)", // titles are escaped
	} {
		if !strings.Contains(html, want) {
			t.Errorf("HTML report missing %q", want)
		}
	}
	if strings.Contains(html, "<script>alert(1)") {
		t.Error("finding title was not HTML-escaped")
	}
	if strings.Contains(html, "postura permissions check") || strings.Contains(html, "supply assessments") {
		t.Error("HTML report includes removed CLI guidance")
	}
	if strings.Index(html, "Start checks") > strings.Index(html, "Mature checks") {
		t.Error("roadmap tiers should be ordered start before mature")
	}
}
