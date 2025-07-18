package clientbound

import (
	"Veloce/internal/interfaces"
)

type FinishConfigurationPacket struct { /*No Fields*/
}

func (p FinishConfigurationPacket) ID() int32 {
	return 0x03
}

func (p FinishConfigurationPacket) Write(*interfaces.Buffer) {
	//Nothing to write
}
