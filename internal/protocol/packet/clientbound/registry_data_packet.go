package clientbound

import (
	"Veloce/internal/network/common"
)

type RegistryDataPacket struct {
	data []byte
}

func (p *RegistryDataPacket) ID() int32 {
	return 0x07
}

func (p *RegistryDataPacket) Write(buf *common.Buffer) {
	for i := range p.data {
		buf.WriteByte(p.data[i])
	}
}
