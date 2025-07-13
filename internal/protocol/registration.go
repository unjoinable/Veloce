package protocol

import (
	"Veloce/internal/interfaces"
	"Veloce/internal/network"
	"Veloce/internal/protocol/packet"
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
		return &packet.HandshakePacket{}
	})
}

func registerStatus(reg *network.PacketRegistry) {
	reg.RegisterServerBound(interfaces.Status, 0x00, func() interfaces.ServerboundPacket { return &packet.StatusRequestPacket{} })
	reg.RegisterServerBound(interfaces.Status, 0x01, func() interfaces.ServerboundPacket { return &packet.PingRequestPacket{} })
}

func registerLogin(reg *network.PacketRegistry) {
	reg.RegisterServerBound(interfaces.Login, 0x00, func() interfaces.ServerboundPacket { return &packet.LoginStartPacket{} })
	reg.RegisterServerBound(interfaces.Login, 0x03, func() interfaces.ServerboundPacket { return &packet.LoginAcknowledgedPacket{} })
}

func registerConfiguration(reg *network.PacketRegistry) {
	reg.RegisterServerBound(interfaces.Configuration, 0x00, func() interfaces.ServerboundPacket { return &packet.ClientInformationPacket{} })
	reg.RegisterServerBound(interfaces.Configuration, 0x02, func() interfaces.ServerboundPacket { return &packet.PluginMessagePacket{} })
	reg.RegisterServerBound(interfaces.Configuration, 0x03, func() interfaces.ServerboundPacket { return &packet.AcknowledgeFinishConfigurationPacket{} })
	reg.RegisterServerBound(interfaces.Configuration, 0x07, func() interfaces.ServerboundPacket { return &packet.ServerBoundKnownPacksPacket{} })
}

func registerPlay(reg *network.PacketRegistry) {
	reg.RegisterServerBound(interfaces.Play, 0x0B, func() interfaces.ServerboundPacket { return &packet.ClientTickEndPacket{} })
	reg.RegisterServerBound(interfaces.Play, 0x1C, func() interfaces.ServerboundPacket { return &packet.MovePlayerPosPacket{} })
	reg.RegisterServerBound(interfaces.Play, 0x1D, func() interfaces.ServerboundPacket { return &packet.MovePlayerPosRotPacket{} })
}
