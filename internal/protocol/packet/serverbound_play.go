package packet

import (
	"Veloce/internal/interfaces"
	"Veloce/internal/objects/protocol"
)

type ClientTickEndPacket struct { /*No Fields*/
}

func (c ClientTickEndPacket) ID() int32 {
	return 0x0B
}

func (c ClientTickEndPacket) Read(*interfaces.Buffer) {
	//No Fields
}

type ConfirmTeleportationPacket struct {
	TeleportId protocol.VarInt
}

func (c ConfirmTeleportationPacket) ID() int32 {
	return 0x00
}

func (c ConfirmTeleportationPacket) Read(buf *interfaces.Buffer) {
	c.TeleportId, _ = buf.ReadVarInt()
}

type MovePlayerPosPacket struct {
	X     int64
	Y     int64
	Z     int64
	Flags byte
}

func (m MovePlayerPosPacket) ID() int32 {
	return 0x1C
}

func (m MovePlayerPosPacket) Read(buf *interfaces.Buffer) {
	m.X, _ = buf.ReadInt64()
	m.Y, _ = buf.ReadInt64()
	m.Z, _ = buf.ReadInt64()
	m.Flags, _ = buf.ReadByte()
}

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
