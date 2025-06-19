package event

import (
	"Veloce/internal/entity/player"
)

// Event is a marker interface for all events
type Event interface {
	IsEvent() bool
}

// CancellableEvent represents an event that can be cancelled
type CancellableEvent interface {
	Event
	IsCancelled() bool
	SetCancelled(cancel bool)
}

// PlayerEvent links an event to a specific player
type PlayerEvent interface {
	Event
	GetPlayer() *player.Player
}

// Global Event Handler
var globalEventHandler = NewNode("global")

func GetGlobalEventHandler() *EventNode {
	return globalEventHandler
}

func GetGlobalHandle[T Event]() *ListenerHandle[T] {
	return GetHandle[T](globalEventHandler)
}
