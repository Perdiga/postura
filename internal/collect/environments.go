package collect

import (
	"context"

	"github.com/google/go-github/v84/github"
)

func repoEnvironments(ctx context.Context, c *github.Client, org, repo string) ([]any, bool) {
	var out []any
	page := 1
	for {
		var body struct {
			Environments []struct {
				Name            string `json:"name"`
				ProtectionRules []struct {
					Type string `json:"type"`
				} `json:"protection_rules"`
			} `json:"environments"`
		}
		resp, err := getJSON(ctx, c, "repos/"+org+"/"+repo+"/environments?per_page=100&page="+itoa(int64(page)), &body)
		if err != nil {
			return nil, false
		}
		for _, environment := range body.Environments {
			requiredReviewers := false
			for _, rule := range environment.ProtectionRules {
				if rule.Type == "required_reviewers" {
					requiredReviewers = true
					break
				}
			}
			out = append(out, map[string]any{
				"name":                  environment.Name,
				"required_reviewers":    requiredReviewers,
				"protection_rule_count": len(environment.ProtectionRules),
			})
		}
		if resp == nil || resp.NextPage == 0 {
			return emptyIfNil(out), true
		}
		if resp.NextPage == page {
			return nil, false
		}
		page = resp.NextPage
	}
}
