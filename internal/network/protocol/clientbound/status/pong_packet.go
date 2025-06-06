package status

import (
	"Veloce/internal/network/buffer"
	"fmt"
)

// PongPacket represents a pong response to a ping request
type PongPacket struct {
	Number int64
}

func (p *PongPacket) ID() int32 {
	return 0x01
}

func (p *PongPacket) Write(buf *buffer.Buffer) {
	fmt.Println("Write PongPacket")
	buf.WriteInt64(p.Number)
}
