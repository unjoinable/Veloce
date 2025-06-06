package configuration

import (
	"Veloce/internal/network"
	"Veloce/internal/network/buffer"
	"fmt"
)

type ServerBoundKnownPacksPacket struct {
	//TODO
}

func (p *ServerBoundKnownPacksPacket) ID() int32 {
	return 0x07
}

func (p *ServerBoundKnownPacksPacket) Read(buf *buffer.Buffer) {
	fmt.Println("Read ServerBoundKnownPacksPacket")
}

func (p *ServerBoundKnownPacksPacket) Handle(pc *network.PlayerConnection) {
	fmt.Println("ServerBoundKnownPacksPacket Handle")
}
