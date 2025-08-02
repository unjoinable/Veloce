package events

import "Veloce/internal/network/common"

// PacketIncomingEvent is for C -> S (C2S)
type PacketIncomingEvent struct {
	Packet     *common.ServerboundPacket
	Connection *common.PlayerConnection
}

func (evt *PacketIncomingEvent) IsEvent() {}

// PacketOutgoingEvent is for S -> C (S2C)
type PacketOutgoingEvent struct {
	Packet     *common.ClientboundPacket
	Connection *common.PlayerConnection
}

func (evt *PacketOutgoingEvent) IsEvent() {}