package json

import "encoding/json"

const (
	InvalidType JSONType = iota + 1
	NullType
	BooleanType
	NumberType
	IntegerType
	StringType
	ArrayType
	ObjectType
)

type JSONType int

func (j JSONType) String() string {
	switch j {
	case NullType:
		return "null"
	case BooleanType:
		return "boolean"
	case NumberType:
		return "number"
	case IntegerType:
		return "integer"
	case StringType:
		return "string"
	case ArrayType:
		return "array"
	case ObjectType:
		return "object"
	}
	return ""
}

func (j JSONType) MarshalText() ([]byte, error) {
	return []byte(j.String()), nil
}

func TypeOf(v any) JSONType {
	switch v.(type) {
	case nil:
		return NullType
	case bool:
		return BooleanType
	case json.Number, float32, float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return NumberType
	case string:
		return StringType
	case []any:
		return ArrayType
	case map[string]any:
		return ObjectType
	default:
		return InvalidType
	}
}

type JSONSchema struct {
	// TODO(naivary): id,schema must be type markers
	ID string `json:"$id"`
	// TODO(naivary): add custom type for draft
	Schema            string                  `json:"$schema"`
	Title             string                  `json:"title,omitempty"`
	Description       string                  `json:"description,omitempty"`
	Type              JSONType                `json:"type"`
	DependentRequired map[string][]string     `json:"dependentRequired,omitempty"`
	Properties        map[string]JSONProperty `json:"properties"`
}

// TODO(naivary): some of these fields are valid with some types.
// It must be validated
type JSONProperty struct {
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
	Enum        []any  `json:"enum,omitempty"`

	// Possible Property Types
	StringProperty
	NumericProperty
	ArrayProperty
}

type StringProperty struct {
	MaxLength int    `json:"maxLength,omitempty"`
	MinLength int    `json:"minLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
	Format    string `json:"format,omitempty"`
}

type NumericProperty struct {
	Maximum          int  `json:"maximum,omitempty"`
	Minimum          int  `json:"minimum,omitempty"`
	MultipleOf       int  `json:"multipleOf,omitempty"`
	ExclusiveMaximum bool `json:"exclusiveMaximum,omitempty"`
	ExclusiveMinimum bool `json:"exclusiveMinimum,omitempty"`
}

type ArrayProperty struct {
	Items       *ArrayItem `json:"items,omitempty"`
	MaxItems    int        `json:"maxItems,omitempty"`
	MinItems    int        `json:"minItems,omitempty"`
	UniqueItems bool       `json:"uniqueItems,omitempty"`
}

type ArrayItem struct {
	NumericProperty
	StringProperty
}
