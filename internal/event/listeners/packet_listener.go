package listener

import (
	"Veloce/internal/event"
	"Veloce/internal/event/events"
	"fmt"
)

// RegisterPacketListeners registers both incoming and outgoing packet listeners.
func RegisterPacketListeners() {
	event.GetGlobalHandle[*events.PacketOutgoingEvent]().Add(event.NewEventListener(func(e *events.PacketOutgoingEvent) {
		fmt.Printf("Packet outgoing: %+v\n", e.GetPacket())
	}))

	event.GetGlobalHandle[*events.PacketIncomingEvent]().Add(event.NewEventListener(func(e *events.PacketIncomingEvent) {
		fmt.Printf("Packet incoming: %+v\n", e.GetPacket())
	}))
}
