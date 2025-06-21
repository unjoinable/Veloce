package protocol

import (
	"Veloce/internal/network"
	"Veloce/internal/network/protocol/serverbound/configuration"
	"Veloce/internal/network/protocol/serverbound/handshake"
	"Veloce/internal/network/protocol/serverbound/login"
	"Veloce/internal/network/protocol/serverbound/play"
	"Veloce/internal/network/protocol/serverbound/status"
)

func RegisterAllPackets() {
	reg := network.GlobalPacketRegistry

	registerHandshake(reg)
	registerStatus(reg)
	registerLogin(reg)
	registerConfiguration(reg)
	registerPlay(reg)
}

func registerHandshake(reg *network.PacketRegistry) {
	reg.RegisterServerBound(network.Handshake, 0x00, func() network.ServerboundPacket {
		return &handshake.HandshakePacket{}
	})
}

func registerStatus(reg *network.PacketRegistry) {
	reg.RegisterServerBound(network.Status, 0x00, func() network.ServerboundPacket { return &status.StatusRequestPacket{} })
	reg.RegisterServerBound(network.Status, 0x01, func() network.ServerboundPacket { return &status.PingRequestPacket{} })
}

func registerLogin(reg *network.PacketRegistry) {
	reg.RegisterServerBound(network.Login, 0x00, func() network.ServerboundPacket { return &login.LoginStartPacket{} })
	reg.RegisterServerBound(network.Login, 0x03, func() network.ServerboundPacket { return &login.LoginAcknowledgedPacket{} })
}

func registerConfiguration(reg *network.PacketRegistry) {
	reg.RegisterServerBound(network.Configuration, 0x00, func() network.ServerboundPacket { return &configuration.ClientInformationPacket{} })
	reg.RegisterServerBound(network.Configuration, 0x02, func() network.ServerboundPacket { return &configuration.PluginMessagePacket{} })
	reg.RegisterServerBound(network.Configuration, 0x03, func() network.ServerboundPacket { return &configuration.AcknowledgeFinishConfigurationPacket{} })
	reg.RegisterServerBound(network.Configuration, 0x07, func() network.ServerboundPacket { return &configuration.ServerBoundKnownPacksPacket{} })
}

func registerPlay(reg *network.PacketRegistry) {
	reg.RegisterServerBound(network.Play, 0x0B, func() network.ServerboundPacket { return &play.ClientTickEndPacket{} })
	reg.RegisterServerBound(network.Play, 0x1C, func() network.ServerboundPacket { return &play.MovePlayerPosPacket{} })
	reg.RegisterServerBound(network.Play, 0x1D, func() network.ServerboundPacket { return &play.MovePlayerPosRotPacket{} })
}
