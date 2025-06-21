package play

import (
	"Veloce/internal/network"
	"Veloce/internal/network/buffer"
)

type ClientTickEndPacket struct { /*No Fields*/
}

func (c ClientTickEndPacket) ID() int32 {
	return 0x0B
}

func (c ClientTickEndPacket) Read(*buffer.Buffer) {
	//No Fields
}

func (c ClientTickEndPacket) Handle(*network.PlayerConnection) {
	//No Handling Required
}
