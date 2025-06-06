package configuration

import (
	"Veloce/internal/network/buffer"
	"fmt"
)

type ClientBoundKnownPacksPacket struct {
	//TODO
}

func (p *ClientBoundKnownPacksPacket) ID() int32 {
	return 0x0E
}

func (p *ClientBoundKnownPacksPacket) Write(buf *buffer.Buffer) {
	fmt.Println("Write ClientBoundKnownPacksPacket")
	buf.WriteVarInt(1)
	buf.WriteString("minecraft")
	buf.WriteString("core")
	buf.WriteString("1.21.5")
}
