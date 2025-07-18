package serverbound

import (
	common2 "Veloce/internal/network/common"
	"Veloce/internal/protocol/packet/clientbound"
)

type LoginAcknowledgedPacket struct {
	// No Fields
}

func (p *LoginAcknowledgedPacket) ID() int32 {
	return 0x03
}

func (p *LoginAcknowledgedPacket) Read(_ *common2.Buffer) {
	// No Reading
}

func (p *LoginAcknowledgedPacket) Handle(pc *common2.PlayerConnection) {
	pc.SetState(common2.Configuration)
	_ = pc.SendPacket(&clientbound.ClientBoundKnownPacksPacket{})
}
