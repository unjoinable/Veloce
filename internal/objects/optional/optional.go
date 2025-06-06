package optional

// Optional represents a container that may or may not contain a value
type Optional[T any] struct {
	value *T
}

// Of creates an Optional with a value
func Of[T any](value T) *Optional[T] {
	return &Optional[T]{value: &value}
}

// OfNilable creates an Optional from a potentially nil pointer
func OfNilable[T any](ptr *T) *Optional[T] {
	if ptr == nil {
		return Empty[T]()
	}
	return Of(*ptr)
}

// Empty returns an empty Optional
func Empty[T any]() *Optional[T] {
	return &Optional[T]{value: nil}
}

// IsPresent returns true if a value is present
func (o *Optional[T]) IsPresent() bool {
	return o.value != nil
}

// IsEmpty returns true if no value is present
func (o *Optional[T]) IsEmpty() bool {
	return o.value == nil
}

// Get returns the value if present, panics otherwise
func (o *Optional[T]) Get() T {
	if o.value == nil {
		panic("optional: no value present")
	}
	return *o.value
}

// OrElse returns the value if present, otherwise returns other
func (o *Optional[T]) OrElse(other T) T {
	if o.value != nil {
		return *o.value
	}
	return other
}

// OrElseGet returns the value if present, otherwise calls supplier
func (o *Optional[T]) OrElseGet(supplier func() T) T {
	if o.value != nil {
		return *o.value
	}
	return supplier()
}

// OrElsePanic returns the value if present, otherwise panics with msg
func (o *Optional[T]) OrElsePanic(msg string) T {
	if o.value != nil {
		return *o.value
	}
	panic(msg)
}

// IfPresent calls consumer with the value if present
func (o *Optional[T]) IfPresent(consumer func(T)) {
	if o.value != nil {
		consumer(*o.value)
	}
}

// IfPresentOrElse calls consumer if present, otherwise calls emptyAction
func (o *Optional[T]) IfPresentOrElse(consumer func(T), emptyAction func()) {
	if o.value != nil {
		consumer(*o.value)
	} else {
		emptyAction()
	}
}
