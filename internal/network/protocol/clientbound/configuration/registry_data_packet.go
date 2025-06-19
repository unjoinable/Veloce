package configuration

import "Veloce/internal/network/buffer"

type RegistryDataPacket struct {
	data []byte
}

func (p *RegistryDataPacket) ID() int32 {
	return 0x07
}

func (p *RegistryDataPacket) Write(buf *buffer.Buffer) {
	buf.WriteBytes(p.data)

}
