package serverbound

import (
	"Veloce/internal/interfaces"
)

type MovePlayerPosRotPacket struct {
	X     int64
	Y     int64
	Z     int64
	Yaw   float32
	Pitch float32
	Flags byte
}

func (m MovePlayerPosRotPacket) ID() int32 {
	return 0x1D
}

func (m MovePlayerPosRotPacket) Read(buf *interfaces.Buffer) {
	m.X, _ = buf.ReadInt64()
	m.Y, _ = buf.ReadInt64()
	m.Z, _ = buf.ReadInt64()
	m.Yaw, _ = buf.ReadFloat32()
	m.Pitch, _ = buf.ReadFloat32()
	m.Flags, _ = buf.ReadByte()
}

func (m MovePlayerPosRotPacket) Handle(pc *interfaces.PlayerConnection) {
	// No handling
}
