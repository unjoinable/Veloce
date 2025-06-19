package network

import (
	"Veloce/internal/network/buffer"
)

type Packet interface {
	ID() int32
}

type ServerboundPacket interface {
	Packet
	Read(buf *buffer.Buffer)
	Handle(pc *PlayerConnection)
}

type ClientboundPacket interface {
	Packet
	Write(buf *buffer.Buffer)
}
