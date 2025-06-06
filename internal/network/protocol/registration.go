package protocol

import (
	"Veloce/internal/network"
	"Veloce/internal/network/protocol/serverbound/configuration"
	"Veloce/internal/network/protocol/serverbound/handshake"
	"Veloce/internal/network/protocol/serverbound/login"
	"Veloce/internal/network/protocol/serverbound/status"
)

func RegisterAllPackets() {
	// Handshaking state serverbound packets
	network.GlobalPacketRegistry.RegisterServerBound(network.Handshake, 0x00, func() network.ServerboundPacket {
		return &handshake.HandshakePacket{}
	})

	// Status state serverbound packets
	network.GlobalPacketRegistry.RegisterServerBound(network.Status, 0x00, func() network.ServerboundPacket {
		return &status.StatusRequestPacket{}
	})
	network.GlobalPacketRegistry.RegisterServerBound(network.Status, 0x01, func() network.ServerboundPacket {
		return &status.PingRequestPacket{}
	})

	// Login state serverbound packets
	network.GlobalPacketRegistry.RegisterServerBound(network.Login, 0x00, func() network.ServerboundPacket {
		return &login.LoginStartPacket{}
	})

	network.GlobalPacketRegistry.RegisterServerBound(network.Login, 0x03, func() network.ServerboundPacket {
		return &login.LoginAcknowledgedPacket{}
	})

	// Configuration State serverbound packets
	network.GlobalPacketRegistry.RegisterServerBound(network.Configuration, 0x07, func() network.ServerboundPacket {
		return &configuration.ServerBoundKnownPacksPacket{}
	})

	network.GlobalPacketRegistry.RegisterServerBound(network.Configuration, 0x02, func() network.ServerboundPacket {
		return &configuration.PluginMessagePacket{}
	})

	network.GlobalPacketRegistry.RegisterServerBound(network.Configuration, 0x00, func() network.ServerboundPacket {
		return &configuration.ClientInformationPacket{}
	})

	network.GlobalPacketRegistry.RegisterServerBound(network.Configuration, 0x03, func() network.ServerboundPacket {
		return &configuration.AcknowledgeFinishConfigurationPacket{}
	})
}
