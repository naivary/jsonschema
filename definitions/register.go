package definitions

import (
	"reflect"

	"sigs.k8s.io/controller-tools/pkg/markers"
)

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

func (d *DefinitionWithHelp) clone() *DefinitionWithHelp {
	newDef, newHelp := *d.Definition, *d.Help
	return &DefinitionWithHelp{
		Definition: &newDef,
		Help:       &newHelp,
	}
}

type Helper interface {
	Help() *markers.DefinitionHelp
}

func MustMakeAllWithPrefix(prefix string, target markers.TargetType, objs ...any) []*DefinitionWithHelp {
	defs := make([]*DefinitionWithHelp, len(objs))
	for i, obj := range objs {
		name := prefix + reflect.TypeOf(obj).Name()
		def, err := markers.MakeDefinition(name, target, obj)
		if err != nil {
			panic(err)
		}
		defs[i] = &DefinitionWithHelp{Definition: def, Help: obj.(Helper).Help()}
	}
	return defs
}
