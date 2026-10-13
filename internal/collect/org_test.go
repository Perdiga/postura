package collect

import (
	"context"
	"io"
	"net/http"
	"testing"
)

func TestOrganizationActionsFacts(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/orgs/acme/actions/permissions":
			_, _ = io.WriteString(w, `{"enabled_repositories":"selected","allowed_actions":"selected"}`)
		case "/orgs/acme/actions/permissions/workflow":
			_, _ = io.WriteString(w, `{"default_workflow_permissions":"read","can_approve_pull_request_reviews":false}`)
		default:
			t.Errorf("unexpected request path %q", r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
		}
	})
	facts := organizationActionsFacts(context.Background(), c, "acme")
	if facts["actions_enabled_repositories"] != "selected" || facts["actions_allowed_actions"] != "selected" ||
		facts["actions_default_workflow_permissions"] != "read" || facts["actions_can_approve_pull_request_reviews"] != false {
		t.Fatalf("organization Actions facts = %+v", facts)
	}
}

func TestOrganizationActionsFactsUnreadable(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})
	if facts := organizationActionsFacts(context.Background(), c, "acme"); facts != nil {
		t.Fatalf("unreadable organization Actions settings should be omitted, got %+v", facts)
	}
}

func TestOrganizationRunnerGroups(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/orgs/acme/actions/runner-groups" {
			t.Errorf("request path = %q, want runner groups", r.URL.Path)
		}
		_, _ = io.WriteString(w, `{"total_count":1,"runner_groups":[{"id":7,"name":"private","visibility":"selected","allows_public_repositories":false}]}`)
	})
	groups, ok := organizationRunnerGroups(context.Background(), c, "acme")
	if !ok || len(groups) != 1 {
		t.Fatalf("runner groups = (%v, %v), want one known group", groups, ok)
	}
	group := groups[0].(map[string]any)
	if group["visibility"] != "selected" || group["allows_public_repositories"] != false {
		t.Errorf("runner group facts = %+v", group)
	}
}

func TestOrganizationRunnerGroupsUnreadable(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})
	if _, ok := organizationRunnerGroups(context.Background(), c, "acme"); ok {
		t.Fatal("unreadable runner groups must be unknown")
	}
}
