package schema

import (
	"strings"

	"sigs.k8s.io/controller-tools/pkg/markers"
)

const (
	JSONTypeInvalid JSONType = iota + 1
	JSONTypeNull
	JSONTypeBoolean
	JSONTypeNumber
	JSONTypeInteger
	JSONTypeString
	JSONTypeArray
	JSONTypeObject
)

type JSONType int

func (t JSONType) String() string {
	switch t {
	case JSONTypeNull:
		return "null"
	case JSONTypeBoolean:
		return "boolean"
	case JSONTypeNumber:
		return "number"
	case JSONTypeInteger:
		return "integer"
	case JSONTypeString:
		return "string"
	case JSONTypeArray:
		return "array"
	case JSONTypeObject:
		return "object"
	}
	return "INVALID_TYPE"
}

func (j JSONType) MarshalText() ([]byte, error) {
	return []byte(j.String()), nil
}

func NewJSON() *JSON {
	return &JSON{}
}

type JSON struct {
	// type agnostic fields
	ID     string   `json:"$id,omitempty"`
	Draft  string   `json:"$schema,omitempty"`
	Type   JSONType `json:"type"`
	Format string   `json:"format,omitempty"`

	// annotations
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Default     *any   `json:"default,omitempty"`
	Comment     string `json:"comment,omitempty"`
	ReadOnly    bool   `json:"readOnly,omitempty"`
	WriteOnly   bool   `json:"writeOnly,omitempty"`
	Examples    []any  `json:"examples,omitempty"`
	Deprecated  bool   `json:"deprecated,omitempty"`

	// object
	MaxProperties     *int                `json:"maxProperties,omitempty"`
	MinProperties     *int                `json:"minProperties,omitempty"`
	Required          []string            `json:"required,omitempty"`
	Properties        map[string]*JSON    `json:"properties,omitempty"`
	DependentRequired map[string][]string `json:"dependentRequired,omitempty"`

	// string
	MaxLength        *int   `json:"maxLength,omitempty"`
	MinLength        *int   `json:"minLength,omitempty"`
	Pattern          string `json:"pattern,omitempty"`
	ContentEncoding  string `json:"contentEncoding,omitempty"`
	ContentMediatype string `json:"contentMediatype,omitempty"`

	// numertic
	Maximum          *int     `json:"maximum,omitempty"`
	Minimum          *int     `json:"minimum,omitempty"`
	MultipleOf       *float64 `json:"multipleOf,omitempty"`
	ExclusiveMaximum bool     `json:"exclusiveMaximum,omitempty"`
	ExclusiveMinimum bool     `json:"exclusiveMinimum,omitempty"`

	// array
	Items       *JSON `json:"items,omitempty"`
	MaxItems    *int  `json:"maxItems,omitempty"`
	MinItems    *int  `json:"minItems,omitempty"`
	UniqueItems bool  `json:"uniqueItems,omitempty"`
}

func (j *JSON) IsType(t JSONType) bool {
	return j.Type == t
}

func (j *JSON) IsInvalidType() bool {
	return j.Type == JSONTypeInvalid
}

func (j *JSON) IsObjectType() bool {
	return j.Type == JSONTypeObject
}

func JSONNameForField(info *markers.FieldInfo) string {
	jsonName := info.Tag.Get("json")
	if jsonName != "" {
		return strings.Split(jsonName, ",")[0]
	}
	return strings.ToLower(info.Name)
}

func NewJSONSchemaTypeConvert() TypeConverter[JSONType] {
	return jsonSchemaTypeConverter{}
}

var _ TypeConverter[JSONType] = (*jsonSchemaTypeConverter)(nil)

type jsonSchemaTypeConverter struct{}

func (jsonSchemaTypeConverter) Invalid() JSONType {
	return JSONTypeInvalid
}

func (jsonSchemaTypeConverter) Int() JSONType {
	return JSONTypeNumber
}
func (jsonSchemaTypeConverter) UInt() JSONType {
	return JSONTypeNumber
}
func (jsonSchemaTypeConverter) Float() JSONType {
	return JSONTypeNumber
}
func (jsonSchemaTypeConverter) String() JSONType {
	return JSONTypeString
}
func (jsonSchemaTypeConverter) Struct() JSONType {
	return JSONTypeObject
}
func (jsonSchemaTypeConverter) Map() JSONType {
	return JSONTypeObject
}
func (jsonSchemaTypeConverter) Array() JSONType {
	return JSONTypeObject
}
func (jsonSchemaTypeConverter) Slice() JSONType {
	return JSONTypeArray
}
func (jsonSchemaTypeConverter) Bool() JSONType {
	return JSONTypeArray
}
func (jsonSchemaTypeConverter) Complex() JSONType {
	return JSONTypeNumber
}
