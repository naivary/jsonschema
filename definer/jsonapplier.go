package definer

import (
	"errors"

	"github.com/naivary/specraft/schema"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

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
