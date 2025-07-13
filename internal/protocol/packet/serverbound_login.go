package packet

import (
	"Veloce/internal/interfaces"
	"github.com/google/uuid"
)

type LoginAcknowledgedPacket struct { /*No Fields*/
}

func (p *LoginAcknowledgedPacket) ID() int32 {
	return 0x03
}

func (p *LoginAcknowledgedPacket) Read(_ *interfaces.Buffer) { /*Nothing to read*/ }

type LoginStartPacket struct {
	Username string
	Uuid     uuid.UUID
}

func (h *LoginStartPacket) ID() int32 {
	return 0x00
}

func (h *LoginStartPacket) Read(buf *interfaces.Buffer) {
	h.Username, _ = buf.ReadString()
	h.Uuid, _ = buf.ReadUUID()
}
