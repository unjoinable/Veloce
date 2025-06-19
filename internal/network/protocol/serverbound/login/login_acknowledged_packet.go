package login

import (
	"Veloce/internal/network"
	"Veloce/internal/network/buffer"
	"Veloce/internal/network/protocol/clientbound/configuration"
	"fmt"
)

type LoginAcknowledgedPacket struct { /*No Fields*/
}

func (p *LoginAcknowledgedPacket) ID() int32 {
	return 0x03
}

func (p *LoginAcknowledgedPacket) Read(_ *buffer.Buffer) { /*Nothing to read*/ }

func (p *LoginAcknowledgedPacket) Handle(pc *network.PlayerConnection) {
	pc.SetState(network.Configuration)

	err := pc.SendPacket(&configuration.ClientBoundKnownPacksPacket{})
	if err != nil {
		fmt.Println("Error sending configuration packet:", err)
	}
}
