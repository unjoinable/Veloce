package status

import (
	"Veloce/internal/network"
	"Veloce/internal/network/buffer"
	"Veloce/internal/network/protocol/clientbound/status"
	"fmt"
)

// StatusRequestPacket represents a status request from the client
type StatusRequestPacket struct { /*No Fields*/
}

func (p *StatusRequestPacket) ID() int32 {
	return 0x00
}

func (p *StatusRequestPacket) Read(*buffer.Buffer) { /*Nothing to read*/ }

func (p *StatusRequestPacket) Handle(pc *network.PlayerConnection) {
	err := pc.SendPacket(&status.StatusResponsePacket{})

	if err != nil {
		fmt.Println("Error sending status packet:", err)
	}
}
