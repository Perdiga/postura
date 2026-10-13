package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestPermissionProbe(t *testing.T) {
	facts := map[string]any{"two_factor_requirement_enabled": true, "admin_count": 2}
	check := probe("org", "acme", "org settings", facts,
		[]string{"two_factor_requirement_enabled", "admin_count", "organization_roles"},
		"grant admin:org")
	if len(check.missing) != 1 || check.missing[0] != "organization_roles" {
		t.Fatalf("missing facts = %v, want organization_roles", check.missing)
	}
}

func TestPrivateReportingProbeOnlyForPublicRepos(t *testing.T) {
	privateChecks := repoPermissionChecks(map[string]any{
		"name": "acme/private", "visibility": "private", "archived": false,
	})
	for _, check := range privateChecks {
		if strings.Contains(check.capability, "private vulnerability reporting") {
			t.Fatal("private repositories should not be probed for private vulnerability reporting")
		}
	}

	publicChecks := repoPermissionChecks(map[string]any{
		"name": "acme/public", "visibility": "public", "archived": false,
	})
	found := false
	for _, check := range publicChecks {
		if strings.Contains(check.capability, "private vulnerability reporting") {
			found = true
			if len(check.missing) != 1 || check.missing[0] != "private_vulnerability_reporting" {
				t.Fatalf("public repo check missing facts = %v", check.missing)
			}
		}
	}
	if !found {
		t.Fatal("public repositories should be probed for private vulnerability reporting")
	}
}

func TestWritePermissionChecks(t *testing.T) {
	checks := []permissionCheck{
		{scope: "org", target: "acme", capability: "Actions settings"},
		{scope: "repo", target: "acme/api", capability: "private reporting", missing: []string{"private_vulnerability_reporting"}, guidance: "Metadata read"},
	}
	var out bytes.Buffer
	if err := writePermissionChecks(&out, checks); err == nil {
		t.Fatal("incomplete permission check should return an error")
	}
	text := out.String()
	for _, expected := range []string{"[OK] org acme", "[INCOMPLETE] repo acme/api", "Metadata read", "Policy failures are not permission failures"} {
		if !strings.Contains(text, expected) {
			t.Errorf("permission output missing %q:\n%s", expected, text)
		}
	}
}
