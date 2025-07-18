package serverbound

import (
	common2 "Veloce/internal/network/common"
)

type ClientTickEndPacket struct {
	// No Fields
}

func (c ClientTickEndPacket) ID() int32 {
	return 0x0B
}

func (c ClientTickEndPacket) Read(*common2.Buffer) {
	// No Fields
}

func (c ClientTickEndPacket) Handle(*common2.PlayerConnection) {
	// No Handling
}
