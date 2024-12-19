package utils

import "fmt"

type Set[T comparable] struct {
	data map[T]struct{}
}

func (set *Set[T]) ForEach(block func (value T)) {
	for key := range set.data {
		block(key)
	}
}

func (set *Set[T]) Get() []T {
	keys := make([]T, 0, len(set.data))
	for k := range set.data {
		keys = append(keys, k)
	}
	return keys
}

func (set *Set[T]) String() string {
	return fmt.Sprintf("%v", set.data)
}

func (set *Set[T]) Copy() *Set[T] {
	new := NewSet[T]()
	for key := range set.data {
		new.Add(key)
	}
	return new
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{ data: make(map[T]struct{}) }
}

func (set *Set[T]) Exists(data T) bool {
	_, exists := set.data[data]
	return exists
}

func (set *Set[T]) Add(data T) *Set[T] {
	set.data[data] = struct{}{}
	return set
}

func (set *Set[T]) Remove(data T) *Set[T] {
	delete(set.data, data)
	return set
}

func (set *Set[T]) Size() int {
	return len(set.data)
}
