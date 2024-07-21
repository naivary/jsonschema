package main

import (
	"fmt"
	"log/slog"
	"os"

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
	for _, def := range ValidationMarkers {
		if err := reg.Register(def.Definition); err != nil {
			return err
		}
		reg.AddHelp(def.Definition, def.Help)
	}

	// collect all the markers in the given project
	col := &markers.Collector{
		Registry: reg,
	}
	pkgs, err := loader.LoadRoots("examples/auth_req.go")
	if err != nil {
		return err
	}

	infos := make([]*markers.TypeInfo, 0)
	for _, pkg := range pkgs {
		err := markers.EachType(col, pkg, func(info *markers.TypeInfo) {
			infos = append(infos, info)
		})
		if err != nil {
			return err
		}
	}
	fmt.Println(infos[0].Fields)
	return nil
}
