package event

import (
	"fmt"
	"reflect"
	"sync"
)

// EventNode routes events to appropriate listeners
type EventNode struct {
	name     string
	mu       sync.RWMutex
	handles  map[reflect.Type]anyListenerHandle
	children []*EventNode
	parent   *EventNode
}

func NewNode(name string) *EventNode {
	return &EventNode{
		name:    name,
		handles: make(map[reflect.Type]anyListenerHandle),
	}
}

func (n *EventNode) AddChild(child *EventNode) {
	n.mu.Lock()
	defer n.mu.Unlock()
	child.parent = n
	n.children = append(n.children, child)
}

// Call dispatches a non-cancellable event to all listeners
func (n *EventNode) Call(event Event) {
	fmt.Println("[Call] Dispatching:", reflect.TypeOf(event))
	t := reflect.TypeOf(event)

	n.mu.RLock()
	defer n.mu.RUnlock()

	if h, ok := n.handles[t]; ok {
		fmt.Println("[Call] Found handle for", t)
		h.call(event)
	} else {
		fmt.Println("[Call] No handle found for", t)
	}

	for _, child := range n.children {
		child.Call(event)
	}
}

// CallCancellable dispatches a cancellable event and runs success() if not cancelled
func (n *EventNode) CallCancellable(event CancellableEvent, success func()) {
	n.Call(event)
	if !event.IsCancelled() {
		success()
	}
}

type anyListenerHandle interface {
	call(Event)
}

// ListenerHandle holds listeners for a specific event type
type ListenerHandle[T Event] struct {
	mu        sync.RWMutex
	listeners []EventListener[T]
}

func (h *ListenerHandle[T]) Add(listener EventListener[T]) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.listeners = append(h.listeners, listener)
}

func (h *ListenerHandle[T]) Remove(listener EventListener[T]) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for i, l := range h.listeners {
		if l == listener {
			h.listeners = append(h.listeners[:i], h.listeners[i+1:]...)
			break
		}
	}
}

func (h *ListenerHandle[T]) call(event Event) {
	fmt.Println("[Handle.call] Called with:", reflect.TypeOf(event))
	typed, ok := event.(T)
	if !ok {
		fmt.Println("[Handle.call] Type assertion FAILED")
		return
	}
	fmt.Println("[Handle.call] Type assertion SUCCESS")

	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, listener := range h.listeners {
		fmt.Println("[Handle.call] Calling listener")
		listener.Run(typed)
	}
}

// GetHandle -> Generic function for accessing typed handle
func GetHandle[T Event](n *EventNode) *ListenerHandle[T] {
	n.mu.Lock()
	defer n.mu.Unlock()

	t := reflect.TypeOf((*T)(nil)).Elem()
	fmt.Println("[GetHandle] Looking for", t)

	if h, ok := n.handles[t]; ok {
		fmt.Println("[GetHandle] Found existing handle for", t)
		return h.(*ListenerHandle[T])
	}

	
	handle := &ListenerHandle[T]{}
	n.handles[t] = handle
	return handle
}
