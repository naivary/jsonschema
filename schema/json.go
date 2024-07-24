package schema

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
	return ""
}

func (j JSONType) MarshalText() ([]byte, error) {
	return []byte(j.String()), nil
}

type JSON struct {
	ID string `json:"$id,omitempty"`
	// TODO(naivary): add custom type for draft
	Draft             string              `json:"$schema,omitempty"`
	Title             string              `json:"title,omitempty"`
	Description       string              `json:"description,omitempty"`
	Type              JSONType            `json:"type"`
	DependentRequired map[string][]string `json:"dependentRequired,omitempty"`
	Properties        map[string]*JSON    `json:"properties,omitempty"`

	// string properties
	MaxLength *int   `json:"maxLength,omitempty"`
	MinLength *int   `json:"minLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
	Format    string `json:"format,omitempty"`

	// numertic properties
	Maximum          *int `json:"maximum,omitempty"`
	Minimum          *int `json:"minimum,omitempty"`
	MultipleOf       *int `json:"multipleOf,omitempty"`
	ExclusiveMaximum bool `json:"exclusiveMaximum,omitempty"`
	ExclusiveMinimum bool `json:"exclusiveMinimum,omitempty"`

	// array properties
	Items       *JSON `json:"items,omitempty"`
	MaxItems    *int  `json:"maxItems,omitempty"`
	MinItems    *int  `json:"minItems,omitempty"`
	UniqueItems bool  `json:"uniqueItems,omitempty"`
}
