package main

import (
	"log/slog"
	"os"

	"github.com/naivary/specraft/definer"
	"github.com/naivary/specraft/generator"
	"github.com/naivary/specraft/runtime"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

func main() {
	if err := run(); err != nil {
		slog.Error("something went wrong", "err_msg", err.Error())
		os.Exit(1)
	}
}

func run() error {
	reg := &markers.Registry{}
	defn, err := definer.JSONSchema(reg)
	if err != nil {
		return err
	}

	// collect all the markers in the given project
    // TODO(naivary) should be handled probably by the runtime
	col := &markers.Collector{
		Registry: defn.Registry(),
	}
	pkgs, err := loader.LoadRoots("examples/auth_req.go")
	if err != nil {
		return err
	}

	pkgInfos := make(map[*loader.Package][]*markers.TypeInfo, 0)
	for _, pkg := range pkgs {
		pkg.NeedTypesInfo()
		pkg.NeedSyntax()
		infos := make([]*markers.TypeInfo, 0)
		err := markers.EachType(col, pkg, func(info *markers.TypeInfo) {
			infos = append(infos, info)
		})
		if err != nil {
			return err
		}
		pkgInfos[pkg] = infos

	}

	rt := runtime.NewRuntime(pkgInfos)
	return generator.JSONSchema().Generate(rt, os.Stdout)
}
