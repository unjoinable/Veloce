package network

type Packet interface {
	ID() int32
}

type ServerboundPacket interface {
	Packet
	Read(buf *Buffer)
	Handle(pc *PlayerConnection)
}

type ClientboundPacket interface {
	Packet
	Write(buf *Buffer)
}
