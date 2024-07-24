package definer

import (
	"errors"
	"strings"

	"github.com/naivary/specraft/schema"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

const (
	JSONSchemaValidationPrefix = "jsonschema:validation:"
	JSONSChemaMetaPrefix       = "jsonschema:meta:"
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

func (j *jsonSchemaDefiner) Define(defs ...*DefinitionWithHelp) error {
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
	return nil
}

func (j *jsonSchemaDefiner) Registry() *markers.Registry {
	return j.reg
}

func (j *jsonSchemaDefiner) ApplierFor(marker string, val []any) Applier[*schema.JSON] {
	name := strings.TrimPrefix(marker, JSONSchemaValidationPrefix)
	name = strings.TrimPrefix(name, JSONSChemaMetaPrefix)
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
			Prefix:     JSONSChemaMetaPrefix,
			TargetType: markers.DescribesType,
			Objs: []any{
				ID(""),
				Draft(""),
			},
		},
	}

	for _, ms := range markerSets {
		defs, err := makeAllWithPrefix(ms.Prefix, ms.TargetType, ms.Objs...)
		if err != nil {
			return err
		}
		err = j.Define(defs...)
		if err != nil {
			return err
		}
	}
	return nil
}

var _ Applier[*schema.JSON] = (*Maximum)(nil)

type Maximum float64

func (m Maximum) Value() float64 {
	return float64(m)
}

func (Maximum) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "maximum integer",
			Details: "maximum integer which the specified has to be",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (m Maximum) ApplyToSchema(s *schema.JSON) error {
	if s.Type != schema.JSONTypeNumber {
		return errors.New("cannot apply `Maximum` to non numeric type")
	}
	val := int(m)
	s.Maximum = &val
	return nil
}

var _ Applier[*schema.JSON] = (*Minimum)(nil)

type Minimum float64

func (m Minimum) Value() float64 {
	return float64(m)
}

func (Minimum) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "JSON Schema validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "minimum integer",
			Details: "minimum integer which the specified has to be",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (m Minimum) ApplyToSchema(s *schema.JSON) error {
	if s.Type != schema.JSONTypeNumber {
		return errors.New("cannot apply `Minimum` to non numeric type")
	}
	val := int(m)
	s.Minimum = &val
	return nil
}

var _ Applier[*schema.JSON] = (*ID)(nil)

type ID string

func (i ID) Value() string {
	return string(i)
}

func (ID) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "JSON Schema validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "id of the JSON Schema",
			Details: "id of the JSON Schema used for cross-referencing",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (id ID) ApplyToSchema(s *schema.JSON) error {
	s.ID = string(id)
	return nil
}

var _ Applier[*schema.JSON] = (*Draft)(nil)

type Draft string

func (d Draft) Value() string {
	return string(d)
}

func (Draft) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "JSON Schema validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "JSON Schema draft to use",
			Details: "specified JSON Schema draf for this type to use",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (d Draft) ApplyToSchema(s *schema.JSON) error {
	s.Draft = string(d)
	return nil
}
