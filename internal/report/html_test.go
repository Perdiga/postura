package report

import (
	"strings"
	"testing"
	"time"

	"github.com/jackchuka/postura/internal/engine"
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
		if got := maturityLevel(c.fs); got != c.want {
			t.Errorf("%s: got %d, want %d", c.name, got, c.want)
		}
	}
}

func TestHTMLReport(t *testing.T) {
	fs := []engine.Finding{
		{ID: "ENT-1", Scope: "enterprise", Severity: "error", Title: "SSO", Target: "acme-ent", Status: engine.StatusPass, PrincipleDomain: "Governance", MaturityLevel: "start"},
		{ID: "ORG-1", Scope: "org", Severity: "error", Title: "2FA", Target: "acme", Status: engine.StatusFail, PrincipleDomain: "Governance", MaturityLevel: "start", FixHint: "require 2FA"},
		{ID: "REPO-1", Scope: "repo", Severity: "warn", Title: "<script>alert(1)</script>", Target: "acme/a", Status: engine.StatusUnknown, PrincipleDomain: "Application Security", MaturityLevel: "mature"},
		{ID: "REPO-1", Scope: "repo", Severity: "warn", Title: "<script>alert(1)</script>", Target: "acme/b", Status: engine.StatusPass, PrincipleDomain: "Application Security", MaturityLevel: "mature"},
	}
	out, err := HTML(fs, "enterprise + org + 2 repos", time.Date(2026, 1, 2, 3, 4, 0, 0, time.UTC))
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
	if strings.Index(html, "Start checks") > strings.Index(html, "Mature checks") {
		t.Error("roadmap tiers should be ordered start before mature")
	}
}
