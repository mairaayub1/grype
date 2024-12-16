package dotnet

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

func NewDotnetMatcher(cfg MatcherConfig) *Matcher {
	return &Matcher{
		cfg: cfg,
	}
}

func (m *Matcher) PackageTypes() []syftPkg.Type {
	return []syftPkg.Type{syftPkg.DotnetPkg}
}

func (m *Matcher) Type() match.MatcherType {
	return match.DotnetMatcher
}

func (m *Matcher) Match(store v5.VulnerabilityProvider, d *distro.Distro, p pkg.Package) ([]match.Match, error) {
	criteria := search.CommonCriteria
	if m.cfg.UseCPEs {
		criteria = append(criteria, search.ByCPE)
	}
	return search.ByCriteria(store, d, p, m.Type(), criteria...)
}
