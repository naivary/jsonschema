package definer

import (
	"reflect"

	"sigs.k8s.io/controller-tools/pkg/markers"
)

type Prefix = string

type Helper interface {
	Help() *markers.DefinitionHelp
}

type Applier[T any] interface {
	ApplyToSchema(schm T) error
}

type Definer[T any] interface {
	Define(def ...*DefinitionWithHelp) error

	Prefixes() map[Prefix]markers.TargetType

	Registry() *markers.Registry

	ApplierFor(marker string, val []any) Applier[T]
}

type MarkerSet struct {
	Prefix     string
	TargetType markers.TargetType
	Objs       []any
}

type DefinitionWithHelp struct {
	*markers.Definition
	Help *markers.DefinitionHelp
}

func (d *DefinitionWithHelp) WithHelp(help *markers.DefinitionHelp) *DefinitionWithHelp {
	d.Help = help
	return d
}

func (d *DefinitionWithHelp) Register(reg *markers.Registry) error {
	if err := reg.Register(d.Definition); err != nil {
		return err
	}
	if d.Help != nil {
		reg.AddHelp(d.Definition, d.Help)
	}
	return nil
}

func makeAllWithPrefix(prefix string, target markers.TargetType, objs ...any) ([]*DefinitionWithHelp, error) {
	defs := make([]*DefinitionWithHelp, len(objs))
	for i, obj := range objs {
		name := prefix + reflect.TypeOf(obj).Name()
		def, err := markers.MakeDefinition(name, target, obj)
		if err != nil {
			return nil, err
		}
		defs[i] = &DefinitionWithHelp{Definition: def, Help: obj.(Helper).Help()}
	}
	return defs, nil
}
