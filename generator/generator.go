package generator

import (
	"io"

	"github.com/naivary/specraft/runtime"
)

type Generator[T any, S any] interface {
	// Generate will be given all the defined types with set marker comments
	Generate(rt runtime.Runtime[T, S], w io.Writer) error
}
