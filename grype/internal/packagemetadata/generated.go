// DO NOT EDIT: generated by grype/internal/packagemetadata/generate/main.go

package packagemetadata

import "github.com/mairaayub1/grype/grype/pkg"

// AllTypes returns a list of all pkg metadata types that grype supports (that are represented in the pkg.Package.Metadata field).
func AllTypes() []any {
	return []any{pkg.ApkMetadata{}, pkg.GolangBinMetadata{}, pkg.GolangModMetadata{}, pkg.JavaMetadata{}, pkg.JavaVMInstallationMetadata{}, pkg.RpmMetadata{}}
}
