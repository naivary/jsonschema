package schema

import (
	"go/ast"

	"golang.org/x/tools/go/packages"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

type Schema[T any] interface {
	TypeOf(field *ast.Field, pkg *packages.Package) T

	NameForField(info *markers.FieldInfo) string
}
