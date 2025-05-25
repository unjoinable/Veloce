package serverbound

import (
	"Veloce/internal/network"
	"Veloce/internal/network/protocol/clientbound"
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

func (h *LoginStartPacket) Read(buf *network.Buffer) {
	fmt.Println("Read LoginStartPacket")
	h.Username, _ = buf.ReadString()
	h.Uuid, _ = buf.ReadUUID()
}

func (h *LoginStartPacket) Handle(pc *network.PlayerConnection) {
	fmt.Println("Handle LoginStartPacket")
	packet := &clientbound.LoginSuccessPacket{}
	pc.SendPacket(packet)
}

type LoginAcknowledgedPacket struct{}

func (p *LoginAcknowledgedPacket) ID() int32 {
	return 0x03
}

func (p *LoginAcknowledgedPacket) Read(_ *network.Buffer) {
	fmt.Println("Read LoginAcknowledgedPacket")
}

func (p *LoginAcknowledgedPacket) Handle(pc *network.PlayerConnection) {
	fmt.Println("Handle LoginAcknowledgedPacket")
	pc.SetState(network.Configuration)
	packet := &clientbound.ClientBoundKnownPacksPacket{}
	pc.SendPacket(packet)
}
