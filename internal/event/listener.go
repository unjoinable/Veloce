package event

// Listener is a function that handles an Event
type Listener func(Event)

// NewListener creates a new Listener from a simple function
func NewListener(fn func(Event)) Listener {
    return fn
}