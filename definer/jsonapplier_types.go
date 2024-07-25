package definer

import (
	"errors"
	"regexp"

	"github.com/naivary/specraft/schema"
)

var (
	ErrFieldSchema = errors.New("applier is only appliable to the top level schema")
)

var _ Applier[*schema.JSON] = (*ExclusiveMaximum)(nil)

type ExclusiveMaximum bool

func (e ExclusiveMaximum) Apply(typeReq *TypeApplyRequest[*schema.JSON], fieldReq *FieldApplyRequest[*schema.JSON]) error {
	s := fieldReq.Schema
	if s.Type != schema.JSONTypeNumber {
		return errors.New("cannot apply `ExclusiveMaximum` to non numeric field")
	}
	if s.Maximum == nil {
		return errors.New("missing `Maximum` marker")
	}
	s.ExclusiveMaximum = true
	return nil
}

var _ Applier[*schema.JSON] = (*ExclusiveMinimum)(nil)

type ExclusiveMinimum bool

func (e ExclusiveMinimum) Apply(typeReq *TypeApplyRequest[*schema.JSON], fieldReq *FieldApplyRequest[*schema.JSON]) error {
	s := fieldReq.Schema
	if s.Type != schema.JSONTypeNumber {
		return errors.New("cannot apply `ExclusiveMinimum` to non numeric field")
	}
	if s.Minimum == nil {
		return errors.New("missing `Minimum` marker")
	}
	s.ExclusiveMinimum = true
	return nil
}

var _ Applier[*schema.JSON] = (*MultipleOf)(nil)

type MultipleOf float64

func (m MultipleOf) Apply(typeReq *TypeApplyRequest[*schema.JSON], fieldReq *FieldApplyRequest[*schema.JSON]) error {
	s := fieldReq.Schema
	if s.Type != schema.JSONTypeNumber {
		return errors.New("cannot apply `MultipleOf` to non numeric field")
	}
	val := float64(m)
	s.MultipleOf = &val
	return nil
}

var _ Applier[*schema.JSON] = (*Maximum)(nil)

type Maximum float64

func (m Maximum) Apply(typeReq *TypeApplyRequest[*schema.JSON], fieldReq *FieldApplyRequest[*schema.JSON]) error {
	s := fieldReq.Schema
	if s.Type != schema.JSONTypeNumber {
		return errors.New("cannot apply `Maximum` to non numeric type")
	}
	val := int(m)
	s.Maximum = &val
	return nil
}

var _ Applier[*schema.JSON] = (*Minimum)(nil)

type Minimum float64

func (m Minimum) Apply(typeReq *TypeApplyRequest[*schema.JSON], fieldReq *FieldApplyRequest[*schema.JSON]) error {
	s := fieldReq.Schema
	if s.Type != schema.JSONTypeNumber {
		return errors.New("cannot apply `Minimum` to non numeric type")
	}
	val := int(m)
	s.Minimum = &val
	return nil
}

// string appliers
var _ Applier[*schema.JSON] = (*MaxLength)(nil)

type MaxLength int

func (m MaxLength) Apply(typeReq *TypeApplyRequest[*schema.JSON], fieldReq *FieldApplyRequest[*schema.JSON]) error {
	s := fieldReq.Schema
	if s.Type != schema.JSONTypeString {
		return errors.New("cannot apply `MaxLength` to non string type")
	}
	val := int(m)
	s.MaxLength = &val
	return nil
}

var _ Applier[*schema.JSON] = (*MinLength)(nil)

type MinLength int

func (m MinLength) Apply(typeReq *TypeApplyRequest[*schema.JSON], fieldReq *FieldApplyRequest[*schema.JSON]) error {
	s := fieldReq.Schema
	if s.Type != schema.JSONTypeString {
		return errors.New("cannot apply `MinLength` to non string type")
	}
	val := int(m)
	s.MinLength = &val
	return nil
}

var _ Applier[*schema.JSON] = (*Pattern)(nil)

type Pattern string

func (p Pattern) Apply(typeReq *TypeApplyRequest[*schema.JSON], fieldReq *FieldApplyRequest[*schema.JSON]) error {
	s := fieldReq.Schema
	if s.Type != schema.JSONTypeString {
		return errors.New("cannot apply `Pattern` to non string type")
	}
	val := string(p)
	if _, err := regexp.Compile(val); err != nil {
		return err
	}
	s.Pattern = val
	return nil
}

