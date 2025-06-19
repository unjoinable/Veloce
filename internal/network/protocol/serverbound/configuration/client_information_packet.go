package configuration

import (
	"Veloce/internal/network"
	"Veloce/internal/network/buffer"
	"fmt"
)

type ClientInformationPacket struct {
	locale        string
	render        byte
	chatMode      int32
	chatColor     bool
	skin          byte
	mainHand      int32
	filter        bool
	serverListing bool
	particle      int32
}

func (p *ClientInformationPacket) ID() int32 {
	return 0x00
}

func (p *ClientInformationPacket) Read(buf *buffer.Buffer) {
	fmt.Println("Read ClientInformationPacket")
	p.locale, _ = buf.ReadString()
	p.render, _ = buf.ReadByte()
	p.chatMode, _ = buf.ReadVarInt()
	p.chatColor, _ = buf.ReadBool()
	p.skin, _ = buf.ReadByte()
	p.mainHand, _ = buf.ReadVarInt()
	p.filter, _ = buf.ReadBool()
	p.serverListing, _ = buf.ReadBool()
	p.particle, _ = buf.ReadVarInt()
}

func (p *ClientInformationPacket) Handle(*network.PlayerConnection) {
	fmt.Println("ClientInformationPacket Handle")
}
