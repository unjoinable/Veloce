package login

import (
	"Veloce/internal/network"
	"Veloce/internal/network/buffer"
	"Veloce/internal/network/protocol/clientbound/configuration"
	"fmt"
	"time"
)

type LoginAcknowledgedPacket struct{}

func (p *LoginAcknowledgedPacket) ID() int32 {
	return 0x03
}

func (p *LoginAcknowledgedPacket) Read(_ *buffer.Buffer) {
	fmt.Println("Read LoginAcknowledgedPacket")
}

func (p *LoginAcknowledgedPacket) Handle(pc *network.PlayerConnection) {
	fmt.Println("Handle LoginAcknowledgedPacket")
	pc.SetState(network.Configuration)

	time.Sleep(50 * time.Millisecond) // wait 1 tick
	packet := &configuration.ClientBoundKnownPacksPacket{}
	pc.SendPacket(packet)
}
