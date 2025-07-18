package serverbound

import (
	"Veloce/internal/interfaces"
	"Veloce/internal/protocol/packet/clientbound"
)

type PingRequestPacket struct {
	Number int64
}

func (p *PingRequestPacket) ID() int32 {
	return 0x01
}

func (p *PingRequestPacket) Read(buf *interfaces.Buffer) {
	p.Number, _ = buf.ReadInt64()
}

func (p *PingRequestPacket) Handle(pc *interfaces.PlayerConnection) {
	_ = pc.SendPacket(&clientbound.PongPacket{Number: p.Number})
}
