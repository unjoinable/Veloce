package events

import (
	"Veloce/internal/interfaces"
)

// S -> C or Clientbound

type PacketOutgoingEvent struct {
	packet interfaces.ClientboundPacket
}

func (p *PacketOutgoingEvent) IsEvent() bool {
	return true
}

func (p *PacketOutgoingEvent) GetPacket() interfaces.ClientboundPacket {
	return p.packet
}

func NewPacketOutgoingEvent(p interfaces.ClientboundPacket) *PacketOutgoingEvent {
	return &PacketOutgoingEvent{packet: p}
}

// C -> S or Serverbound

type PacketIncomingEvent struct {
	packet interfaces.ServerboundPacket
}

func (p PacketIncomingEvent) IsEvent() bool {
	return true
}

func (p PacketIncomingEvent) GetPacket() interfaces.ServerboundPacket {
	return p.packet
}

func NewPacketIncomingEvent(p interfaces.ServerboundPacket) *PacketIncomingEvent {
	return &PacketIncomingEvent{packet: p}
}
