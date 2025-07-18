package serverbound

import (
	"Veloce/internal/interfaces"
	"Veloce/internal/protocol/packet/clientbound"
)

type StatusRequestPacket struct {
	/*No Fields*/
}

func (p *StatusRequestPacket) ID() int32 {
	return 0x00
}

func (p *StatusRequestPacket) Read(*interfaces.Buffer) {
	/*Nothing to read*/
}

func (p *StatusRequestPacket) Handle(pc *interfaces.PlayerConnection) {
	_ = pc.SendPacket(&clientbound.StatusResponsePacket{})
}
