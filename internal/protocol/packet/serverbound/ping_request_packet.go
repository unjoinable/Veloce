package serverbound

import (
	common2 "Veloce/internal/network/common"
	"Veloce/internal/protocol/packet/clientbound"
)

type PingRequestPacket struct {
	Number int64
}

func (p *PingRequestPacket) ID() int32 {
	return 0x01
}

func (p *PingRequestPacket) Read(buf *common2.Buffer) {
	p.Number, _ = buf.ReadInt64()
}

func (p *PingRequestPacket) Handle(pc *common2.PlayerConnection) {
	_ = pc.SendPacket(&clientbound.PongPacket{Number: p.Number})
}
