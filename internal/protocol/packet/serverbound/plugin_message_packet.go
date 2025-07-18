package serverbound

import (
	common2 "Veloce/internal/network/common"
)

type PluginMessagePacket struct {
	identifier string
	data       []byte
}

func (p *PluginMessagePacket) ID() int32 {
	return 0x02
}

func (p *PluginMessagePacket) Read(buf *common2.Buffer) {
	p.identifier, _ = buf.ReadString()
	p.data = buf.Bytes()
}

func (p *PluginMessagePacket) Handle(*common2.PlayerConnection) {
	// Nothing to handle
}
