package serverbound

import (
	"Veloce/internal/interfaces"
)

type ClientTickEndPacket struct {
	// No Fields
}

func (c ClientTickEndPacket) ID() int32 {
	return 0x0B
}

func (c ClientTickEndPacket) Read(*interfaces.Buffer) {
	// No Fields
}

func (c ClientTickEndPacket) Handle(*interfaces.PlayerConnection) {
	// No Handling
}