var _ Applier[*schema.JSON] = (*ContentEncoding)(nil)

type ContentEncoding string

func (c ContentEncoding) Apply(typeReq *TypeApplyRequest[*schema.JSON], fieldReq *FieldApplyRequest[*schema.JSON]) error {
	s := fieldReq.Schema
	if s.Type != schema.JSONTypeString {
		return errors.New("cannot apply `ContentEncoding` to non string type")
	}
	s.ContentEncoding = string(c)
	return nil
}

var _ Applier[*schema.JSON] = (*ContentMediatype)(nil)

type ContentMediatype string

func (c ContentMediatype) Apply(typeReq *TypeApplyRequest[*schema.JSON], fieldReq *FieldApplyRequest[*schema.JSON]) error {
	s := fieldReq.Schema
	if s.Type != schema.JSONTypeString {
		return errors.New("cannot apply `ContentMediatype` to non string type")
	}
	s.ContentMediatype = string(c)
	return nil
}

// array applier
var _ Applier[*schema.JSON] = (*MaxItems)(nil)

type MaxItems int

func (m MaxItems) Apply(typeReq *TypeApplyRequest[*schema.JSON], fieldReq *FieldApplyRequest[*schema.JSON]) error {
	s := fieldReq.Schema
	if s.Type != schema.JSONTypeArray {
		return errors.New("cannot apply `MaxItems` to non array type")
	}
	val := int(m)
	s.MaxItems = &val
	return nil
}

var _ Applier[*schema.JSON] = (*MinItems)(nil)

type MinItems int

func (m MinItems) Apply(typeReq *TypeApplyRequest[*schema.JSON], fieldReq *FieldApplyRequest[*schema.JSON]) error {
	s := fieldReq.Schema
	if s.Type != schema.JSONTypeArray {
		return errors.New("cannot apply `MinItems` to non array type")
	}
	val := int(m)
	s.MinItems = &val
	return nil
}

var _ Applier[*schema.JSON] = (*UniqueItems)(nil)

type UniqueItems bool

func (u UniqueItems) Apply(typeReq *TypeApplyRequest[*schema.JSON], fieldReq *FieldApplyRequest[*schema.JSON]) error {
	s := fieldReq.Schema
	if s.Type != schema.JSONTypeArray {
		return errors.New("cannot apply `UniqueItems` to non array type")
	}
	s.UniqueItems = bool(u)
	return nil

}

// object appliers
var _ Applier[*schema.JSON] = (*MaxProperties)(nil)

type MaxProperties int

func (m MaxProperties) Apply(typeReq *TypeApplyRequest[*schema.JSON], fieldReq *FieldApplyRequest[*schema.JSON]) error {
	s := typeReq.Schema
	if !s.IsObjectType() {
		return errors.New("cannot apply `MaxProperties` to non object type")
	}
	val := int(m)
	s.MaxProperties = &val
	return nil
}

var _ Applier[*schema.JSON] = (*MinProperties)(nil)

type MinProperties int

func (m MinProperties) Apply(typeReq *TypeApplyRequest[*schema.JSON], fieldReq *FieldApplyRequest[*schema.JSON]) error {
	s := typeReq.Schema
	if !s.IsObjectType() {
		return errors.New("cannot apply `MinProperties` to non object type")
	}
	val := int(m)
	s.MinProperties = &val
	return nil
}

var _ Applier[*schema.JSON] = (*Required)(nil)

type Required struct{}

func (r Required) Apply(typeReq *TypeApplyRequest[*schema.JSON], fieldReq *FieldApplyRequest[*schema.JSON]) error {
	name := schema.JSONNameForField(fieldReq.Info)
	typeReq.Schema.Required = append(typeReq.Schema.Required, name)
	return nil
}

// JSON Schema meta validations
var _ Applier[*schema.JSON] = (*ID)(nil)

type ID string

func (id ID) Apply(typeReq *TypeApplyRequest[*schema.JSON], fieldReq *FieldApplyRequest[*schema.JSON]) error {
	s := typeReq.Schema
	s.ID = string(id)
	return nil
}

var _ Applier[*schema.JSON] = (*Draft)(nil)

type Draft string

func (d Draft) Apply(typeReq *TypeApplyRequest[*schema.JSON], fieldReq *FieldApplyRequest[*schema.JSON]) error {
	s := typeReq.Schema
	s.Draft = string(d)
	return nil
}
