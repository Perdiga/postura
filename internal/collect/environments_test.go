package collect

import (
	"context"
	"io"
	"net/http"
	"testing"
)

func TestRepoEnvironments(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/repos/acme/web/environments" {
			t.Errorf("request path = %q, want repo environments", r.URL.Path)
		}
		_, _ = io.WriteString(w, `{"total_count":2,"environments":[
			{"name":"production","protection_rules":[{"type":"required_reviewers"}]},
			{"name":"preview","protection_rules":[]}
		]}`)
	})
	environments, ok := repoEnvironments(context.Background(), c, "acme", "web")
	if !ok || len(environments) != 2 {
		t.Fatalf("environments = (%v, %v), want two known environments", environments, ok)
	}
	production := environments[0].(map[string]any)
	preview := environments[1].(map[string]any)
	if production["required_reviewers"] != true || preview["required_reviewers"] != false {
		t.Fatalf("environment protection facts = %v, %v", production, preview)
	}
}

func TestRepoEnvironmentsUnreadable(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})
	if _, ok := repoEnvironments(context.Background(), c, "acme", "web"); ok {
		t.Fatal("unreadable environments must be unknown")
	}
}
