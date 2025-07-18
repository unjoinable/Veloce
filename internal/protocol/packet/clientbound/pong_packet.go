package clientbound

import (
	"Veloce/internal/network/common"
)

// PongPacket represents a pong response to a ping request
type PongPacket struct {
	Number int64
}

func (p *PongPacket) ID() int32 {
	return 0x01
}

func (p *PongPacket) Write(buf *common.Buffer) {
	buf.WriteInt64(p.Number)
}
