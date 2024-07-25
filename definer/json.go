package definer

import (
	"strings"

	"github.com/naivary/specraft/schema"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

const (
	JSONSchemaValidationPrefix = "jsonschema:validation:"
	JSONSchemaMetaPrefix       = "jsonschema:meta:"
)

var _ Definer[*schema.JSON] = (*jsonSchemaDefiner)(nil)

type jsonSchemaDefiner struct {
	reg *markers.Registry
}

func JSONSchema(reg *markers.Registry) (Definer[*schema.JSON], error) {
	definer := &jsonSchemaDefiner{
		reg: reg,
	}
	err := definer.init()
	return definer, err
}

func (j *jsonSchemaDefiner) define(defs ...*DefinitionWithHelp) error {
	for _, def := range defs {
		err := j.reg.Register(def.Definition)
		if err != nil {
			return err
		}
		j.reg.AddHelp(def.Definition, def.Help)
	}
	return nil
}

func (j *jsonSchemaDefiner) Prefixes() map[Prefix]markers.TargetType {
	return map[Prefix]markers.TargetType{
		JSONSchemaValidationPrefix: markers.DescribesField,
		JSONSchemaMetaPrefix:       markers.DescribesType,
	}
}


func (j *jsonSchemaDefiner) ApplierForMarker(marker string, val []any) Applier[*schema.JSON] {
	name := strings.TrimPrefix(marker, JSONSchemaValidationPrefix)
	name = strings.TrimPrefix(name, JSONSchemaMetaPrefix)
	switch name {
	case "ID":
		return (val[0]).(ID)
	case "Draft":
		return (val[0]).(Draft)
	case "Maximum":
		return (val[0]).(Maximum)
	case "Minimum":
		return (val[0]).(Minimum)
	default:
		return nil
	}
}

func (j *jsonSchemaDefiner) init() error {
	markerSets := []MarkerSet{
		{
			Prefix:     JSONSchemaValidationPrefix,
			TargetType: markers.DescribesField,
			Objs: []any{
				// numeric markers
				Maximum(0),
				Minimum(0),
			},
		},
		{
			Prefix:     JSONSchemaMetaPrefix,
			TargetType: markers.DescribesType,
			Objs: []any{
				ID(""),
				Draft(""),
			},
		},
	}

	for _, ms := range markerSets {
		defs, err := makeDefsWithPrefix(ms.Prefix, ms.TargetType, ms.Objs...)
		if err != nil {
			return err
		}
		if err := j.define(defs...); err != nil {
			return err
		}
	}
	return nil
}
