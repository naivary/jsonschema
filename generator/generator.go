package generator

import "sigs.k8s.io/controller-tools/pkg/markers"

type Generator interface {
    // Generate will be given all the defined types
    // with set marker comments
    Generate(infos []*markers.TypeInfo) error
}
