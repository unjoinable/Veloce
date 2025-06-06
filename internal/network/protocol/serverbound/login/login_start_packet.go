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
	fmt.Println("Read LoginStartPacket")
	h.Username, _ = buf.ReadString()
	h.Uuid, _ = buf.ReadUUID()
}

func (h *LoginStartPacket) Handle(pc *network.PlayerConnection) {
	fmt.Println("Handle LoginStartPacket")
	fmt.Println(h.Username)
	fmt.Println(h.Uuid)
	packet := &login.LoginSuccessPacket{}
	pc.SendPacket(packet)
}
