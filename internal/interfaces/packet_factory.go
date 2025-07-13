package interfaces

type PacketFactory interface {
	CreateServerBound(state ConnectionState, id int32) (ServerboundPacket, bool)
}
