package report

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"sort"
	"time"

	"github.com/Perdiga/postura/internal/engine"
	"github.com/Perdiga/postura/internal/rules"
)

//go:embed html.tmpl
var htmlTemplate string

const wellArchitectedURL = "https://learn.github.com/well-architected"

type pillarInfo struct {
	Name, URL, Summary string
}

// pillars are the GitHub Well-Architected pillars, in framework order.
var pillars = []pillarInfo{
	{rules.DomainProductivity, wellArchitectedURL + "/productivity/overview", "Streamline workflows, automate processes, and accelerate team velocity."},
	{rules.DomainCollaboration, wellArchitectedURL + "/collaboration/overview", "Inclusive practices and transparent project management that foster teamwork."},
	{rules.DomainApplicationSecurity, wellArchitectedURL + "/application-security/overview", "Embed security into every stage of development, from code scanning to compliance."},
	{rules.DomainGovernance, wellArchitectedURL + "/governance/overview", "Policies and controls that balance innovation with oversight and accountability."},
	{rules.DomainArchitecture, wellArchitectedURL + "/architecture/overview", "Scalable, resilient, and efficient GitHub environments."},
}

const unclassified = "Unclassified"

// tiers are the maturity tiers in ascending order.
var tiers = []string{rules.MaturityStart, rules.MaturityMature, rules.MaturityAdvance}

var levelNames = []string{"Start", "Mature", "Advance"}

var scopeOrder = []struct{ key, label string }{
	{string(rules.ScopeEnterprise), "Enterprise"},
	{string(rules.ScopeOrg), "Organizations"},
	{string(rules.ScopeRepo), "Repositories"},
}

type tally struct{ Pass, Fail, Unknown int }

func (t *tally) add(s engine.Status) {
	switch s {
	case engine.StatusPass:
		t.Pass++
	case engine.StatusFail:
		t.Fail++
	case engine.StatusUnknown:
		t.Unknown++
	}
}

func (t tally) Total() int { return t.Pass + t.Fail + t.Unknown }

// Pct is the pass rate, 0–100. Unknown never counts as a pass.
func (t tally) Pct() int {
	if t.Total() == 0 {
		return 0
	}
	return t.Pass * 100 / t.Total()
}

// KnownPct is the share of checks postura could actually score, 0–100.
func (t tally) KnownPct() int {
	if t.Total() == 0 {
		return 0
	}
	return (t.Pass + t.Fail) * 100 / t.Total()
}

type htmlPillar struct {
	pillarInfo
	tally
	Scopes []tally // aligned with htmlReport.Scopes
}

type htmlTarget struct {
	Name  string
	Level int
	tally
	Pillars  []tally // aligned with htmlReport.PillarNames
	Findings []engine.Finding
}

func (t htmlTarget) LevelName() string { return levelNames[t.Level] }

type htmlScope struct {
	Key, Label   string
	Targets      []htmlTarget
	Distribution []int // targets per level, indexed like levelNames
}

type htmlGap struct {
	ID, Title, Severity, Pillar, PillarURL, Scope, FixHint string
	Failing, Unknown                                       []string
}

type htmlTier struct {
	Heading string
	Gaps    []htmlGap
}

type htmlReport struct {
	Generated, ScopeDesc string
	FrameworkURL         string
	Overall              tally
	Pillars              []htmlPillar
	PillarNames          []string
	Scopes               []htmlScope
	Roadmap              []htmlTier
	LevelNames           []string
	Maturity             rules.MaturityPolicy
}

