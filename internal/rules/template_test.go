package rules

import (
	"bytes"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestGenerateAssessmentTemplate(t *testing.T) {
	rs := &Ruleset{Rules: []Rule{
		{ID: "ENT-H1", Scope: ScopeEnterprise, Assessment: true},
		{ID: "ORG-H1", Scope: ScopeOrg, Assessment: true, Title: "Backups are tested", Question: "When was the last restore test?"},
		{ID: "REPO-H1", Scope: ScopeRepo, Assessment: true},
		{ID: "REPO-AUTO", Scope: ScopeRepo},
	}}
	targets := AssessmentTargets{
		Enterprises: []string{"zeta-inc"},
		Orgs:        []string{"zeta", "acme", "acme"},
		Repos:       []string{"zeta/api", "acme/web"},
	}
	first, err := GenerateAssessmentTemplate(rs, targets, "2026-10-13", 180)
	if err != nil {
		t.Fatal(err)
	}
	second, err := GenerateAssessmentTemplate(rs, targets, "2026-10-13", 180)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(first, second) {
		t.Fatal("generated template should be deterministic")
	}

	var got struct {
		Assessments AssessmentConfig `yaml:"assessments"`
	}
	if err := yaml.Unmarshal(first, &got); err != nil {
		t.Fatalf("unmarshal generated template: %v", err)
	}
	if got.Assessments.AsOf != "2026-10-13" || got.Assessments.MaxAgeDays != 180 {
		t.Errorf("assessment freshness = (%q, %d)", got.Assessments.AsOf, got.Assessments.MaxAgeDays)
	}
	if len(got.Assessments.Orgs) != 2 || len(got.Assessments.Repos) != 2 || len(got.Assessments.Enterprises) != 1 {
		t.Fatalf("generated targets = %+v", got.Assessments)
	}
	want := AssessmentInput{Title: "Backups are tested", Question: "When was the last restore test?"}
	if got.Assessments.Orgs["acme"]["ORG-H1"] != want {
		t.Errorf("org assessment should carry only rule guidance, got %+v", got.Assessments.Orgs["acme"]["ORG-H1"])
	}
	if _, exists := got.Assessments.Repos["acme/web"]["REPO-AUTO"]; exists {
		t.Error("non-assessment rules must not appear in the template")
	}
	if !bytes.Contains(first, []byte("Valid status values")) {
		t.Error("template should document allowed assessment status values")
	}
}

func TestGenerateAssessmentTemplateRejectsInvalidInput(t *testing.T) {
	rs := &Ruleset{Rules: []Rule{{ID: "ORG-H1", Scope: ScopeOrg, Assessment: true}}}
	cases := []struct {
		name    string
		targets AssessmentTargets
		asOf    string
		maxAge  int
	}{
		{name: "no targets", asOf: "2026-10-13", maxAge: 180},
		{name: "invalid date", targets: AssessmentTargets{Orgs: []string{"acme"}}, asOf: "yesterday", maxAge: 180},
		{name: "invalid freshness", targets: AssessmentTargets{Orgs: []string{"acme"}}, asOf: "2026-10-13", maxAge: 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := GenerateAssessmentTemplate(rs, tc.targets, tc.asOf, tc.maxAge); err == nil {
				t.Fatal("expected invalid input error")
			}
		})
	}
}
