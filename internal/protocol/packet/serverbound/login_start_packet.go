package serverbound

import (
	"Veloce/internal/interfaces"
	"Veloce/internal/protocol/packet/clientbound"
	"github.com/google/uuid"
)

type LoginStartPacket struct {
	Username string
	Uuid     uuid.UUID
}

func (h *LoginStartPacket) ID() int32 {
	return 0x00
}

func (h *LoginStartPacket) Read(buf *interfaces.Buffer) {
	h.Username, _ = buf.ReadString()
	h.Uuid, _ = buf.ReadUUID()
}

func (h *LoginStartPacket) Handle(pc *interfaces.PlayerConnection) {
	_ = pc.SendPacket(&clientbound.LoginSuccessPacket{})
}
