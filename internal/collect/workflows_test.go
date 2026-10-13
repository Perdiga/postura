package collect

import (
	"context"
	"encoding/base64"
	"io"
	"net/http"
	"testing"
)

func TestUnpinnedActionRefs(t *testing.T) {
	data := []byte(`name: test
on: [push]
jobs:
  build:
    uses: acme/workflows/.github/workflows/build.yml@main
  test:
    steps:
      - uses: actions/checkout@v4
      - uses: acme/setup@0123456789abcdef0123456789abcdef01234567
      - uses: ./.github/actions/local
      - uses: docker://alpine:3.20
`)
	got, err := unpinnedActionRefs(data)
	if err != nil {
		t.Fatal(err)
	}
	if got != 2 {
		t.Fatalf("unpinnedActionRefs = %d, want 2", got)
	}
}

func TestUnpinnedActionRefsMalformedYAML(t *testing.T) {
	if _, err := unpinnedActionRefs([]byte("jobs: [")); err == nil {
		t.Fatal("malformed workflow YAML should return an error")
	}
}

func TestUnpinnedActionRefsWithoutJobs(t *testing.T) {
	got, err := unpinnedActionRefs([]byte("name: empty workflow\n"))
	if err != nil {
		t.Fatal(err)
	}
	if got != 0 {
		t.Fatalf("unpinnedActionRefs = %d, want 0", got)
	}
}

func TestWorkflowFilesAndUnpinnedActionRefs(t *testing.T) {
	workflow1 := "permissions:\n  contents: read\njobs:\n  test:\n    steps:\n      - uses: actions/checkout@v4\n"
	workflow2 := "permissions:\n  contents: read\njobs:\n  test:\n    steps:\n      - uses: acme/test@0123456789abcdef0123456789abcdef01234567\n"
	body := treeBody(t, []treeEnt{
		workflowBlob(".github/workflows/ci.yml", "blob-one"),
		workflowBlob(".github/workflows/deploy.yaml", "blob-two"),
		workflowBlob(".github/workflows/nested/ignored.yml", "blob-three"),
		workflowBlob(".github/workflows/README.md", "blob-four"),
	}, false)
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/repos/acme/web/git/trees/main":
			if r.URL.Query().Get("recursive") != "1" {
				t.Errorf("tree request recursive = %q, want 1", r.URL.Query().Get("recursive"))
			}
			_, _ = io.WriteString(w, body)
		case "/repos/acme/web/git/blobs/blob-one":
			_, _ = io.WriteString(w, blobResponse(t, workflow1))
		case "/repos/acme/web/git/blobs/blob-two":
			_, _ = io.WriteString(w, blobResponse(t, workflow2))
		default:
			t.Errorf("unexpected request path %q", r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
		}
	})

	facts := repoFiles(context.Background(), c, "acme", "web", "main")
	if len(facts.workflowFiles) != 2 || !facts.workflowFilesKnown {
		t.Fatalf("workflow file facts = %+v, want two known top-level YAML workflows", facts)
	}
	counts, known := workflowSecurityCountsForRepo(context.Background(), c, "acme", "web", facts)
	if !known || counts.unpinnedActionRefCount != 1 || counts.permissionIssueCount != 0 {
		t.Fatalf("workflow security counts = (%+v, %v), want one unpinned ref and zero permission issues", counts, known)
	}
}

func TestUnpinnedWorkflowActionRefsUnreadableBlob(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})
	counts, known := workflowSecurityCountsForRepo(context.Background(), c, "acme", "web", fileFacts{
		workflowFiles: map[string]string{".github/workflows/ci.yml": "blob"}, workflowFilesKnown: true,
	})
	if known || counts != (workflowSecurityCounts{}) {
		t.Fatalf("unreadable workflow should be unknown, got (%+v, %v)", counts, known)
	}
}

func TestWorkflowPermissionIssues(t *testing.T) {
	cases := []struct {
		name string
		data string
		want int
	}{
		{
			name: "workflow-level read-only permissions are inherited",
			data: `permissions:
  contents: read
jobs:
  test:
    steps: []
`,
			want: 0,
		},
		{
			name: "OIDC id-token write is a narrow exception",
			data: `permissions:
  contents: read
  id-token: write
jobs:
  deploy:
    steps: []
`,
			want: 0,
		},
		{
			name: "job-level write grant is reported",
			data: `jobs:
  test:
    permissions:
      contents: read
      issues: write
`,
			want: 1,
		},
		{
			name: "implicit permissions are reported",
			data: `jobs:
  test:
    steps: []
`,
			want: 1,
		},
		{
			name: "write-all is reported",
			data: `permissions: write-all
jobs:
  test:
    steps: []
`,
			want: 1,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := workflowPermissionIssues([]byte(tc.data))
			if err != nil {
				t.Fatal(err)
			}
			if got != tc.want {
				t.Errorf("workflowPermissionIssues = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestInspectWorkflow(t *testing.T) {
	data := []byte(`name: secure-ci
on:
  pull_request:
    paths: ["src/**"]
  push:
    branches: [main]
concurrency:
  group: ci-${{ github.ref }}
jobs:
  build:
    strategy:
      matrix:
        go: ["1.25", "1.26"]
    steps:
      - uses: github/codeql-action/init@0123456789abcdef0123456789abcdef01234567
      - uses: actions/dependency-review-action@0123456789abcdef0123456789abcdef01234567
      - uses: actions/cache@0123456789abcdef0123456789abcdef01234567
      - uses: actions/stale@0123456789abcdef0123456789abcdef01234567
      - uses: googleapis/release-please-action@0123456789abcdef0123456789abcdef01234567
  lint:
    uses: acme/workflows/.github/workflows/lint.yml@0123456789abcdef0123456789abcdef01234567
`)
	got, err := inspectWorkflow(data)
	if err != nil {
		t.Fatal(err)
	}
	want := workflowStaticFacts{
		pullRequestTrigger: true, narrowTrigger: true,
		codeQLActionCount: 1, dependencyReviewCount: 1, cacheActionCount: 1,
		concurrency: true, matrixJobs: 1, reusableWorkflowCount: 1,
		releaseAutomationCount: 1, staleIssueAutomationCount: 1,
	}
	if got != want {
		t.Errorf("inspectWorkflow = %+v, want %+v", got, want)
	}
}

func blobResponse(t *testing.T, content string) string {
	t.Helper()
	return `{"encoding":"base64","content":"` + base64.StdEncoding.EncodeToString([]byte(content)) + `"}`
}
