package event

import "reflect"

// Node represents a node in the event tree, capable of registering listeners,
// dispatching events, and propagating them up the tree.
type Node struct {
    parent    *Node
    children  []*Node
    listeners map[reflect.Type][]Listener
}

// NewNode creates and returns a new root node with no parent.
func NewNode() *Node {
    return &Node{
        listeners: make(map[reflect.Type][]Listener),
    }
}

// AddChild attaches an existing child node to this node,
// re-parenting it if it already has a different parent.
func (n *Node) AddChild(child *Node) {
    if child == nil {
        return
    }

    // Detach from any previous parent
    if child.parent != nil {
        child.parent.removeChild(child)
    }

    child.parent = n
    n.children = append(n.children, child)
}

// removeChild detaches the specified child from this node's children list.
func (n *Node) removeChild(child *Node) {
    for i, c := range n.children {
        if c == child {
            n.children = append(n.children[:i], n.children[i+1:]...)
            return
        }
    }
}

// Register adds a listener function for the given event type.
func (n *Node) Register(event Event, listener Listener) {
    eventType := reflect.TypeOf(event)
    n.listeners[eventType] = append(n.listeners[eventType], listener)
}

// CallEvent sends an event to all listeners on this node and then bubbles it up to the parent.
func (n *Node) CallEvent(event Event) {
    eventType := reflect.TypeOf(event)

    if listeners, ok := n.listeners[eventType]; ok {
        for _, listener := range listeners {
            listener(event)
        }
    }

    if n.parent != nil {
        n.parent.CallEvent(event)
    }
}

// CallCancelledEvent dispatches a cancellable event and calls successCallback if it's not cancelled.
func (n *Node) CallCancelledEvent(event CancellableEvent, successCallback func()) {
    n.CallEvent(event)
    if !event.IsCancelled() {
        successCallback()
    }
}