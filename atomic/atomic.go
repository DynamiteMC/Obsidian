package atomic

import "sync/atomic"

func New[T any](val T) Value[T] {
	v := Value[T]{}
	v.Set(val)

	return v
}

type Value[T any] struct {
	val atomic.Value
}

func (v *Value[T]) Get() T {
	var val T
	if v == nil {
		return val
	}
	if v, ok := v.val.Load().(T); ok {
		val = v
	}
	return val
}

func (v *Value[T]) Set(val T) {
	if v == nil {
		v = new(Value[T])
	}
	v.val.Store(val)
}
