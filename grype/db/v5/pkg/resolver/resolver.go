package resolver

import (
	grypePkg "github.com/mairaayub1/grype/grype/pkg"
)

type Resolver interface {
	Normalize(string) string
	Resolve(p grypePkg.Package) []string
}
