package generator

import (
	"errors"
	"fmt"
	"go/types"

	"github.com/naivary/specraft/definer"
	"github.com/naivary/specraft/schema"
	"github.com/naivary/specraft/utils/typesutil"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

var (
	ErrNonStructType = errors.New("type is not a struct")
)

func JSONSchema() Generator[*schema.JSON] {
	return jsonSchemaGenerator{
		tc: schema.NewJSONSchemaTypeConvert(),
	}
}

var _ Generator[*schema.JSON] = (*jsonSchemaGenerator)(nil)

type jsonSchemaGenerator struct {
	tc schema.TypeConverter[schema.JSONType]
}

func (j jsonSchemaGenerator) Generate(defn definer.Definer[*schema.JSON], pkg *loader.Package, typeInfo *markers.TypeInfo) (*schema.JSON, error) {
	typeType := pkg.TypesInfo.TypeOf(typeInfo.RawSpec.Type)
	if _, ok := typesutil.IsType[*types.Struct](typeType); !ok {
		return nil, ErrNonStructType
	}

	typeSchm := &schema.JSON{
		// TODO(naivary): set a sane id
		ID:          "default-id",
		Description: typeInfo.Doc,
		Title:       typeInfo.Name,
		Type:        typesutil.TypeConversion(typeType, j.tc),
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
		if _, ok := typesutil.IsType[types.Struct](fieldType); ok {
			// make a reference to the otherhema of the
		}
		fieldSchm := &schema.JSON{
			Description: fieldInfo.Doc,
			Type:        typesutil.TypeConversion(fieldType, j.tc),
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
