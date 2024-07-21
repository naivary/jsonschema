package schema

type Property interface {
    TypeOf() any
}

type Applier interface {
    ApplyToProperty() error
}
