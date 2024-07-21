package jsondefs

import (
	"github.com/naivary/specraft/definitions"
	jsonschm "github.com/naivary/specraft/schema/json"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

const (
	validationPrefix = "jsonschema:validation:"
	metaPrefix       = "jsonschema:meta:"
)

var FieldMarkers = definitions.MustMakeAllWithPrefix(validationPrefix, markers.DescribesField,
	// numeric markers
	Maximum(0),
	Minimum(0),

	// array markers
	// string markers
	// object markers

)

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

var TypeMarkers = definitions.MustMakeAllWithPrefix(metaPrefix, markers.DescribesType,
	ID(""),
	Draft(""),
)

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

func (id ID) ApplyToSchema(s *jsonschm.Schema) error {
    s.ID = string(id)
    return nil
}


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

func (d Draft) ApplyToSchema(s *jsonschm.Schema) error {
    s.Schema = string(d)
    return nil
}