// HTML renders a self-contained maturity report aligned to the GitHub
// Well-Architected pillars, aimed at enterprise/org owners: an executive
// summary, pillar scores per scope, per-target maturity, and a tiered roadmap.
func HTML(fs []engine.Finding, scopeDesc string, generated time.Time, maturity rules.MaturityPolicy) ([]byte, error) {
	if err := maturity.Validate(); err != nil {
		return nil, err
	}
	tmpl, err := template.New("report").Funcs(template.FuncMap{
		"levelClass": func(l int) string { return fmt.Sprintf("lvl%d", l) },
		"add":        func(a, b int) int { return a + b },
		"pct": func(n, total int) int {
			if total == 0 {
				return 0
			}
			return n * 100 / total
		},
		"statusIcon": func(s engine.Status) string {
			switch s {
			case engine.StatusFail:
				return "✗"
			case engine.StatusUnknown:
				return "?"
			}
			return "✓"
		},
	}).Parse(htmlTemplate)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, buildHTMLReport(fs, scopeDesc, generated, maturity)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func pillarOf(f engine.Finding) string {
	if f.PrincipleDomain == "" {
		return unclassified
	}
	return f.PrincipleDomain
}

func buildHTMLReport(fs []engine.Finding, scopeDesc string, generated time.Time, maturity rules.MaturityPolicy) htmlReport {
	r := htmlReport{
		Generated:    generated.UTC().Format("2006-01-02 15:04 MST"),
		ScopeDesc:    scopeDesc,
		FrameworkURL: wellArchitectedURL,
		LevelNames:   levelNames,
		Maturity:     maturity,
	}

	pillarList := append([]pillarInfo(nil), pillars...)
	for _, f := range fs {
		if f.PrincipleDomain == "" {
			pillarList = append(pillarList, pillarInfo{Name: unclassified, Summary: "Checks without a principle_domain."})
			break
		}
	}
	pillarIdx := map[string]int{}
	for i, p := range pillarList {
		pillarIdx[p.Name] = i
		r.PillarNames = append(r.PillarNames, p.Name)
	}

	var scopesPresent []int
	for i, s := range scopeOrder {
		for _, f := range fs {
			if f.Scope == s.key {
				scopesPresent = append(scopesPresent, i)
				break
			}
		}
	}
	scopeIdx := map[string]int{}
	for i, si := range scopesPresent {
		scopeIdx[scopeOrder[si].key] = i
	}

	r.Pillars = make([]htmlPillar, len(pillarList))
	for i, p := range pillarList {
		r.Pillars[i] = htmlPillar{pillarInfo: p, Scopes: make([]tally, len(scopesPresent))}
	}
	for _, f := range fs {
		r.Overall.add(f.Status)
		p := &r.Pillars[pillarIdx[pillarOf(f)]]
		p.add(f.Status)
		if si, ok := scopeIdx[f.Scope]; ok {
			p.Scopes[si].add(f.Status)
		}
	}

	for _, si := range scopesPresent {
		key := scopeOrder[si].key
		sc := htmlScope{Key: key, Label: scopeOrder[si].label, Distribution: make([]int, len(levelNames))}
		byTarget := map[string]int{}
		for _, f := range fs {
			if f.Scope != key {
				continue
			}
			i, ok := byTarget[f.Target]
			if !ok {
				i = len(sc.Targets)
				byTarget[f.Target] = i
				sc.Targets = append(sc.Targets, htmlTarget{Name: f.Target, Pillars: make([]tally, len(pillarList))})
			}
			t := &sc.Targets[i]
			t.add(f.Status)
			t.Pillars[pillarIdx[pillarOf(f)]].add(f.Status)
			t.Findings = append(t.Findings, f)
		}
		for i := range sc.Targets {
			t := &sc.Targets[i]
			t.Level = maturityLevel(t.Findings, maturity)
			sort.SliceStable(t.Findings, func(a, b int) bool { return findingLess(t.Findings[a], t.Findings[b]) })
			sc.Distribution[t.Level]++
		}
		sort.SliceStable(sc.Targets, func(a, b int) bool {
			if sc.Targets[a].Level != sc.Targets[b].Level {
				return sc.Targets[a].Level < sc.Targets[b].Level
			}
			if sc.Targets[a].Pct() != sc.Targets[b].Pct() {
				return sc.Targets[a].Pct() < sc.Targets[b].Pct()
			}
			return sc.Targets[a].Name < sc.Targets[b].Name
		})
		r.Scopes = append(r.Scopes, sc)
	}

	r.Roadmap = buildRoadmap(fs, pillarList)
	return r
}

// maturityLevel returns the lowest blocked tier. Warning allowances are counted
// separately per target and tier; failed and unknown errors always block.
func maturityLevel(fs []engine.Finding, policy rules.MaturityPolicy) int {
	for i, tier := range tiers[:len(tiers)-1] {
		warnings := 0
		for _, f := range fs {
			if f.MaturityLevel == tier && f.Severity != "info" && f.Status != engine.StatusPass {
				if f.Severity != "warn" {
					return i
				}
				warnings++
			}
		}
		if !policy.IgnoreWarnings && warnings > policy.MaxWarningsPerTier {
			return i
		}
	}
	return len(tiers) - 1
}

func tierRank(level string) int {
	for i, t := range tiers {
		if t == level {
			return i
		}
	}
	return len(tiers)
}

func findingLess(a, b engine.Finding) bool {
	if (a.Status == engine.StatusPass) != (b.Status == engine.StatusPass) {
		return a.Status != engine.StatusPass
	}
	if tierRank(a.MaturityLevel) != tierRank(b.MaturityLevel) {
		return tierRank(a.MaturityLevel) < tierRank(b.MaturityLevel)
	}
	if sevRank[a.Severity] != sevRank[b.Severity] {
		return sevRank[a.Severity] < sevRank[b.Severity]
	}
	return a.ID < b.ID
}

// buildRoadmap groups non-passing rules by the maturity tier they block, then
// orders each tier by severity and blast radius.
func buildRoadmap(fs []engine.Finding, pillarList []pillarInfo) []htmlTier {
	urls := map[string]string{}
	for _, p := range pillarList {
		urls[p.Name] = p.URL
	}
	byRule := map[string]*htmlGap{}
	tierOf := map[string]int{}
	var order []string
	for _, f := range fs {
		if f.Status == engine.StatusPass {
			continue
		}
		g, ok := byRule[f.ID]
		if !ok {
			pillar := pillarOf(f)
			g = &htmlGap{
				ID: f.ID, Title: f.Title, Severity: f.Severity, Pillar: pillar,
				PillarURL: urls[pillar], Scope: f.Scope, FixHint: f.FixHint,
			}
			byRule[f.ID] = g
			tierOf[f.ID] = tierRank(f.MaturityLevel)
			order = append(order, f.ID)
		}
		if f.Status == engine.StatusFail {
			g.Failing = append(g.Failing, f.Target)
		} else {
			g.Unknown = append(g.Unknown, f.Target)
		}
	}

	out := make([]htmlTier, len(tiers)+1)
	for i := range tiers {
		out[i].Heading = levelNames[i] + " checks"
		if i+1 < len(levelNames) {
			out[i].Heading += " — fix to reach " + levelNames[i+1]
		}
	}
	out[len(tiers)].Heading = "Other gaps"
	for _, id := range order {
		out[tierOf[id]].Gaps = append(out[tierOf[id]].Gaps, *byRule[id])
	}
	var nonEmpty []htmlTier
	for _, t := range out {
		if len(t.Gaps) == 0 {
			continue
		}
		sort.SliceStable(t.Gaps, func(a, b int) bool {
			ga, gb := t.Gaps[a], t.Gaps[b]
			if sevRank[ga.Severity] != sevRank[gb.Severity] {
				return sevRank[ga.Severity] < sevRank[gb.Severity]
			}
			if len(ga.Failing) != len(gb.Failing) {
				return len(ga.Failing) > len(gb.Failing)
			}
			return ga.ID < gb.ID
		})
		nonEmpty = append(nonEmpty, t)
	}
	return nonEmpty
}
