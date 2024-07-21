package generator

import (
	"encoding/json"
	"errors"
	"fmt"
	"go/ast"
	"io"

	"github.com/naivary/specraft/schema"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

func NewJSONSchema() Generator {
	return jsonSchemaGenerator{}
}

var _ Generator = (*jsonSchemaGenerator)(nil)

type jsonSchemaGenerator struct{}

func (j jsonSchemaGenerator) Generate(info *markers.TypeInfo, w io.Writer) error {
	if info.Fields == nil {
		return errors.New("empty struct")
	}

	schm := schema.JSONSchema{
		ID:     "test-id",
		Schema: "test-schema",
		Type:   schema.ObjectType,
		Title:  info.Name,
	}

	properties := make(map[string]schema.Property)
	for _, field := range info.Fields {
		// TODO(naivary): embedded types are still not handled
		prop := schema.Property{}
		typ := schema.TypeOf(field.RawField.Type.(*ast.Ident).String())
		if typ == schema.InvalidType {
			return errors.New("invalid json type")
		}
        fmt.Println(field.Markers)

		prop.Type = typ.String()
		prop.Description = field.Doc
		properties[field.Name] = prop
	}

	schm.Properties = properties
	return json.NewEncoder(w).Encode(&schm)
}
