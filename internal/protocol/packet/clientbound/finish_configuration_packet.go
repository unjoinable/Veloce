package clientbound

import (
	"Veloce/internal/network/common"
)

type FinishConfigurationPacket struct { /*No Fields*/
}

func (p FinishConfigurationPacket) ID() int32 {
	return 0x03
}

func (p FinishConfigurationPacket) Write(*common.Buffer) {
	//Nothing to write
}
