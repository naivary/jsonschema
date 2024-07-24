package runtime

import (
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

type Runtime[T any, S any] interface {
	// TODO(naivary): should have definer, global registry, generator for each
	// type to generate

	// Package is returning the package which this runtime is made for
	Packages() map[*loader.Package][]*markers.TypeInfo
}
