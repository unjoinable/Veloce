package handler

import (
	"Veloce/internal/interfaces"
	"Veloce/internal/protocol/packet"
)

type StatusRequestPacketHandler struct{}

func (h *StatusRequestPacketHandler) HandlePacket(pc interfaces.Connection, p interfaces.ServerboundPacket) {
	_ = pc.SendPacket(&packet.StatusResponsePacket{})
}

type PingRequestPacketHandler struct{}

func (h *PingRequestPacketHandler) HandlePacket(pc interfaces.Connection, p interfaces.ServerboundPacket) {
	pingPacket := p.(*packet.PingRequestPacket)
	_ = pc.SendPacket(&packet.PongPacket{Number: pingPacket.Number})
}
