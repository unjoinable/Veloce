package configuration

import (
	"Veloce/internal/network"
	"Veloce/internal/network/buffer"
	"fmt"
)

type PluginMessagePacket struct {
	identifier string
	data       []byte
}

func (p *PluginMessagePacket) ID() int32 {
	return 0x02
}

func (p *PluginMessagePacket) Read(buf *buffer.Buffer) {
	fmt.Println("Read PluginMessagePacket")
	p.identifier, _ = buf.ReadString()
	p.data = buf.Data()
	fmt.Println(p.identifier)
}

func (p *PluginMessagePacket) Handle(pc *network.PlayerConnection) {
	fmt.Println("PluginMessagePacket Handle")
}
