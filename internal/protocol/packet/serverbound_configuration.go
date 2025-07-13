package packet

import (
	"Veloce/internal/interfaces"
	"Veloce/internal/objects/protocol"
	"fmt"
)

type AcknowledgeFinishConfigurationPacket struct { /*No Fields*/
}

func (p *AcknowledgeFinishConfigurationPacket) ID() int32 {
	return 0x03
}

func (p *AcknowledgeFinishConfigurationPacket) Read(*interfaces.Buffer) {
	fmt.Println("Read AcknowledgeFinishConfigurationPacket")
}

type ClientInformationPacket struct {
	locale        string
	render        byte
	chatMode      protocol.VarInt
	chatColor     bool
	skin          byte
	mainHand      protocol.VarInt
	filter        bool
	serverListing bool
	particle      protocol.VarInt
}

func (p *ClientInformationPacket) ID() int32 {
	return 0x00
}

func (p *ClientInformationPacket) Read(buf *interfaces.Buffer) {
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

type PluginMessagePacket struct {
	identifier string
	data       []byte
}

func (p *PluginMessagePacket) ID() int32 {
	return 0x02
}

func (p *PluginMessagePacket) Read(buf *interfaces.Buffer) {
	fmt.Println("Read PluginMessagePacket")
	p.identifier, _ = buf.ReadString()
	p.data = buf.Data()
	fmt.Println(p.identifier)
}

type ServerBoundKnownPacksPacket struct {
	//TODO
}

func (p *ServerBoundKnownPacksPacket) ID() int32 {
	return 0x07
}

func (p *ServerBoundKnownPacksPacket) Read(buf *interfaces.Buffer) {
	fmt.Println("Read ServerBoundKnownPacksPacket")
}
