package namespace

import (
	"github.com/mairaayub1/grype/grype/db/v5/pkg/resolver"
)

type Namespace interface {
	Provider() string
	Resolver() resolver.Resolver
	String() string
}
