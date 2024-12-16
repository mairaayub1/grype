package qualifier

import (
	"fmt"

	"github.com/mairaayub1/grype/grype/pkg/qualifier"
)

type Qualifier interface {
	fmt.Stringer
	Parse() qualifier.Qualifier
}
