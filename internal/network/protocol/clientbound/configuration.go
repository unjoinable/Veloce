package clientbound

import (
	"Veloce/internal/network"
	"fmt"
)

type ClientBoundKnownPacksPacket struct {
	//TODO
}

func (p *ClientBoundKnownPacksPacket) ID() int32 {
	return 0x0E
}

func (p *ClientBoundKnownPacksPacket) Write(buf *network.Buffer) {
	fmt.Println("Write ClientBoundKnownPacksPacket")
	buf.WriteByte(0)
}

type FinishConfigurationPacket struct { /*No Fields*/
}

func (p *FinishConfigurationPacket) ID() int32 {
	return 0x03
}

func (p *FinishConfigurationPacket) Write(*network.Buffer) {
	fmt.Println("Write FinishConfigurationPacket")
}
