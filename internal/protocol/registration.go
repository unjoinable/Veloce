package protocol

import (
	"Veloce/internal/network"
	"Veloce/internal/network/common"
	"Veloce/internal/protocol/packet/serverbound"
)

func RegisterAllPackets(reg *network.PacketRegistry) {
	registerHandshake(reg)
	registerStatus(reg)
	registerLogin(reg)
	registerConfiguration(reg)
	registerPlay(reg)
}

func registerHandshake(reg *network.PacketRegistry) {
	reg.RegisterServerBound(common.Handshake, 0x00, func() common.ServerboundPacket {
		return &serverbound.HandshakePacket{}
	})
}

func registerStatus(reg *network.PacketRegistry) {
	reg.RegisterServerBound(common.Status, 0x00, func() common.ServerboundPacket { return &serverbound.StatusRequestPacket{} })
	reg.RegisterServerBound(common.Status, 0x01, func() common.ServerboundPacket { return &serverbound.PingRequestPacket{} })
}

func registerLogin(reg *network.PacketRegistry) {
	reg.RegisterServerBound(common.Login, 0x00, func() common.ServerboundPacket { return &serverbound.LoginStartPacket{} })
	reg.RegisterServerBound(common.Login, 0x03, func() common.ServerboundPacket { return &serverbound.LoginAcknowledgedPacket{} })
}

func registerConfiguration(reg *network.PacketRegistry) {
	reg.RegisterServerBound(common.Configuration, 0x00, func() common.ServerboundPacket { return &serverbound.ClientInformationPacket{} })
	reg.RegisterServerBound(common.Configuration, 0x02, func() common.ServerboundPacket { return &serverbound.PluginMessagePacket{} })
	reg.RegisterServerBound(common.Configuration, 0x03, func() common.ServerboundPacket { return &serverbound.AcknowledgeFinishConfigurationPacket{} })
	reg.RegisterServerBound(common.Configuration, 0x07, func() common.ServerboundPacket { return &serverbound.ServerBoundKnownPacksPacket{} })
}

func registerPlay(reg *network.PacketRegistry) {
	reg.RegisterServerBound(common.Play, 0x0B, func() common.ServerboundPacket { return &serverbound.ClientTickEndPacket{} })
	reg.RegisterServerBound(common.Play, 0x1C, func() common.ServerboundPacket { return &serverbound.MovePlayerPosPacket{} })
	reg.RegisterServerBound(common.Play, 0x1D, func() common.ServerboundPacket { return &serverbound.MovePlayerPosRotPacket{} })
}
