package handler

import (
	"Veloce/internal/interfaces"
	"Veloce/internal/protocol/packet"
)

type HandshakePacketHandler struct{}

func (h *HandshakePacketHandler) HandlePacket(conn interfaces.Connection, p interfaces.ServerboundPacket) {
	handshakePacket := p.(*packet.HandshakePacket)
	conn.SetState(interfaces.ConnectionState(handshakePacket.NextState))
}
