package atomicvalue

import "testing"

func TestAtomicValue(t *testing.T) {
	var av AtomicValue[int]
	av.Store(10)
	if av.Load() != 10 {
		t.Error("Atomic value not set correctly")
	}
}

func TestNewAtomicValue(t *testing.T) {
	av := NewAtomicValue(10)
	if av.Load() != 10 {
		t.Error("Atomic value not set correctly")
	}
}

func TestAtomicValueSwap(t *testing.T) {
	av := NewAtomicValue(10)
	prev := av.Swap(20)
	if prev != 10 {
		t.Error("previous value should be 10")
	}
	if av.Load() != 20 {
		t.Error("Atomic value not set correctly")
	}
}

func TestAtomicValueSwapNullInitial(t *testing.T) {
	var av AtomicValue[int]
	prev := av.Swap(20)
	if prev != 0 {
		t.Error("previous value should be 0")
	}
	if av.Load() != 20 {
		t.Error("Atomic value not set correctly")
	}
}

func TestAtomicValueCompareAndSwap(t *testing.T) {
	av := NewAtomicValue(10)
	if !av.CompareAndSwap(10, 20) {
		t.Error("Atomic value not set correctly")
	}
	if av.CompareAndSwap(30, 30) {
		t.Error("Atomic value not set correctly")
	}
}

func TestAtomicValueLoadNil(t *testing.T) {
	av := AtomicValue[error]{}
	if err := av.Load(); err != nil {
		t.Error("Atomic value not set correctly")
	}
}
