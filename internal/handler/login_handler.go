package handler

import (
	"Veloce/internal/interfaces"
	"Veloce/internal/protocol/packet"
)

type LoginStartPacketHandler struct{}

func (h *LoginStartPacketHandler) HandlePacket(pc interfaces.Connection, p interfaces.ServerboundPacket) {
	_ = pc.SendPacket(&packet.LoginSuccessPacket{})
}

type LoginAcknowledgedPacketHandler struct{}

func (h *LoginAcknowledgedPacketHandler) HandlePacket(pc interfaces.Connection, p interfaces.ServerboundPacket) {
	pc.SetState(interfaces.Configuration)

	_ = pc.SendPacket(&packet.ClientBoundKnownPacksPacket{})
}
