package generator

import (
	"encoding/json"
	"errors"
	"io"

	jsonschm "github.com/naivary/specraft/schema/json"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

func NewJSONSchema() Generator {
	return jsonSchemaGenerator{}
}

var _ Generator = (*jsonSchemaGenerator)(nil)

type jsonSchemaGenerator struct{}

// 1. take one Field
// 2. create a new property for that field with the correct type
// 3. set all the defined markers on the property
// 4. add the property to the Schema.
func (j jsonSchemaGenerator) Generate(info *markers.TypeInfo, w io.Writer) error {
	if info.Fields == nil {
		return errors.New("empty struct")
	}
	schm := jsonschm.JSONSchema{
        Type: jsonschm.ObjectType,
        Description: info.Doc,
        Title: info.Name,
    }
	_ = schm
	return json.NewEncoder(w).Encode(nil)
}
