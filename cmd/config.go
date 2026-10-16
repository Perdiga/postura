package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/v84/github"
	"github.com/Perdiga/postura/internal/collect"
	"github.com/Perdiga/postura/internal/rules"
	"github.com/spf13/cobra"
)

var (
	configTemplateRules      string
	configTemplateOutput     string
	configTemplateEnterprise string
	configTemplateOrgs       []string
	configTemplateAsOf       string
	configTemplateMaxAge     int
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Generate policy configuration files",
}

var configTemplateCmd = &cobra.Command{
	Use:   "template [owner/repo ...]",
	Short: "Generate a blank human-assessment config for selected targets",
	Long:  "Generate an assessment-config template containing every assessment-marked rule for the selected enterprise, organizations, and repositories. Passing an org without repo arguments discovers all non-forked repositories in that org.",
	RunE:  runConfigTemplate,
}

func init() {
	f := configTemplateCmd.Flags()
	f.StringVar(&configTemplateRules, "rules", "", "ruleset YAML containing assessment-marked rules (required)")
	f.StringVar(&configTemplateOutput, "output", "config.yaml", "output path; use - for stdout")
	f.StringVar(&configTemplateEnterprise, "enterprise", "", "enterprise slug to include")
	f.StringSliceVar(&configTemplateOrgs, "org", nil, "organization(s) to include (repeatable/comma-separated)")
	f.StringVar(&configTemplateAsOf, "as-of", "", "assessment as-of date (YYYY-MM-DD; defaults to today in UTC)")
	f.IntVar(&configTemplateMaxAge, "max-age-days", 180, "maximum age of an assessment before it becomes unknown")
	_ = configTemplateCmd.MarkFlagRequired("rules")
	configCmd.AddCommand(configTemplateCmd)
	rootCmd.AddCommand(configCmd)
}

func runConfigTemplate(cmd *cobra.Command, args []string) error {
	rs, err := loadRuleset(configTemplateRules)
	if err != nil {
		return err
	}
	specs, err := orgsFromArgs(configTemplateOrgs, args)
	if err != nil {
		return err
	}
	if len(specs) == 0 && configTemplateEnterprise == "" {
		return fmt.Errorf("pass --enterprise, --org, or owner/repo targets")
	}

	repoTargets := make([]string, 0)
	if len(specs) > 0 {
		ctx := context.Background()
		var client *github.Client
		for _, spec := range specs {
			repoNames := spec.repos
			if len(repoNames) == 0 {
				if client == nil {
					token, err := resolveToken()
					if err != nil {
						return err
					}
					client = collect.NewClient(ctx, token)
				}
				repoNames, err = collect.RepositoryNames(ctx, client, spec.name)
				if err != nil {
					return fmt.Errorf("list repositories for org %s: %w", spec.name, err)
				}
			}
			for _, repo := range repoNames {
				repoTargets = append(repoTargets, spec.name+"/"+repo)
			}
		}
	}

	asOf := configTemplateAsOf
	if asOf == "" {
		asOf = time.Now().UTC().Format("2006-01-02")
	}
	data, err := rules.GenerateAssessmentTemplate(rs, rules.AssessmentTargets{
		Enterprises: nonEmpty(configTemplateEnterprise),
		Orgs:        orgNames(specs),
		Repos:       repoTargets,
	}, asOf, configTemplateMaxAge)
	if err != nil {
		return err
	}
	if configTemplateOutput == "-" {
		_, err = cmd.OutOrStdout().Write(data)
		return err
	}
	file, err := os.OpenFile(configTemplateOutput, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		return fmt.Errorf("create %s (will not overwrite): %w", configTemplateOutput, err)
	}
	if _, err := file.Write(data); err != nil {
		_ = file.Close()
		return fmt.Errorf("write %s: %w", configTemplateOutput, err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("close %s: %w", configTemplateOutput, err)
	}
	_, _ = fmt.Fprintf(cmd.ErrOrStderr(), "generated assessment template: %s\n", configTemplateOutput)
	return nil
}

func orgNames(specs []orgSpec) []string {
	out := make([]string, 0, len(specs))
	for _, spec := range specs {
		out = append(out, spec.name)
	}
	return out
}

func nonEmpty(value string) []string {
	if value == "" {
		return nil
	}
	return []string{value}
}
