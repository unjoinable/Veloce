package event

// EventListener defines an object that handles a specific event type
type EventListener[T Event] interface {
	Run(event T)
}

// NewEventListener adapts a func(T) to EventListener[T]
func NewEventListener[T Event](fn func(T)) EventListener[T] {
	return listenerFunc[T](fn)
}

type listenerFunc[T Event] func(T)

func (f listenerFunc[T]) Run(event T) {
	f(event)
}
