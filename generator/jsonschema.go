package generator

import (
	"fmt"
	"io"
	"reflect"

	"github.com/naivary/specraft/runtime"
	"github.com/naivary/specraft/schema"
)

func JSONSchema() Generator[schema.JSONType, *schema.JSON] {
    return jsonSchemaGenerator{}
}

var _ Generator[schema.JSONType, *schema.JSON] = (*jsonSchemaGenerator)(nil)

type jsonSchemaGenerator struct{}

// 1. take one Field
// 2. create a new property for that field with the correct type
// 3. set all the defined markers on the property
// 4. add the property to the Schema.
func (j jsonSchemaGenerator) Generate(rt runtime.Runtime[schema.JSONType, *schema.JSON], w io.Writer) error {
	pkgs := rt.Packages()
	for pkg, infos := range pkgs {
		for _, info := range infos {
			for _, field := range info.Fields {
				typ := pkg.TypesInfo.Types[field.RawField.Type].Type
				fmt.Println(reflect.TypeOf(typ))
			}
		}
	}
	return nil
}
