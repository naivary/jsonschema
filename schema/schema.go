package schema

import "encoding/json"

type JSONType int

const (
	InvalidType JSONType = iota +1
	NullType    	
    BooleanType
	NumberType
	IntegerType
	StringType
	ArrayType
	ObjectType
)

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
	ID         string `json:"$id"`
    // TODO(naivary): add custom type for draft
	Schema     string `json:"$schema"`
	Title      string `json:"title"`
	Type       JSONType `json:"type"`
	Properties map[string]Property
}

// TODO(naivary): add the rest of the defined validation properties
type Property struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
