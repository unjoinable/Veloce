package serverbound

import (
	common2 "Veloce/internal/network/common"
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

func (h *LoginStartPacket) Read(buf *common2.Buffer) {
	h.Username, _ = buf.ReadString()
	h.Uuid, _ = buf.ReadUUID()
}

func (h *LoginStartPacket) Handle(pc *common2.PlayerConnection) {
	_ = pc.SendPacket(&clientbound.LoginSuccessPacket{})
}
