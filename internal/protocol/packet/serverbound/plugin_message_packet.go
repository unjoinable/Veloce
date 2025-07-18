package serverbound

import (
	"Veloce/internal/interfaces"
)

type PluginMessagePacket struct {
	identifier string
	data       []byte
}

func (p *PluginMessagePacket) ID() int32 {
	return 0x02
}

func (p *PluginMessagePacket) Read(buf *interfaces.Buffer) {
	p.identifier, _ = buf.ReadString()
	p.data = buf.Bytes()
}

func (p *PluginMessagePacket) Handle(*interfaces.PlayerConnection) {
	// Nothing to handle
}
