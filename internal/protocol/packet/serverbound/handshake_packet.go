package serverbound

import (
	"Veloce/internal/network/common"
)

type HandshakePacket struct {
	ProtocolVersion int32
	ServerAddress   string
	ServerPort      uint16
	NextState       int32
}

func (h *HandshakePacket) ID() int32 {
	return 0x00
}

func (h *HandshakePacket) Read(buf *common.Buffer) {
	h.ProtocolVersion, _ = buf.ReadVarInt()
	h.ServerAddress, _ = buf.ReadString()
	h.ServerPort, _ = buf.ReadUint16()
	h.NextState, _ = buf.ReadVarInt()
}

func (h *HandshakePacket) Handle(pc *common.PlayerConnection) {
	pc.SetState(common.ConnectionState(h.NextState))
}
