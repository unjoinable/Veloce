package set

type HashSet[T comparable] struct {
	data map[T]struct{}
}

func NewHashSet[T comparable](values ...T) *HashSet[T] {
	s := &HashSet[T]{data: make(map[T]struct{})}
	for _, v := range values {
		s.Add(v)
	}
	return s
}

func (s *HashSet[T]) Add(value T) {
	s.data[value] = struct{}{}
}

func (s *HashSet[T]) Remove(value T) {
	delete(s.data, value)
}

func (s *HashSet[T]) Contains(value T) bool {
	_, ok := s.data[value]
	return ok
}

func (s *HashSet[T]) Len() int {
	return len(s.data)
}

func (s *HashSet[T]) Values() []T {
	values := make([]T, 0, len(s.data))
	for k := range s.data {
		values = append(values, k)
	}
	return values
}

func (s *HashSet[T]) Clear() {
	s.data = make(map[T]struct{})
}
