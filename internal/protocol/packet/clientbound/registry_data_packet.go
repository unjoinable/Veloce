package clientbound

import (
	"Veloce/internal/interfaces"
)

type RegistryDataPacket struct {
	data []byte
}

func (p *RegistryDataPacket) ID() int32 {
	return 0x07
}

func (p *RegistryDataPacket) Write(buf *interfaces.Buffer) {
	for i := range p.data {
		buf.WriteByte(p.data[i])
	}
}
