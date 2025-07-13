package packet

import (
	"Veloce/internal/interfaces"
	"Veloce/internal/objects/protocol"
)

type HandshakePacket struct {
	ProtocolVersion protocol.VarInt
	ServerAddress   string
	ServerPort      uint16
	NextState       protocol.VarInt
}

func (h *HandshakePacket) ID() int32 {
	return 0x00
}

func (h *HandshakePacket) Read(buf *interfaces.Buffer) {
	h.ProtocolVersion, _ = buf.ReadVarInt()
	h.ServerAddress, _ = buf.ReadString()
	h.ServerPort, _ = buf.ReadUint16()
	h.NextState, _ = buf.ReadVarInt()
}
