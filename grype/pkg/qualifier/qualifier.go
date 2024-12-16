package qualifier

import (
	"github.com/mairaayub1/grype/grype/distro"
	"github.com/mairaayub1/grype/grype/pkg"
)

type Qualifier interface {
	Satisfied(d *distro.Distro, p pkg.Package) (bool, error)
}
