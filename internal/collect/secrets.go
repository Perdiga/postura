package collect

import (
	"context"
	"time"

	"github.com/google/go-github/v84/github"
)

func oldestActionsSecretAge(ctx context.Context, c *github.Client, path string) (int, bool) {
	now := time.Now()
	oldest := 0
	page := 1
	for {
		var body struct {
			Secrets []struct {
				UpdatedAt time.Time `json:"updated_at"`
			} `json:"secrets"`
		}
		resp, err := getJSON(ctx, c, path+"?per_page=100&page="+itoa(int64(page)), &body)
		if err != nil {
			return 0, false
		}
		for _, secret := range body.Secrets {
			if secret.UpdatedAt.IsZero() {
				return 0, false
			}
			if age := ageDays(now, secret.UpdatedAt); age > oldest {
				oldest = age
			}
		}
		if resp == nil || resp.NextPage == 0 {
			return oldest, true
		}
		if resp.NextPage == page {
			return 0, false
		}
		page = resp.NextPage
	}
}
