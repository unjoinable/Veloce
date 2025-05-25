package serverbound

import (
	"Veloce/internal/network"
	"Veloce/internal/network/protocol/clientbound"
	"fmt"
)

// StatusRequestPacket represents a status request from the client
type StatusRequestPacket struct{}

func (p *StatusRequestPacket) ID() int32 {
	return 0x00
}

func (p *StatusRequestPacket) Read(*network.Buffer) {
	fmt.Println("Reading StatusRequestPacket")
}

func (p *StatusRequestPacket) Handle(pc *network.PlayerConnection) {
	fmt.Println("Processing StatusRequestPacket")
	packet := &clientbound.StatusResponsePacket{}
	pc.SendPacket(packet)
}

// PingRequestPacket represents a ping request from the client
type PingRequestPacket struct {
	Number int64
}

func (p *PingRequestPacket) ID() int32 {
	return 0x01
}

func (p *PingRequestPacket) Read(buf *network.Buffer) {
	fmt.Println("Reading PingRequestPacket")
	p.Number, _ = buf.ReadInt64()
}

func (p *PingRequestPacket) Handle(pc *network.PlayerConnection) {
	fmt.Println("Processing PingRequestPacket")
	packet := &clientbound.PongPacket{Number: p.Number}
	pc.SendPacket(packet)
}
