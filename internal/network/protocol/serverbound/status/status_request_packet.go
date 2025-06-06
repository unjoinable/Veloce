package status

import (
	"Veloce/internal/network"
	"Veloce/internal/network/buffer"
	"Veloce/internal/network/protocol/clientbound/status"
	"fmt"
)

// StatusRequestPacket represents a status request from the client
type StatusRequestPacket struct{}

func (p *StatusRequestPacket) ID() int32 {
	return 0x00
}

func (p *StatusRequestPacket) Read(*buffer.Buffer) {
	fmt.Println("Reading StatusRequestPacket")
}

func (p *StatusRequestPacket) Handle(pc *network.PlayerConnection) {
	fmt.Println("Processing StatusRequestPacket")
	packet := &status.StatusResponsePacket{}
	pc.SendPacket(packet)
}
