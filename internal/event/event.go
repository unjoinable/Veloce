package event

import "Veloce/internal/entity/player"

// Event represents a generic event type.
type Event interface {
    IsEvent()
}

// CancellableEvent is an event that can be cancelled.
type CancellableEvent interface {
    Event
    IsCancelled() bool
    SetCancelled(isCancelled bool)
}

// PlayerEvent is an event associated with a specific player.
type PlayerEvent interface {
    Event
    GetPlayer() player.Player
}