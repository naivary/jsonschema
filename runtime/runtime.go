package runtime

import (
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
