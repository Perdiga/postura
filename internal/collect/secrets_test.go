package collect

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestOldestActionsSecretAge(t *testing.T) {
	updated := time.Now().Add(-45 * 24 * time.Hour).UTC().Format(time.RFC3339)
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/repos/acme/web/actions/secrets" {
			t.Errorf("request path = %q, want repo Actions secrets", r.URL.Path)
		}
		_, _ = fmt.Fprintf(w, `{"total_count":1,"secrets":[{"name":"DEPLOY_TOKEN","updated_at":%q}]}`, updated)
	})
	age, ok := oldestActionsSecretAge(context.Background(), c, "repos/acme/web/actions/secrets")
	if !ok || age < 44 || age > 45 {
		t.Fatalf("oldest secret age = (%d, %v), want roughly 45 days", age, ok)
	}
}

func TestOldestActionsSecretAgeEmpty(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, `{"total_count":0,"secrets":[]}`)
	})
	age, ok := oldestActionsSecretAge(context.Background(), c, "orgs/acme/actions/secrets")
	if !ok || age != 0 {
		t.Fatalf("empty secret age = (%d, %v), want (0, true)", age, ok)
	}
}

func TestOldestActionsSecretAgeUnknown(t *testing.T) {
	cases := map[string]http.HandlerFunc{
		"missing update date": func(w http.ResponseWriter, _ *http.Request) {
			_, _ = io.WriteString(w, `{"secrets":[{"name":"UNDATED"}]}`)
		},
		"forbidden": func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusForbidden)
		},
	}
	for name, handler := range cases {
		t.Run(name, func(t *testing.T) {
			c := testClient(t, handler)
			if _, ok := oldestActionsSecretAge(context.Background(), c, "repos/acme/web/actions/secrets"); ok {
				t.Fatal("incomplete or unreadable secret list must be unknown")
			}
		})
	}
}
