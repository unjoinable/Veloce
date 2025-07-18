package serverbound

import (
	"Veloce/internal/interfaces"
)

type ConfirmTeleportationPacket struct {
	TeleportId int32
}

func (c ConfirmTeleportationPacket) ID() int32 {
	return 0x00
}

func (c ConfirmTeleportationPacket) Read(buf *interfaces.Buffer) {
	c.TeleportId, _ = buf.ReadVarInt()
}

func (c ConfirmTeleportationPacket) Handle(*interfaces.PlayerConnection) {
	// No Handling
}
