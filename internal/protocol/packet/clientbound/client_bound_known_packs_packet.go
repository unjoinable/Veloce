package clientbound

import (
	"Veloce/internal/interfaces"
)

type ClientBoundKnownPacksPacket struct {
	//TODO
}

func (p *ClientBoundKnownPacksPacket) ID() int32 {
	return 0x0E
}

func (p *ClientBoundKnownPacksPacket) Write(buf *interfaces.Buffer) {
	buf.WriteVarInt(1)
	buf.WriteString("minecraft")
	buf.WriteString("core")
	buf.WriteString("1.21.5")
}
