package login

import (
	"Veloce/internal/network"
	"Veloce/internal/network/buffer"
	"Veloce/internal/network/protocol/clientbound/login"
	"fmt"
	"github.com/google/uuid"
)

type LoginStartPacket struct {
	Username string
	Uuid     uuid.UUID
}

func (h *LoginStartPacket) ID() int32 {
	return 0x00
}

func (h *LoginStartPacket) Read(buf *buffer.Buffer) {
	h.Username, _ = buf.ReadString()
	h.Uuid, _ = buf.ReadUUID()
}

func (h *LoginStartPacket) Handle(pc *network.PlayerConnection) {
	err := pc.SendPacket(&login.LoginSuccessPacket{})

	if err != nil {
		fmt.Println("Error sending login success packet:", err)
	}
}
