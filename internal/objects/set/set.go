package set

type Set[T comparable] interface {
	Add(value T)
	Remove(value T)
	Contains(value T) bool
	Len() int
	Values() []T
	Clear()
}
