package serverbound

import (
	common2 "Veloce/internal/network/common"
)

type ConfirmTeleportationPacket struct {
	TeleportId int32
}

func (c ConfirmTeleportationPacket) ID() int32 {
	return 0x00
}

func (c ConfirmTeleportationPacket) Read(buf *common2.Buffer) {
	c.TeleportId, _ = buf.ReadVarInt()
}

func (c ConfirmTeleportationPacket) Handle(*common2.PlayerConnection) {
	// No Handling
}
