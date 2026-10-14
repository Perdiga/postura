package rules

import "testing"

func TestLoadRuleMetadata(t *testing.T) {
	rs, err := Load([]byte(`rules:
  - id: ORG-1
    scope: org
    severity: error
    title: Require two-factor authentication
    principle_domain: Governance
    maturity_level: start
    pass: two_factor_requirement_enabled == true
`))
	if err != nil {
		t.Fatal(err)
	}
	if got := rs.Rules[0].PrincipleDomain; got != DomainGovernance {
		t.Errorf("principle domain = %q, want %q", got, DomainGovernance)
	}
	if got := rs.Rules[0].MaturityLevel; got != MaturityStart {
		t.Errorf("maturity level = %q, want %q", got, MaturityStart)
	}
}

func TestLoadRejectsInvalidRuleMetadata(t *testing.T) {
	for _, tc := range []struct {
		name  string
		field string
		value string
	}{
		{name: "invalid domain", field: "principle_domain", value: "Security"},
		{name: "invalid maturity", field: "maturity_level", value: "expert"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			data := []byte("rules:\n  - id: TEST-1\n    scope: repo\n    severity: info\n    title: Example\n    " + tc.field + ": " + tc.value + "\n    pass: true\n")
			if _, err := Load(data); err == nil {
				t.Fatal("invalid metadata should fail ruleset loading")
			}
		})
	}
}

func TestLoadAllowsMissingRuleMetadata(t *testing.T) {
	if _, err := Load([]byte(`rules:
  - id: TEST-1
    scope: repo
    severity: info
    title: Existing custom rule
    pass: true
`)); err != nil {
		t.Fatalf("rules without optional metadata should remain valid: %v", err)
	}
}
