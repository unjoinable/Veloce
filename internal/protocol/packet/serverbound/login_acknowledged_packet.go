package serverbound

import (
	"Veloce/internal/interfaces"
	"Veloce/internal/protocol/packet/clientbound"
)

type LoginAcknowledgedPacket struct {
	// No Fields
}

func (p *LoginAcknowledgedPacket) ID() int32 {
	return 0x03
}

func (p *LoginAcknowledgedPacket) Read(_ *interfaces.Buffer) {
	// No Reading
}

func (p *LoginAcknowledgedPacket) Handle(pc *interfaces.PlayerConnection) {
	pc.SetState(interfaces.Configuration)
	_ = pc.SendPacket(&clientbound.ClientBoundKnownPacksPacket{})
}
