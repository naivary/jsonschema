package runtime

import (
	"github.com/naivary/specraft/schema"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

func NewRuntime(pkgs map[*loader.Package][]*markers.TypeInfo) Runtime[schema.JSONType, *schema.JSON] {
	return &jsonSchemaRuntime{pkgs}
}

var _ Runtime[schema.JSONType, *schema.JSON] = (*jsonSchemaRuntime)(nil)

type jsonSchemaRuntime struct {
	pkgs map[*loader.Package][]*markers.TypeInfo
}

func (jrt jsonSchemaRuntime) NameForField(info *markers.FieldInfo) string {
	return ""
}

func (jrt jsonSchemaRuntime) Packages() map[*loader.Package][]*markers.TypeInfo {
	return jrt.pkgs
}
