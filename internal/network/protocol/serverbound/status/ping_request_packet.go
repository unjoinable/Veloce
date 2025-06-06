package status

import (
	"Veloce/internal/network"
	"Veloce/internal/network/buffer"
	"Veloce/internal/network/protocol/clientbound/status"
	"fmt"
)

// PingRequestPacket represents a ping request from the client
type PingRequestPacket struct {
	Number int64
}

func (p *PingRequestPacket) ID() int32 {
	return 0x01
}

func (p *PingRequestPacket) Read(buf *buffer.Buffer) {
	fmt.Println("Reading PingRequestPacket")
	p.Number, _ = buf.ReadInt64()
}

func (p *PingRequestPacket) Handle(pc *network.PlayerConnection) {
	fmt.Println("Processing PingRequestPacket")
	packet := &status.PongPacket{Number: p.Number}
	pc.SendPacket(packet)
}
