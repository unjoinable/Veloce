package serverbound

import (
	common2 "Veloce/internal/network/common"
	"Veloce/internal/protocol/packet/clientbound"
)

type StatusRequestPacket struct {
	/*No Fields*/
}

func (p *StatusRequestPacket) ID() int32 {
	return 0x00
}

func (p *StatusRequestPacket) Read(*common2.Buffer) {
	/*Nothing to read*/
}

func (p *StatusRequestPacket) Handle(pc *common2.PlayerConnection) {
	_ = pc.SendPacket(&clientbound.StatusResponsePacket{})
}
