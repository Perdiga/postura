package report

import (
	"strings"
	"testing"

	"github.com/Perdiga/postura/internal/engine"
)

func TestShouldFail(t *testing.T) {
	fs := []engine.Finding{
		{ID: "REPO-4", Severity: "warn", Status: engine.StatusFail},
		{ID: "REPO-5", Severity: "info", Status: engine.StatusUnknown},
	}
	if ShouldFail(fs, "error", false) {
		t.Error("no error fail present; should not fail on error")
	}
	if !ShouldFail(fs, "warn", false) {
		t.Error("a warn fail present; should fail on warn")
	}
	if ShouldFail(fs, "none", false) {
		t.Error("none threshold never fails")
	}

	// failOnUnknown, isolated to an info-severity unknown (no fail present).
	unk := []engine.Finding{{ID: "REPO-5", Severity: "info", Status: engine.StatusUnknown}}
	if ShouldFail(unk, "error", true) {
		t.Error("info unknown is below error threshold; should not fail")
	}
	if !ShouldFail(unk, "info", true) {
		t.Error("info unknown at info threshold with failOnUnknown; should fail")
	}
	if ShouldFail(unk, "info", false) {
		t.Error("unknown without failOnUnknown never gates")
	}
	if ShouldFail(unk, "none", true) {
		t.Error("none threshold never fails, even with failOnUnknown")
	}
}

func TestMarkdownDropsPassAndOrders(t *testing.T) {
	fs := []engine.Finding{
		{ID: "REPO-4", Severity: "warn", Title: "deps", Target: "o/a", Status: engine.StatusFail},
		{ID: "REPO-1", Severity: "error", Title: "pr", Target: "o/a", Status: engine.StatusFail},
		{ID: "REPO-6", Severity: "info", Title: "license", Target: "o/a", Status: engine.StatusPass},
	}
	md := Markdown(fs, "1 repos")

	// error rule sorts before warn.
	if strings.Index(md, "REPO-1") > strings.Index(md, "REPO-4") {
		t.Error("REPO-1 (error) should sort before REPO-4 (warn)")
	}
	// an all-pass rule shows the summary line, not a table row.
	if !strings.Contains(md, "✅ all 1 pass") {
		t.Error("REPO-6 is all-pass; expected '✅ all 1 pass'")
	}
	if strings.Contains(md, "| o/a | ✅") {
		t.Error("pass rows must be dropped from tables")
	}
}

func TestMarkdownIncludesPassingAssessmentEvidence(t *testing.T) {
	fs := []engine.Finding{{
		ID: "ORG-11", Severity: "warn", Title: "Recovery plan", Target: "acme",
		Status: engine.StatusPass, Notes: "Assessment: compliant; Evidence: restore test passed",
	}}
	md := Markdown(fs, "org")
	if !strings.Contains(md, "| acme | ✅ pass | Assessment: compliant; Evidence: restore test passed |") {
		t.Fatalf("passing assessment evidence missing from Markdown:\n%s", md)
	}
}

func TestMarkdownEscapesAssessmentEvidence(t *testing.T) {
	fs := []engine.Finding{{
		ID: "ORG-11", Severity: "warn", Title: "Recovery plan", Target: "acme",
		Status: engine.StatusPass, Notes: "Evidence: rollback | restore\nverified",
	}}
	md := Markdown(fs, "org")
	if !strings.Contains(md, `rollback \| restore verified`) {
		t.Fatalf("assessment evidence breaks Markdown table formatting:\n%s", md)
	}
}

func TestSARIFIncludesPassingAssessmentEvidence(t *testing.T) {
	fs := []engine.Finding{{
		ID: "ORG-11", Scope: "org", Severity: "warn", Title: "Recovery plan", Target: "acme",
		Status: engine.StatusPass, Notes: "Assessment: compliant; Evidence: restore test passed",
	}}
	data, err := SARIF(fs)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "restore test passed") {
		t.Fatalf("passing assessment evidence missing from SARIF: %s", data)
	}
}

func TestSARIFProjectURL(t *testing.T) {
	data, err := SARIF(nil)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), `"informationUri": "https://github.com/Perdiga/postura"`) {
		t.Fatalf("SARIF must link to the fork: %s", data)
	}
}
