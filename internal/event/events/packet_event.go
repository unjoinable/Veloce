package events

import (
	"Veloce/internal/network/common"
)

// S -> C or Clientbound

type PacketOutgoingEvent struct {
	packet common.ClientboundPacket
}

func (p *PacketOutgoingEvent) IsEvent() bool {
	return true
}

func (p *PacketOutgoingEvent) GetPacket() common.ClientboundPacket {
	return p.packet
}

func NewPacketOutgoingEvent(p common.ClientboundPacket) *PacketOutgoingEvent {
	return &PacketOutgoingEvent{packet: p}
}

// C -> S or Serverbound

type PacketIncomingEvent struct {
	packet common.ServerboundPacket
}

func (p PacketIncomingEvent) IsEvent() bool {
	return true
}

func (p PacketIncomingEvent) GetPacket() common.ServerboundPacket {
	return p.packet
}

func NewPacketIncomingEvent(p common.ServerboundPacket) *PacketIncomingEvent {
	return &PacketIncomingEvent{packet: p}
}
