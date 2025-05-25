package clientbound

import (
	"Veloce/internal/entity"
	"Veloce/internal/network"
	"fmt"
)

type LoginSuccessPacket struct {
	GameProfile entity.GameProfile
}

func (p *LoginSuccessPacket) ID() int32 {
	return 0x02
}

func (p *LoginSuccessPacket) Write(buf *network.Buffer) {
	fmt.Println("Write LoginSuccessPacket")
	buf.WriteUUID(p.GameProfile.UUID)
	buf.WriteString(p.GameProfile.Name)

	//TODO: A basic way of writing this, I am not sure if this is the correct way.
	ar := p.GameProfile.Properties
	buf.WriteVarInt(int32(len(ar)))
	for i := range ar {
		buf.WriteString(ar[i].Name)
		buf.WriteString(ar[i].Value)
	}
}
