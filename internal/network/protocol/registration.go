package protocol

import (
	"Veloce/internal/network"
	"Veloce/internal/network/protocol/serverbound"
)

func RegisterAllPackets() {
	// Handshaking state serverbound packets
	network.GlobalPacketRegistry.RegisterServerBound(network.Handshake, 0x00, func() network.ServerboundPacket {
		return &serverbound.HandshakePacket{}
	})

	// Status state serverbound packets
	network.GlobalPacketRegistry.RegisterServerBound(network.Status, 0x00, func() network.ServerboundPacket {
		return &serverbound.StatusRequestPacket{}
	})
	network.GlobalPacketRegistry.RegisterServerBound(network.Status, 0x01, func() network.ServerboundPacket {
		return &serverbound.PingRequestPacket{}
	})

	// Login state serverbound packets
	network.GlobalPacketRegistry.RegisterServerBound(network.Login, 0x00, func() network.ServerboundPacket {
		return &serverbound.LoginStartPacket{}
	})

	network.GlobalPacketRegistry.RegisterServerBound(network.Login, 0x03, func() network.ServerboundPacket {
		return &serverbound.LoginAcknowledgedPacket{}
	})

	// Configuration State serverbound packets
	network.GlobalPacketRegistry.RegisterServerBound(network.Configuration, 0x03, func() network.ServerboundPacket {
		return &serverbound.AcknowledgeFinishConfigurationPacket{}
	})
}
