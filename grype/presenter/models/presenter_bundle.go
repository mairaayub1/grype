package models

import (
	"github.com/anchore/clio"
	"github.com/anchore/syft/syft/sbom"
	"github.com/mairaayub1/grype/grype/match"
	"github.com/mairaayub1/grype/grype/pkg"
	"github.com/mairaayub1/grype/grype/vulnerability"
)

type PresenterConfig struct {
	ID               clio.Identification
	Matches          match.Matches
	IgnoredMatches   []match.IgnoredMatch
	Packages         []pkg.Package
	Context          pkg.Context
	MetadataProvider vulnerability.MetadataProvider
	SBOM             *sbom.SBOM
	AppConfig        interface{}
	DBStatus         interface{}
}
