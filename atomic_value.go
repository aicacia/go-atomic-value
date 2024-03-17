package types

import "sync/atomic"

type AtomicValue[T any] struct {
	atomic.Value
}

func NewAtomicValue[T any](value T) AtomicValue[T] {
	var av atomic.Value
	av.Store(value)
	return AtomicValue[T]{av}
}

func (av *AtomicValue[T]) Load() T {
	return av.Value.Load().(T)
}

func (av *AtomicValue[T]) Store(value T) {
	av.Value.Store(value)
}

func (av *AtomicValue[T]) Swap(value T) T {
	oldValue := av.Value.Swap(value)
	if oldValue == nil {
		var zero T
		return zero
	} else {
		return oldValue.(T)
	}
}

func (av *AtomicValue[T]) CompareAndSwap(old T, new T) bool {
	return av.Value.CompareAndSwap(old, new)
}
