package packet

import (
	"Veloce/internal/interfaces"
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

type StatusRequestPacket struct { /*No Fields*/
}

func (p *StatusRequestPacket) ID() int32 {
	return 0x00
}

func (p *StatusRequestPacket) Read(*interfaces.Buffer) { /*Nothing to read*/ }
