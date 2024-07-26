package schema

type TypeConverter[T any] interface {
	Invalid() T

	Int() T
	UInt() T
	Float() T
	String() T
	Struct() T
	Map() T
	Array() T
	Slice() T
	Bool() T
	Complex() T
}
