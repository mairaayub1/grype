package matcher

import (
	syftPkg "github.com/anchore/syft/syft/pkg"
	v5 "github.com/mairaayub1/grype/grype/db/v5"
	"github.com/mairaayub1/grype/grype/distro"
	"github.com/mairaayub1/grype/grype/match"
	"github.com/mairaayub1/grype/grype/pkg"
)

type Matcher interface {
	PackageTypes() []syftPkg.Type
	Type() match.MatcherType
	Match(v5.VulnerabilityProvider, *distro.Distro, pkg.Package) ([]match.Match, error)
}
