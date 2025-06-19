package set

// ImmutableSet a Set impl who's data cannot be changed after init
type ImmutableSet[T comparable] struct {
	data map[T]struct{}
}

func NewImmutableSet[T comparable](values ...T) *ImmutableSet[T] {
	m := make(map[T]struct{}, len(values))
	for _, v := range values {
		m[v] = struct{}{}
	}
	return &ImmutableSet[T]{data: m}
}

func (s *ImmutableSet[T]) Add(T) {
	panic("ImmutableSet: cannot Add to an immutable set")
}

func (s *ImmutableSet[T]) Remove(T) {
	panic("ImmutableSet: cannot Remove from an immutable set")
}

func (s *ImmutableSet[T]) Clear() {
	panic("ImmutableSet: cannot Clear an immutable set")
}

func (s *ImmutableSet[T]) Contains(value T) bool {
	_, ok := s.data[value]
	return ok
}

func (s *ImmutableSet[T]) Len() int {
	return len(s.data)
}

func (s *ImmutableSet[T]) Values() []T {
	values := make([]T, 0, len(s.data))
	for v := range s.data {
		values = append(values, v)
	}
	return values
}
