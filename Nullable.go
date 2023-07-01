package nullable

type N[T any] struct {
	value    T
	hasValue bool
}

func (n *N[T]) GetValue() T {
	if n.hasValue {
		return n.value
	}
	panic("not valid value")
}
func (n *N[T]) GetValueOr(def T) T {
	if n.hasValue {
		return n.value
	}
	return def
}

func (n *N[T]) GetValueOrZero() T {
	if n.hasValue {
		return n.value
	}
	var def T
	return def
}

func (n *N[T]) SetValue(value T) {
	n.hasValue = true
	n.value = value
}

func (n *N[T]) HasValue() bool {
	return n.hasValue

}

func NewNullable[T any]() *N[T] {
	return &N[T]{}
}

type Int = *N[int]
type Float32 = *N[float32]
type Int64 = *N[int64]
