package interfaces

// ConnectionState represents the connection state of a client.
type ConnectionState int

const (
	Handshake     ConnectionState = iota // Default state before any packets is received.
	Status                               // Client declares Status intent during handshake.
	Login                                // Client declares Login intent during handshake.
	Configuration                        // Client acknowledged login and is now configuring the game.
	Play                                 // Client (re-)finished configuration.
)

type Packet interface {
	ID() int32
}

type ServerboundPacket interface {
	Packet
	Read(buf *Buffer)
}

type ClientboundPacket interface {
	Packet
	Write(buf *Buffer)
}

type Connection interface {
	SetState(s ConnectionState)
	SendPacket(p ClientboundPacket) error
	SendRaw(data []byte) error
	// Add other methods from PlayerConnection that handlers need to call
}