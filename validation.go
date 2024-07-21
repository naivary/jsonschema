package main

import "sigs.k8s.io/controller-tools/pkg/markers"

const (
	validationPrefix = "jsonschema:validation:"
)

var ValidationMarkers, err = mustMakeAllWithPrefix(validationPrefix, markers.DescribesField, 
    // numeric markers
    Max(0),
    Min(0),
)

type Max float64

func (m Max) Value() float64 {
    return float64(m)
}

func (Max) Help() *markers.DefinitionHelp {
    return &markers.DefinitionHelp{
        Category: "validation",
        DetailedHelp: markers.DetailedHelp{
            Summary: "maximum integer",
            Details: "maximum integer which the specified has to be",
        },
        FieldHelp: map[string]markers.DetailedHelp{},
    }
}

type Min float64

func (m Min) Value() float64 {
    return float64(m)
}

func (Min) Help() *markers.DefinitionHelp {
    return &markers.DefinitionHelp{
        Category: "validation",
        DetailedHelp: markers.DetailedHelp{
            Summary: "minimum integer",
            Details: "minimum integer which the specified has to be",
        },
        FieldHelp: map[string]markers.DetailedHelp{},
    }
}
