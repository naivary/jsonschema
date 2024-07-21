package main

import "sigs.k8s.io/controller-tools/pkg/markers"

const (
	validationPrefix = "jsonschema:validation:"
)

var ValidationMarkers, err = mustMakeAllWithPrefix(validationPrefix, markers.DescribesField,
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
		Category: "validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "minimum integer",
			Details: "minimum integer which the specified has to be",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}
