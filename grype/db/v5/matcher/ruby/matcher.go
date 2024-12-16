package ruby

import (
	syftPkg "github.com/anchore/syft/syft/pkg"
	v5 "github.com/mairaayub1/grype/grype/db/v5"
	"github.com/mairaayub1/grype/grype/db/v5/search"
	"github.com/mairaayub1/grype/grype/distro"
	"github.com/mairaayub1/grype/grype/match"
	"github.com/mairaayub1/grype/grype/pkg"
)

type Matcher struct {
	cfg MatcherConfig
}

type MatcherConfig struct {
	UseCPEs bool
}

func NewRubyMatcher(cfg MatcherConfig) *Matcher {
	return &Matcher{
		cfg: cfg,
	}
}

func (m *Matcher) PackageTypes() []syftPkg.Type {
	return []syftPkg.Type{syftPkg.GemPkg}
}

func (m *Matcher) Type() match.MatcherType {
	return match.RubyGemMatcher
}

func (m *Matcher) Match(store v5.VulnerabilityProvider, d *distro.Distro, p pkg.Package) ([]match.Match, error) {
	criteria := search.CommonCriteria
	if m.cfg.UseCPEs {
		criteria = append(criteria, search.ByCPE)
	}
	return search.ByCriteria(store, d, p, m.Type(), criteria...)
}
