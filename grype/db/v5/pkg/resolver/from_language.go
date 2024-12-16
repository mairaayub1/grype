package resolver

import (
	syftPkg "github.com/anchore/syft/syft/pkg"
	"github.com/mairaayub1/grype/grype/db/v5/pkg/resolver/java"
	"github.com/mairaayub1/grype/grype/db/v5/pkg/resolver/python"
	"github.com/mairaayub1/grype/grype/db/v5/pkg/resolver/stock"
)

func FromLanguage(language syftPkg.Language) (Resolver, error) {
	var r Resolver

	switch language {
	case syftPkg.Python:
		r = &python.Resolver{}
	case syftPkg.Java:
		r = &java.Resolver{}
	default:
		r = &stock.Resolver{}
	}

	return r, nil
}
