package handler

import (
	"Veloce/internal/interfaces"
	"fmt"
)

type PacketHandler interface {
	HandlePacket(conn interfaces.Connection, packet interfaces.ServerboundPacket)
}

type PacketDispatcher struct {
	handlers      map[interfaces.ConnectionState]map[int32]PacketHandler
	packetFactory interfaces.PacketFactory
}

func NewPacketDispatcher(packetFactory interfaces.PacketFactory) *PacketDispatcher {
	return &PacketDispatcher{
		handlers:      make(map[interfaces.ConnectionState]map[int32]PacketHandler),
		packetFactory: packetFactory,
	}
}

func (pd *PacketDispatcher) RegisterHandler(state interfaces.ConnectionState, packetID int32, handler PacketHandler) {
	if pd.handlers[state] == nil {
		pd.handlers[state] = make(map[int32]PacketHandler)
	}
	pd.handlers[state][packetID] = handler
}

func (pd *PacketDispatcher) Dispatch(conn interfaces.Connection, state interfaces.ConnectionState, packetID int32, buf *interfaces.Buffer) error {
	packet, ok := pd.packetFactory.CreateServerBound(state, packetID)
	if !ok {
		return fmt.Errorf("unknown packet ID %d for state %v", packetID, state)
	}

	if packet == nil {
		return fmt.Errorf("packet not implemented: ID %d, State %v", packetID, state)
	}

	packet.Read(buf)

	handler, ok := pd.handlers[state][packetID]
	if !ok {
		return fmt.Errorf("no handler registered for packet ID %d in state %v", packetID, state)
	}

	handler.HandlePacket(conn, packet)
	return nil
}
