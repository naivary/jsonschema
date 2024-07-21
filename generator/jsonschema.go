package generator

import "sigs.k8s.io/controller-tools/pkg/markers"

func NewJSONSchemaGenerator() Generator {
    return jsonSchemaGenerator{}
}

var _ Generator = (*jsonSchemaGenerator)(nil)

type jsonSchemaGenerator struct {}

func (j jsonSchemaGenerator) Generate(infos []*markers.TypeInfo) error {
    return nil
}
