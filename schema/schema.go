package schema

type Property[T any] interface {
    TypeOf() T
}

type Applier interface {
    ApplyToProperty() error
}
