package runtime

import (
	"io"
	"os"

	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

type Runtime[T any] interface {
	// TODO(naivary): should have the following components:
	// 1. definer,
	// 2. global registry,
	// 3. generator for each type to generate
	// 4. collector
	// 5. Packages

	// Package is returning the package which this runtime is made for
	Packages() map[*loader.Package][]*markers.TypeInfo

	Schemas() map[string]*os.File

	Generate() error
}

// internal runtime to use for the generators
// TODO(naivary): how to setup the generator grpc server? 
type runtime interface {
	Packages() []*loader.Package
	Registry() *markers.Registry
	Files() map[string]io.Reader
}

type Generator interface {
	Generate(pkg *loader.Package) (map[string]io.Reader, error)

	Markers() []MarkerSet
}

type Creater interface {
	// --creater=local
	// --creater.dir = something
	// eigene flag set
	Create(name string, r io.Reader) error
}

type MarkerSet struct {
	Prefix     string
	TargetType markers.TargetType
	Objs       []any
}
