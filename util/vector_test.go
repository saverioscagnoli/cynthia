package util

import "testing"

func TestPushAndLength(t *testing.T) {
	v := NewVector[int]()

	v.Push(10)
	v.Push(20)

	if v.Length() != 2 {
		t.Fatalf("expected length 2, got %d", v.Length())
	}

	if *v.At(0) != 10 || *v.At(1) != 20 {
		t.Fatalf("unexpected values")
	}
}

func TestAt(t *testing.T) {
	v := NewVector[int]()
	v.Push(10)
	v.Push(20)
	v.Push(30)

	if *v.At(0) != 10 {
		t.Fatalf("expected 10")
	}
	if *v.At(1) != 20 {
		t.Fatalf("expected 20")
	}
	if *v.At(-1) != 30 {
		t.Fatalf("expected 30")
	}
	if *v.At(-2) != 20 {
		t.Fatalf("expected 20")
	}

	if v.At(100) != nil {
		t.Fatalf("expected nil for out of bounds")
	}
	if v.At(-100) != nil {
		t.Fatalf("expected nil for invalid negative index")
	}
}

func TestFirstLast(t *testing.T) {
	v := NewVector[int]()
	v.Push(1)
	v.Push(2)
	v.Push(3)

	if *v.First() != 1 {
		t.Fatalf("expected first = 1")
	}
	if *v.Last() != 3 {
		t.Fatalf("expected last = 3")
	}
}

func TestPop(t *testing.T) {
	v := NewVector[int]()
	v.Push(1)
	v.Push(2)

	val, ok := v.Pop()
	if !ok || val != 2 {
		t.Fatalf("expected 2")
	}
	if v.Length() != 1 {
		t.Fatalf("expected length 1")
	}

	val, ok = v.Pop()
	if !ok || val != 1 {
		t.Fatalf("expected 1")
	}

	_, ok = v.Pop()
	if ok {
		t.Fatalf("expected false on empty pop")
	}
}

func TestInsert(t *testing.T) {
	v := NewVector[int]()
	v.Push(1)
	v.Push(3)

	ok := v.Insert(1, 2)
	if !ok {
		t.Fatalf("insert failed")
	}

	if v.Length() != 3 {
		t.Fatalf("expected length 3")
	}

	if *v.At(0) != 1 || *v.At(1) != 2 || *v.At(2) != 3 {
		t.Fatalf("unexpected order")
	}
}

func TestInsertNegativeIndex(t *testing.T) {
	v := NewVector[int]()
	v.Push(1)
	v.Push(3)

	ok := v.Insert(-1, 2)
	if !ok {
		t.Fatalf("insert failed")
	}

	if *v.At(1) != 2 {
		t.Fatalf("expected 2 at index 1")
	}
}

func TestRemove(t *testing.T) {
	v := NewVector[int]()
	v.Push(1)
	v.Push(2)
	v.Push(3)

	val, ok := v.Remove(1)
	if !ok || val != 2 {
		t.Fatalf("expected removed 2")
	}

	if v.Length() != 2 {
		t.Fatalf("expected length 2")
	}

	if *v.At(0) != 1 || *v.At(1) != 3 {
		t.Fatalf("unexpected order")
	}
}

func TestRemoveNegativeIndex(t *testing.T) {
	v := NewVector[int]()
	v.Push(1)
	v.Push(2)
	v.Push(3)

	val, ok := v.Remove(-1)
	if !ok || val != 3 {
		t.Fatalf("expected removed 3")
	}

	if v.Length() != 2 {
		t.Fatalf("expected length 2")
	}

	if *v.At(0) != 1 || *v.At(1) != 2 {
		t.Fatalf("unexpected order")
	}
}

func TestOutOfBounds(t *testing.T) {
	v := NewVector[int]()
	v.Push(1)

	if v.Insert(5, 10) {
		t.Fatalf("expected insert to fail")
	}

	if v.Insert(-10, 10) {
		t.Fatalf("expected negative insert to fail")
	}

	_, ok := v.Remove(5)
	if ok {
		t.Fatalf("expected remove to fail")
	}

	_, ok = v.Remove(-10)
	if ok {
		t.Fatalf("expected negative remove to fail")
	}
}
