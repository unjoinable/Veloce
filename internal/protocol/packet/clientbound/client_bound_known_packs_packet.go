package clientbound

import (
	"Veloce/internal/network/common"
)

type ClientBoundKnownPacksPacket struct {
	//TODO
}

func (p *ClientBoundKnownPacksPacket) ID() int32 {
	return 0x0E
}

func (p *ClientBoundKnownPacksPacket) Write(buf *common.Buffer) {
	buf.WriteVarInt(1)
	buf.WriteString("minecraft")
	buf.WriteString("core")
	buf.WriteString("1.21.5")
}
