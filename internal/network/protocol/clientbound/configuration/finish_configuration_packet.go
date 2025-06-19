package configuration

import "Veloce/internal/network/buffer"

type FinishConfigurationPacket struct { /*No Fields*/
}

func (p FinishConfigurationPacket) ID() int32 {
	return 0x03
}

func (p FinishConfigurationPacket) Write(*buffer.Buffer) {
	//Nothing to write
}
