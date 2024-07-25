package generator

import (
	"errors"
	"io"

	"github.com/naivary/specraft/definer"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

var (
	ErrInvalidType = errors.New("invalid type")
)

type Generator[T any] interface {
	// Generate will be given all the defined types with set marker comments
	Generate(defn definer.Definer[T], pkg *loader.Package, typeInfo *markers.TypeInfo, w io.Writer) error
}
