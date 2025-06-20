package play

import (
	"Veloce/internal/network"
	"Veloce/internal/network/buffer"
)

type MovePlayerPosPacket struct {
	X     int64
	Y     int64
	Z     int64
	Flags byte
}

func (m MovePlayerPosPacket) ID() int32 {
	return 0x1C
}

func (m MovePlayerPosPacket) Read(buf *buffer.Buffer) {
	m.X, _ = buf.ReadInt64()
	m.Y, _ = buf.ReadInt64()
	m.Z, _ = buf.ReadInt64()
	m.Flags, _ = buf.ReadByte()
}

func (m MovePlayerPosPacket) Handle(pc *network.PlayerConnection) {
	//TODO
}
