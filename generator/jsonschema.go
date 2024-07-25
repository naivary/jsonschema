package generator

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/naivary/specraft/definer"
	"github.com/naivary/specraft/schema"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

func JSONSchema() Generator[*schema.JSON] {
	return jsonSchemaGenerator{}
}

var _ Generator[*schema.JSON] = (*jsonSchemaGenerator)(nil)

type jsonSchemaGenerator struct{}

func (j jsonSchemaGenerator) Generate(defn definer.Definer[*schema.JSON], pkg *loader.Package, typeInfo *markers.TypeInfo, w io.Writer) error {
	t := pkg.TypesInfo.TypeOf(typeInfo.RawSpec.Type)
	if !schema.IsStructType(t) {
		return nil
	}

	schm := &schema.JSON{
		Description: typeInfo.Doc,
		Title:       typeInfo.Name,
		Type:        schema.JSONTypeOf(typeInfo.RawSpec.Type, pkg),
		Properties:  make(map[string]*schema.JSON),
	}

	for name, val := range typeInfo.Markers {
		err := defn.ApplierForMarker(name, val).ApplyToSchema(schm)
		if err != nil {
			return err
		}
	}

	for _, fieldInfo := range typeInfo.Fields {
		fieldSchm := &schema.JSON{
			Description: fieldInfo.Doc,
			Type:        schema.JSONTypeOf(fieldInfo.RawField.Type, pkg),
		}
		if fieldSchm.IsInvalidType() {
			return fmt.Errorf("invalid JSON Type for field: %s", fieldInfo.Name)
		}
		for name, val := range fieldInfo.Markers {
			err := defn.ApplierForMarker(name, val).ApplyToSchema(fieldSchm)
			if err != nil {
				return err
			}
		}
		name := schema.JSONNameForField(&fieldInfo)
		schm.Properties[name] = fieldSchm
	}
	return json.NewEncoder(w).Encode(schm)
}
