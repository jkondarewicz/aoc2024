package utils

type Set[T comparable] struct {
	data map[T]struct{}
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
