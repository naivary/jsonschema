package jsonschema

import "sigs.k8s.io/controller-tools/pkg/markers"

const (
	validationPrefix = "jsonschema:validation:"
)

var ValidationMarkers, err = mustMakeAllWithPrefix(validationPrefix, markers.DescribesField, 

    Maximum(0),)



// +controllertools:marker:generateHelp:category="validation"
// Minimum specifies the minimum numeric value that this field can have. Negative numbers are supported.
type Maximum float64

func (m Maximum) Value() float64 {
    return float64(m)
}
