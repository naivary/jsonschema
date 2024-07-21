package jsonschema

import (
	"reflect"

	"sigs.k8s.io/controller-tools/pkg/markers"
)

type definitionWithHelp struct {
	*markers.Definition
	Help *markers.DefinitionHelp
}

func (d *definitionWithHelp) WithHelp(help *markers.DefinitionHelp) *definitionWithHelp {
	d.Help = help
	return d
}

func (d *definitionWithHelp) Register(reg *markers.Registry) error {
	if err := reg.Register(d.Definition); err != nil {
		return err
	}
	if d.Help != nil {
		reg.AddHelp(d.Definition, d.Help)
	}
	return nil
}

func (d *definitionWithHelp) clone() *definitionWithHelp {
	newDef, newHelp := *d.Definition, *d.Help
	return &definitionWithHelp{
		Definition: &newDef,
		Help:       &newHelp,
	}
}

// AllDefinitions contains all marker definitions for this package.
var AllDefinitions []*definitionWithHelp

type Helper interface {
	Help() *markers.DefinitionHelp
}

func mustMakeAllWithPrefix(prefix string, target markers.TargetType, objs ...any) ([]*definitionWithHelp, error) {
	defs := make([]*definitionWithHelp, len(objs))
	for i, obj := range objs {
		name := prefix + reflect.TypeOf(obj).Name()
		def, err := markers.MakeDefinition(name, target, obj)
		if err != nil {
			return nil, err
		}
		defs[i] = &definitionWithHelp{Definition: def, Help: obj.(Helper).Help()}
	}
	return defs, nil
}
