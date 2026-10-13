package collect

import (
	"fmt"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

var commitSHA = regexp.MustCompile(`(?i)^[0-9a-f]{40}$`)

type workflowStaticFacts struct {
	pullRequestTrigger        bool
	narrowTrigger             bool
	codeQLActionCount         int
	dependencyReviewCount     int
	cacheActionCount          int
	concurrency               bool
	matrixJobs                int
	reusableWorkflowCount     int
	releaseAutomationCount    int
	staleIssueAutomationCount int
}

// unpinnedActionRefs counts remote Actions and reusable-workflow references that
// do not use a full commit SHA. Local actions and Docker image references are not
// GitHub Action repository refs and are excluded.
func unpinnedActionRefs(data []byte) (int, error) {
	var root yaml.Node
	if err := yaml.Unmarshal(data, &root); err != nil {
		return 0, fmt.Errorf("parse workflow YAML: %w", err)
	}
	if len(root.Content) == 0 || root.Content[0].Kind != yaml.MappingNode {
		return 0, fmt.Errorf("workflow must be a YAML mapping")
	}

	jobs := mappingValue(root.Content[0], "jobs")
	if jobs == nil || jobs.Kind != yaml.MappingNode {
		return 0, nil
	}

	count := 0
	for i := 1; i < len(jobs.Content); i += 2 {
		job := jobs.Content[i]
		if job.Kind != yaml.MappingNode {
			continue
		}
		if uses := mappingValue(job, "uses"); uses != nil && isUnpinned(uses.Value) {
			count++
		}
		steps := mappingValue(job, "steps")
		if steps == nil || steps.Kind != yaml.SequenceNode {
			continue
		}
		for _, step := range steps.Content {
			if step.Kind != yaml.MappingNode {
				continue
			}
			if uses := mappingValue(step, "uses"); uses != nil && isUnpinned(uses.Value) {
				count++
			}
		}
	}
	return count, nil
}

// workflowPermissionIssues counts jobs whose effective token permissions are
// implicit, malformed, or include any write grant. Jobs inherit workflow-level
// permissions when they do not declare their own.
func workflowPermissionIssues(data []byte) (int, error) {
	var root yaml.Node
	if err := yaml.Unmarshal(data, &root); err != nil {
		return 0, fmt.Errorf("parse workflow YAML: %w", err)
	}
	if len(root.Content) == 0 || root.Content[0].Kind != yaml.MappingNode {
		return 0, fmt.Errorf("workflow must be a YAML mapping")
	}
	rootNode := root.Content[0]
	jobs := mappingValue(rootNode, "jobs")
	if jobs == nil || jobs.Kind != yaml.MappingNode {
		return 0, nil
	}
	workflowPermissions := mappingValue(rootNode, "permissions")
	issues := 0
	for i := 1; i < len(jobs.Content); i += 2 {
		job := jobs.Content[i]
		if job.Kind != yaml.MappingNode {
			issues++
			continue
		}
		permissions := mappingValue(job, "permissions")
		if permissions == nil {
			permissions = workflowPermissions
		}
		if permissions == nil || hasWritePermission(permissions) {
			issues++
		}
	}
	return issues, nil
}

func hasWritePermission(permissions *yaml.Node) bool {
	switch permissions.Kind {
	case yaml.ScalarNode:
		return permissions.Value != "read-all"
	case yaml.MappingNode:
		for i := 0; i+1 < len(permissions.Content); i += 2 {
			name := permissions.Content[i]
			level := permissions.Content[i+1]
			if level.Kind != yaml.ScalarNode || (level.Value != "read" && level.Value != "none") {
				if name.Kind == yaml.ScalarNode && name.Value == "id-token" && level.Kind == yaml.ScalarNode && level.Value == "write" {
					continue
				}
				return true
			}
		}
		return false
	default:
		return true
	}
}

func inspectWorkflow(data []byte) (workflowStaticFacts, error) {
	var root yaml.Node
	if err := yaml.Unmarshal(data, &root); err != nil {
		return workflowStaticFacts{}, fmt.Errorf("parse workflow YAML: %w", err)
	}
	if len(root.Content) == 0 || root.Content[0].Kind != yaml.MappingNode {
		return workflowStaticFacts{}, fmt.Errorf("workflow must be a YAML mapping")
	}
	rootNode := root.Content[0]
	facts := workflowStaticFacts{}
	facts.pullRequestTrigger, facts.narrowTrigger = inspectTriggers(mappingValue(rootNode, "on"))
	facts.concurrency = mappingValue(rootNode, "concurrency") != nil
	jobs := mappingValue(rootNode, "jobs")
	if jobs == nil || jobs.Kind != yaml.MappingNode {
		return facts, nil
	}
	for i := 1; i < len(jobs.Content); i += 2 {
		job := jobs.Content[i]
		if job.Kind != yaml.MappingNode {
			continue
		}
		if mappingValue(job, "concurrency") != nil {
			facts.concurrency = true
		}
		strategy := mappingValue(job, "strategy")
		if strategy != nil && mappingValue(strategy, "matrix") != nil {
			facts.matrixJobs++
		}
		if uses := mappingValue(job, "uses"); uses != nil {
			facts.addActionReference(uses.Value, true)
		}
		steps := mappingValue(job, "steps")
		if steps == nil || steps.Kind != yaml.SequenceNode {
			continue
		}
		for _, step := range steps.Content {
			if step.Kind != yaml.MappingNode {
				continue
			}
			if uses := mappingValue(step, "uses"); uses != nil {
				facts.addActionReference(uses.Value, false)
			}
		}
	}
	return facts, nil
}

func (f *workflowStaticFacts) addActionReference(ref string, reusable bool) {
	lower := strings.ToLower(ref)
	switch {
	case strings.Contains(lower, "github/codeql-action/"):
		f.codeQLActionCount++
	case strings.Contains(lower, "actions/dependency-review-action"):
		f.dependencyReviewCount++
	case strings.Contains(lower, "actions/cache"):
		f.cacheActionCount++
	case strings.Contains(lower, "actions/stale") || strings.Contains(lower, "probot/stale"):
		f.staleIssueAutomationCount++
	case strings.Contains(lower, "release-please") || strings.Contains(lower, "semantic-release") ||
		strings.Contains(lower, "changesets/action") || strings.Contains(lower, "softprops/action-gh-release"):
		f.releaseAutomationCount++
	}
	if reusable && strings.Contains(lower, ".github/workflows/") {
		f.reusableWorkflowCount++
	}
}

func inspectTriggers(node *yaml.Node) (pullRequest, narrow bool) {
	if node == nil {
		return false, false
	}
	check := func(name string, details *yaml.Node) {
		if name == "pull_request" || name == "pull_request_target" {
			pullRequest = true
		}
		if details != nil && details.Kind == yaml.MappingNode &&
			(mappingValue(details, "paths") != nil || mappingValue(details, "paths-ignore") != nil ||
				mappingValue(details, "branches") != nil || mappingValue(details, "branches-ignore") != nil) {
			narrow = true
		}
	}
	switch node.Kind {
	case yaml.ScalarNode:
		check(node.Value, nil)
	case yaml.SequenceNode:
		for _, event := range node.Content {
			check(event.Value, nil)
		}
	case yaml.MappingNode:
		for i := 0; i+1 < len(node.Content); i += 2 {
			check(node.Content[i].Value, node.Content[i+1])
		}
	}
	return pullRequest, narrow
}

func mappingValue(node *yaml.Node, key string) *yaml.Node {
	for i := 0; i+1 < len(node.Content); i += 2 {
		if node.Content[i].Value == key {
			return node.Content[i+1]
		}
	}
	return nil
}

func isUnpinned(ref string) bool {
	if strings.HasPrefix(ref, "./") || strings.HasPrefix(ref, "docker://") {
		return false
	}
	i := strings.LastIndexByte(ref, '@')
	return i < 0 || !commitSHA.MatchString(ref[i+1:])
}
