package protocol

import (
	"Veloce/internal/interfaces"
	"Veloce/internal/network"
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
	reg.RegisterServerBound(interfaces.Handshake, 0x00, func() interfaces.ServerboundPacket {
		return &serverbound.HandshakePacket{}
	})
}

func registerStatus(reg *network.PacketRegistry) {
	reg.RegisterServerBound(interfaces.Status, 0x00, func() interfaces.ServerboundPacket { return &serverbound.StatusRequestPacket{} })
	reg.RegisterServerBound(interfaces.Status, 0x01, func() interfaces.ServerboundPacket { return &serverbound.PingRequestPacket{} })
}

func registerLogin(reg *network.PacketRegistry) {
	reg.RegisterServerBound(interfaces.Login, 0x00, func() interfaces.ServerboundPacket { return &serverbound.LoginStartPacket{} })
	reg.RegisterServerBound(interfaces.Login, 0x03, func() interfaces.ServerboundPacket { return &serverbound.LoginAcknowledgedPacket{} })
}

func registerConfiguration(reg *network.PacketRegistry) {
	reg.RegisterServerBound(interfaces.Configuration, 0x00, func() interfaces.ServerboundPacket { return &serverbound.ClientInformationPacket{} })
	reg.RegisterServerBound(interfaces.Configuration, 0x02, func() interfaces.ServerboundPacket { return &serverbound.PluginMessagePacket{} })
	reg.RegisterServerBound(interfaces.Configuration, 0x03, func() interfaces.ServerboundPacket { return &serverbound.AcknowledgeFinishConfigurationPacket{} })
	reg.RegisterServerBound(interfaces.Configuration, 0x07, func() interfaces.ServerboundPacket { return &serverbound.ServerBoundKnownPacksPacket{} })
}

func registerPlay(reg *network.PacketRegistry) {
	reg.RegisterServerBound(interfaces.Play, 0x0B, func() interfaces.ServerboundPacket { return &serverbound.ClientTickEndPacket{} })
	reg.RegisterServerBound(interfaces.Play, 0x1C, func() interfaces.ServerboundPacket { return &serverbound.MovePlayerPosPacket{} })
	reg.RegisterServerBound(interfaces.Play, 0x1D, func() interfaces.ServerboundPacket { return &serverbound.MovePlayerPosRotPacket{} })
}
