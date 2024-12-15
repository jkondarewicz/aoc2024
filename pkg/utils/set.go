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

func (set *Set[T]) String() string {
	return fmt.Sprintf("%v", set.data)
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
