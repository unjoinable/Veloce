package play

import (
	"Veloce/internal/network"
	"Veloce/internal/network/buffer"
	"Veloce/internal/objects/protocol"
)

type ConfirmTeleportationPacket struct {
	TeleportId protocol.VarInt
}

func (c ConfirmTeleportationPacket) ID() int32 {
	return 0x00
}

func (c ConfirmTeleportationPacket) Read(buf *buffer.Buffer) {
	c.TeleportId, _ = buf.ReadVarInt()
}

func (c ConfirmTeleportationPacket) Handle(pc *network.PlayerConnection) {
	//TODO
}
