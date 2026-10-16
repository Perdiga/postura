package rules

import (
	"fmt"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// Config holds per-target rule overrides, layered most-specific-wins on top of
// each rule's built-in defaults:
//
//	rule.Config  ->  enterprises[slug]  ->  orgs[name]  ->  repos[owner/repo]
//
// Each override block is keyed by rule id. Within a block the reserved keys
// `enabled` (bool) and `severity` (string) tune the rule itself; every other key
// is merged into the rule's `cfg` map available to its CEL expressions.
type Config struct {
	Enterprises map[string]map[string]map[string]any `yaml:"enterprises"`
	Orgs        map[string]map[string]map[string]any `yaml:"orgs"`
	Repos       map[string]map[string]map[string]any `yaml:"repos"`
	Assessments AssessmentConfig                     `yaml:"assessments"`
	Maturity    MaturityPolicy                       `yaml:"maturity"`
}

type MaturityPolicy struct {
	MaxWarningsPerTier int  `yaml:"max_warnings_per_tier"`
	IgnoreWarnings     bool `yaml:"ignore_warnings"`
}

func (p *MaturityPolicy) UnmarshalYAML(node *yaml.Node) error {
	if node.Kind != yaml.MappingNode {
		return fmt.Errorf("maturity must be a mapping")
	}
	for i := 0; i+1 < len(node.Content); i += 2 {
		key, value := node.Content[i].Value, node.Content[i+1]
		switch key {
		case "max_warnings_per_tier":
			if value.Tag != "!!int" {
				return fmt.Errorf("maturity.max_warnings_per_tier must be an integer")
			}
		case "ignore_warnings":
			if value.Tag != "!!bool" {
				return fmt.Errorf("maturity.ignore_warnings must be a boolean")
			}
		default:
			return fmt.Errorf("unknown maturity setting %q", key)
		}
	}
	type plain MaturityPolicy
	var decoded plain
	if err := node.Decode(&decoded); err != nil {
		return err
	}
	*p = MaturityPolicy(decoded)
	return p.Validate()
}

func (p MaturityPolicy) Validate() error {
	if p.MaxWarningsPerTier < 0 {
		return fmt.Errorf("maturity.max_warnings_per_tier must be non-negative")
	}
	if p.IgnoreWarnings && p.MaxWarningsPerTier != 0 {
		return fmt.Errorf("maturity.ignore_warnings and a non-zero max_warnings_per_tier cannot be combined")
	}
	return nil
}

type AssessmentConfig struct {
	AsOf        string                                `yaml:"as_of"`
	MaxAgeDays  int                                   `yaml:"max_age_days"`
	Enterprises map[string]map[string]AssessmentInput `yaml:"enterprises"`
	Orgs        map[string]map[string]AssessmentInput `yaml:"orgs"`
	Repos       map[string]map[string]AssessmentInput `yaml:"repos"`
}

type AssessmentInput struct {
	// Title and Question are copied from the rule by `config template` as reviewer guidance; they don't affect scoring.
	Title      string `yaml:"title,omitempty"`
	Question   string `yaml:"question,omitempty"`
	Status     string `yaml:"status"`
	Evidence   string `yaml:"evidence"`
	ReviewedAt string `yaml:"reviewed_at"`
	ReviewedBy string `yaml:"reviewed_by"`
}

// LoadConfig parses a policy-config YAML file. An empty/nil input yields a
// zero Config (defaults only), which is valid.
func LoadConfig(data []byte) (*Config, error) {
	c := &Config{}
	if len(data) == 0 {
		return c, nil
	}
	if err := yaml.Unmarshal(data, c); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	if err := c.Assessments.validate(); err != nil {
		return nil, err
	}
	if err := c.Maturity.Validate(); err != nil {
		return nil, err
	}
	return c, nil
}

func (a AssessmentConfig) validate() error {
	hasInputs := false
	for _, targets := range []map[string]map[string]AssessmentInput{a.Enterprises, a.Orgs, a.Repos} {
		for _, entries := range targets {
			if len(entries) > 0 {
				hasInputs = true
			}
		}
	}
	var asOf time.Time
	if a.AsOf != "" {
		parsed, err := time.Parse("2006-01-02", a.AsOf)
		if err != nil {
			return fmt.Errorf("assessments.as_of must be YYYY-MM-DD")
		}
		asOf = parsed
	}
	if hasInputs && (asOf.IsZero() || a.MaxAgeDays <= 0) {
		return fmt.Errorf("assessments.as_of and positive assessments.max_age_days are required when assessments are supplied")
	}
	for _, group := range []struct {
		scope   string
		targets map[string]map[string]AssessmentInput
	}{
		{"enterprise", a.Enterprises}, {"org", a.Orgs}, {"repo", a.Repos},
	} {
		for target, entries := range group.targets {
			for id, input := range entries {
				if strings.TrimSpace(target) == "" || strings.TrimSpace(id) == "" {
					return fmt.Errorf("%s assessment: target and rule id are required", group.scope)
				}
				switch input.Status {
				case "compliant", "partial", "non_compliant", "not_applicable":
				default:
					return fmt.Errorf("%s assessment %s/%s: invalid status %q", group.scope, target, id, input.Status)
				}
				if strings.TrimSpace(input.Evidence) == "" || strings.TrimSpace(input.ReviewedBy) == "" {
					return fmt.Errorf("%s assessment %s/%s: evidence and reviewed_by are required", group.scope, target, id)
				}
				reviewedAt, err := time.Parse("2006-01-02", input.ReviewedAt)
				if err != nil {
					return fmt.Errorf("%s assessment %s/%s: reviewed_at must be YYYY-MM-DD", group.scope, target, id)
				}
				if !asOf.IsZero() && reviewedAt.After(asOf) {
					return fmt.Errorf("%s assessment %s/%s: reviewed_at is after assessments.as_of", group.scope, target, id)
				}
			}
		}
	}
	return nil
}

// Resolved is a rule's effective settings for one specific target.
type Resolved struct {
	Enabled  bool
	Severity string
	Cfg      map[string]any
}

// Resolve computes the effective settings for rule r against a target located at
// (enterprise, org, repo). Empty location components are skipped. The most
// specific override present wins for each individual key.
func (c *Config) Resolve(r Rule, enterprise, org, repo string) Resolved {
	res := Resolved{Enabled: r.Enabled == nil || *r.Enabled, Severity: r.Severity, Cfg: map[string]any{}}
	for k, v := range r.Config {
		res.Cfg[k] = v
	}

	apply := func(block map[string]any) {
		if block == nil {
			return
		}
		for k, v := range block {
			switch k {
			case "enabled":
				if b, ok := v.(bool); ok {
					res.Enabled = b
				}
			case "severity":
				if s, ok := v.(string); ok {
					res.Severity = s
				}
			default:
				res.Cfg[k] = v
			}
		}
	}

	if enterprise != "" {
		apply(c.Enterprises[enterprise][r.ID])
	}
	if org != "" {
		apply(c.Orgs[org][r.ID])
	}
	if repo != "" {
		apply(c.Repos[repo][r.ID])
	}
	var assessment *AssessmentInput
	if enterprise != "" {
		if a, ok := c.Assessments.Enterprises[enterprise][r.ID]; ok {
			assessment = &a
		}
	}
	if org != "" {
		if a, ok := c.Assessments.Orgs[org][r.ID]; ok {
			assessment = &a
		}
	}
	if repo != "" {
		if a, ok := c.Assessments.Repos[repo][r.ID]; ok {
			assessment = &a
		}
	}
	if assessment != nil {
		stale := false
		if asOf, err := time.Parse("2006-01-02", c.Assessments.AsOf); err == nil && c.Assessments.MaxAgeDays > 0 {
			if reviewedAt, err := time.Parse("2006-01-02", assessment.ReviewedAt); err == nil {
				stale = int(asOf.Sub(reviewedAt).Hours()/24) > c.Assessments.MaxAgeDays
			}
		}
		res.Cfg["assessment"] = map[string]any{
			"status": assessment.Status, "evidence": assessment.Evidence,
			"reviewed_at": assessment.ReviewedAt, "reviewed_by": assessment.ReviewedBy, "stale": stale,
		}
	}
	return res
}
