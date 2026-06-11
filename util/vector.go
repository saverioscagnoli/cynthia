package util

import "encoding/json"

type Vector[T any] struct {
	items []T
}

// Creates an empty vector
func NewVector[T any]() *Vector[T] {
	return &Vector[T]{
		items: make([]T, 0),
	}
}

func (v *Vector[T]) Arr() *[]T {
	return &v.items
}

// Returns the length of the vector
func (v *Vector[T]) Length() int {
	return len(v.items)
}

// Returns an element, if present, at a specific index in the vector.
//
// If the index is negative, it will start counting from the back of the vector, just like python.
//
// So, if you input `vector.At(-1)`, it will return the last item in the vector, if present.
func (v *Vector[T]) At(index int) *T {
	if index < 0 {
		index = v.Length() + index
	}

	if index < 0 || index >= v.Length() {
		return nil
	}

	return &v.items[index]
}

// Returns the first element in the vector, if present.
func (v *Vector[T]) First() *T {
	return v.At(0)
}

// Returns the last element in the vector, if present.
func (v *Vector[T]) Last() *T {
	return v.At(-1)
}

// Pushes an element at the end of the vector.
func (v *Vector[T]) Push(item T) {
	v.items = append(v.items, item)

}

// Removes the last element from the vector.
//
// If it could not be removed (i.e. empty vector), the flag returned with the element will be set to `false`, and `true` otherwise.
func (v *Vector[T]) Pop() (T, bool) {
	n := v.Length()

	if n == 0 {
		var zero T
		return zero, false
	}

	n--
	item := v.items[n]

	clear(v.items[n : n+1])

	v.items = v.items[:n]

	return item, true
}

// Inserts an element at a specific index.
//
// If the index is negative, it will start counting from the back of the vector, just like python.
func (v *Vector[T]) Insert(index int, value T) bool {
	n := len(v.items)

	if index < 0 {
		index = n + index
	}

	if index < 0 || index > n {
		return false
	}

	v.items = append(v.items[:index],
		append([]T{value}, v.items[index:]...)...)

	return true
}

// Removes an element at a specific index.Remove
//
// If the index is negative, it will start counting from the back of the vector, just like python.
func (v *Vector[T]) Remove(index int) (T, bool) {
	var zero T
	n := len(v.items)

	if index < 0 {
		index = n + index
	}

	if index < 0 || index >= n {
		return zero, false
	}

	item := v.items[index]

	copy(v.items[index:], v.items[index+1:])

	// clear last element (for GC safety)
	v.items[n-1] = zero

	v.items = v.items[:n-1]

	return item, true
}

// Calls the defined function over a for loop over all the vector
func (v *Vector[T]) ForEach(fn func(int, T)) {
	for i, val := range v.items {
		fn(i, val)
	}
}

func (v *Vector[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.items)
}

func (v *Vector[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &v.items)
}
