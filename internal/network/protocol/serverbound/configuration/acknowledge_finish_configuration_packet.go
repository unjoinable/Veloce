package configuration

import (
	"Veloce/internal/network"
	"Veloce/internal/network/buffer"
	"fmt"
)

type AcknowledgeFinishConfigurationPacket struct { /*No Fields*/
}

func (p *AcknowledgeFinishConfigurationPacket) ID() int32 {
	return 0x03
}

func (p *AcknowledgeFinishConfigurationPacket) Read(*buffer.Buffer) {
	fmt.Println("Write AcknowledgeFinishConfigurationPacket")
}

func (p *AcknowledgeFinishConfigurationPacket) Handle(pc *network.PlayerConnection) {
	fmt.Println("Handle AcknowledgeFinishConfigurationPacket")
}
