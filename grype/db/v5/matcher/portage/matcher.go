package portage

import (
	"fmt"

	syftPkg "github.com/anchore/syft/syft/pkg"
	v5 "github.com/mairaayub1/grype/grype/db/v5"
	"github.com/mairaayub1/grype/grype/db/v5/search"
	"github.com/mairaayub1/grype/grype/distro"
	"github.com/mairaayub1/grype/grype/match"
	"github.com/mairaayub1/grype/grype/pkg"
)

type Matcher struct {
}

func (m *Matcher) PackageTypes() []syftPkg.Type {
	return []syftPkg.Type{syftPkg.PortagePkg}
}

func (m *Matcher) Type() match.MatcherType {
	return match.PortageMatcher
}

func (m *Matcher) Match(store v5.VulnerabilityProvider, d *distro.Distro, p pkg.Package) ([]match.Match, error) {
	matches, err := search.ByPackageDistro(store, d, p, m.Type())
	if err != nil {
		return nil, fmt.Errorf("failed to find vulnerabilities: %w", err)
	}

	return matches, nil
}
