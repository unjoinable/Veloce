package events

import (
	"Veloce/internal/network"
)

// S -> C or Clientbound

type PacketOutgoingEvent struct {
	packet network.ClientboundPacket
}

func (p *PacketOutgoingEvent) IsEvent() bool {
	return true
}

func (p *PacketOutgoingEvent) GetPacket() network.ClientboundPacket {
	return p.packet
}

func NewPacketOutgoingEvent(p network.ClientboundPacket) *PacketOutgoingEvent {
	return &PacketOutgoingEvent{packet: p}
}

// C -> S or Serverbound

type PacketIncomingEvent struct {
	packet network.ServerboundPacket
}

func (p PacketIncomingEvent) IsEvent() bool {
	return true
}

func (p PacketIncomingEvent) GetPacket() network.ServerboundPacket {
	return p.packet
}

func NewPacketIncomingEvent(p network.ServerboundPacket) *PacketIncomingEvent {
	return &PacketIncomingEvent{packet: p}
}
