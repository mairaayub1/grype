package main

import (
	"fmt"
	"os"

	"github.com/dave/jennifer/jen"

	"github.com/mairaayub1/grype/grype/internal/packagemetadata"
)

// This program is invoked from grype/internal and generates packagemetadata/generated.go

const (
	pkgImport = "github.com/mairaayub1/grype/grype/pkg"
	path      = "packagemetadata/generated.go"
)

func main() {
	typeNames, err := packagemetadata.DiscoverTypeNames()
	if err != nil {
		panic(fmt.Errorf("unable to get all metadata type names: %w", err))
	}

	fmt.Printf("updating package metadata type list with %+v types\n", len(typeNames))

	f := jen.NewFile("packagemetadata")
	f.HeaderComment("DO NOT EDIT: generated by grype/internal/packagemetadata/generate/main.go")
	f.ImportName(pkgImport, "pkg")
	f.Comment("AllTypes returns a list of all pkg metadata types that grype supports (that are represented in the pkg.Package.Metadata field).")

	f.Func().Id("AllTypes").Params().Index().Any().BlockFunc(func(g *jen.Group) {
		g.ReturnFunc(func(g *jen.Group) {
			g.Index().Any().ValuesFunc(func(g *jen.Group) {
				for _, typeName := range typeNames {
					g.Qual(pkgImport, typeName).Values()
				}
			})
		})
	})

	rendered := fmt.Sprintf("%#v", f)

	fh, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(fmt.Errorf("unable to open file: %w", err))
	}
	_, err = fh.WriteString(rendered)
	if err != nil {
		panic(fmt.Errorf("unable to write file: %w", err))
	}
	if err := fh.Close(); err != nil {
		panic(fmt.Errorf("unable to close file: %w", err))
	}
}
