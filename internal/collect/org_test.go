package collect

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestOrgInvitations(t *testing.T) {
	for _, tt := range []struct {
		name         string
		status       int
		body         string
		failNextPage bool
		wantCount    int
		wantWarning  bool
	}{
		{name: "empty", status: http.StatusOK, body: `[]`},
		{name: "pending", status: http.StatusOK, body: `[{"login":"alice"}]`, wantCount: 1},
		{name: "not found", status: http.StatusNotFound, wantWarning: true},
		{name: "forbidden", status: http.StatusForbidden, wantWarning: true},
		{name: "server error", status: http.StatusInternalServerError, wantWarning: true},
		{name: "partial list", status: http.StatusOK, body: `[{"login":"alice"}]`, failNextPage: true, wantWarning: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			continued := false
			c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.Path {
				case "/orgs/acme":
					_, _ = io.WriteString(w, `{"login":"acme"}`)
				case "/orgs/acme/members", "/orgs/acme/outside_collaborators", "/orgs/acme/teams":
					_, _ = io.WriteString(w, `[]`)
				case "/orgs/acme/invitations":
					if r.URL.Query().Get("per_page") != "100" {
						t.Errorf("invitations query = %q, want per_page=100", r.URL.RawQuery)
					}
					if tt.failNextPage {
						if r.URL.Query().Get("page") == "2" {
							w.WriteHeader(http.StatusNotFound)
							return
						}
						w.Header().Set("Link", "<http://"+r.Host+"/orgs/acme/invitations?per_page=100&page=2>; rel=\"next\"")
					}
					w.WriteHeader(tt.status)
					_, _ = io.WriteString(w, tt.body)
				case "/orgs/acme/actions/permissions":
					continued = true
					_, _ = io.WriteString(w, `{"enabled_repositories":"all","allowed_actions":"all"}`)
				default:
					w.WriteHeader(http.StatusNotFound)
				}
			})
			var warnings bytes.Buffer
			facts, err := Org(context.Background(), c, "acme", &warnings)
			if err != nil {
				t.Fatalf("Org: %v", err)
			}
			if !continued || facts["actions_enabled_repositories"] != "all" {
				t.Fatalf("collection did not continue after invitations: %+v", facts)
			}
			invitations, present := facts["invitations"]
			if tt.wantWarning {
				if present {
					t.Fatalf("unavailable invitations must be omitted, got %v", invitations)
				}
				wantError := strconv.Itoa(tt.status)
				if tt.failNextPage {
					wantError = strconv.Itoa(http.StatusNotFound)
				}
				if !strings.Contains(warnings.String(), "warning: collect org acme: pending invitations unavailable") ||
					!strings.Contains(warnings.String(), wantError) {
					t.Fatalf("missing warning with error details: %q", warnings.String())
				}
			} else {
				if warnings.Len() != 0 {
					t.Fatalf("unexpected warning: %s", warnings.String())
				}
				list, ok := invitations.([]any)
				if !present || !ok || len(list) != tt.wantCount {
					t.Fatalf("invitations = %v, want %d entries", invitations, tt.wantCount)
				}
			}
			output, err := json.Marshal(facts)
			if err != nil {
				t.Fatalf("marshal facts: %v", err)
			}
			if strings.Contains(string(output), "warning:") {
				t.Fatalf("warning leaked into facts JSON: %s", output)
			}
		})
	}
}

func TestOrgOtherErrorsRemainFatal(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	})
	var warnings bytes.Buffer
	facts, err := Org(context.Background(), c, "acme", &warnings)
	if err == nil || facts != nil {
		t.Fatalf("Org = (%v, %v), want fatal error", facts, err)
	}
	if warnings.Len() != 0 {
		t.Fatalf("unexpected warning: %s", warnings.String())
	}
}

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
