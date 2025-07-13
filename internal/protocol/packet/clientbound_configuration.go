package packet

import (
	"Veloce/internal/interfaces"
)

type ClientBoundKnownPacksPacket struct {
	//TODO
}

func (p *ClientBoundKnownPacksPacket) ID() int32 {
	return 0x0E
}

func (p *ClientBoundKnownPacksPacket) Write(buf *interfaces.Buffer) {
	
	buf.WriteVarInt(1)
	buf.WriteString("minecraft")
	buf.WriteString("core")
	buf.WriteString("1.21.5")
}

type FinishConfigurationPacket struct { /*No Fields*/
}

func (p FinishConfigurationPacket) ID() int32 {
	return 0x03
}

func (p FinishConfigurationPacket) Write(*interfaces.Buffer) {
	//Nothing to write
}

type RegistryDataPacket struct {
	data []byte
}

func (p *RegistryDataPacket) ID() int32 {
	return 0x07
}

func (p *RegistryDataPacket) Write(buf *interfaces.Buffer) {
	buf.WriteBytes(p.data)
}
