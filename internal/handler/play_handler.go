package handler

import (
	"Veloce/internal/interfaces"
)

type ClientTickEndPacketHandler struct{}

func (c ClientTickEndPacketHandler) HandlePacket(pc interfaces.Connection, p interfaces.ServerboundPacket) {
	//No Handling Required
}

type ConfirmTeleportationPacketHandler struct{}

func (c ConfirmTeleportationPacketHandler) HandlePacket(pc interfaces.Connection, p interfaces.ServerboundPacket) {
}

type MovePlayerPosPacketHandler struct{}

func (m MovePlayerPosPacketHandler) HandlePacket(pc interfaces.Connection, p interfaces.ServerboundPacket) {
}

type MovePlayerPosRotPacketHandler struct{}

func (m MovePlayerPosRotPacketHandler) HandlePacket(pc interfaces.Connection, p interfaces.ServerboundPacket) {
}
