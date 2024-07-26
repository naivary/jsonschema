package typesutil

import (
	"go/types"

	"github.com/naivary/specraft/schema"
)

func IsType[T any](t types.Type) (T, bool) {
	switch v := t.(type) {
	case T:
		return v, true
	case *types.Pointer:
		return IsType[T](v.Elem())
	case *types.Named:
		return IsType[T](v.Underlying())
	default:
		return *new(T), false
	}
}

func Convert[T any](t types.Type, tc schema.TypeConverter[T]) T {
	switch v := t.(type) {
	case *types.Basic:
		return basicKindConversion(v.Kind(), tc)
	case *types.Pointer:
		return Convert(v.Elem(), tc)
	case *types.Slice:
		return tc.Slice()
	case *types.Array:
		return tc.Array()
	case *types.Struct:
		return tc.Struct()
	case *types.Named:
		return Convert(v.Underlying(), tc)
	case *types.Map:
		return tc.Map()
	default:
		return tc.Invalid()
	}

}

func basicKindConversion[T any](kind types.BasicKind, tc schema.TypeConverter[T]) T {
	switch kind {
	case types.Bool:
		return tc.Bool()
	case types.Int, types.Int8, types.Int16, types.Int32, types.Int64:
		return tc.Int()
	case types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64, types.Uintptr:
		return tc.UInt()
	case types.Float32, types.Float64:
		return tc.Float()
	case types.Complex64, types.Complex128:
		return tc.Complex()
	case types.String:
		return tc.String()
	default:
		return tc.Invalid()
	}
}
