package definer

import "sigs.k8s.io/controller-tools/pkg/markers"

type FieldApplyRequest[T any] struct {
	Info   *markers.FieldInfo
	Schema T
}

type TypeApplyRequest[T any] struct {
	Info   *markers.TypeInfo
	Schema T
}

type Applier[T any] interface {
	Apply(typeReq *TypeApplyRequest[T], fieldReq *FieldApplyRequest[T]) error
}

func NewFieldApplyRequest[T any](info *markers.FieldInfo, s T) *FieldApplyRequest[T] {
	return &FieldApplyRequest[T]{
		Info:   info,
		Schema: s,
	}
}

func NewTypeApplyRequest[T any](info *markers.TypeInfo, s T) *TypeApplyRequest[T] {
	return &TypeApplyRequest[T]{
		Info:   info,
		Schema: s,
	}
}
