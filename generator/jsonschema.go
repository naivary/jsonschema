package generator

import (
	"errors"
	"fmt"

	"github.com/naivary/specraft/definer"
	"github.com/naivary/specraft/schema"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

var (
	ErrNonStructType = errors.New("type is not a struct")
)

func JSONSchema() Generator[*schema.JSON] {
	return jsonSchemaGenerator{}
}

var _ Generator[*schema.JSON] = (*jsonSchemaGenerator)(nil)

type jsonSchemaGenerator struct{}

func (j jsonSchemaGenerator) Generate(defn definer.Definer[*schema.JSON], pkg *loader.Package, typeInfo *markers.TypeInfo) (*schema.JSON, error) {
	typeType := pkg.TypesInfo.TypeOf(typeInfo.RawSpec.Type)
	if !schema.IsStructType(typeType) {
		return nil, ErrNonStructType
	}

	typeSchm := &schema.JSON{
		// TODO(naivary): set a sane id
		ID:          "default-id",
		Description: typeInfo.Doc,
		Title:       typeInfo.Name,
		Type:        schema.JSONTypeOf(typeType),
		Properties:  make(map[string]*schema.JSON),
	}

	for name, val := range typeInfo.Markers {
		req := definer.NewTypeApplyRequest(typeInfo, typeSchm)
		err := defn.ApplierForMarker(name, val).Apply(req, nil)
		if err != nil {
			return nil, err
		}
	}

	for _, fieldInfo := range typeInfo.Fields {
		fieldType := pkg.TypesInfo.TypeOf(fieldInfo.RawField.Type)
		if schema.IsStructType(fieldType) {
			// make a reference to the otherhema of the
		}
		fieldSchm := &schema.JSON{
			Description: fieldInfo.Doc,
			Type:        schema.JSONTypeOf(fieldType),
		}
		if fieldSchm.IsInvalidType() {
			return nil, fmt.Errorf("invalid JSON Type for field: %s", fieldInfo.Name)
		}

		typeReq := definer.NewTypeApplyRequest(typeInfo, typeSchm)
		for name, val := range fieldInfo.Markers {
			fieldReq := definer.NewFieldApplyRequest(&fieldInfo, fieldSchm)
			err := defn.ApplierForMarker(name, val).Apply(typeReq, fieldReq)
            if err != nil {
				return nil, err
			}
		}
		name := schema.JSONNameForField(&fieldInfo)
		typeSchm.Properties[name] = fieldSchm
	}
	return typeSchm, nil
}
