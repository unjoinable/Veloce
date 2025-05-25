package network

// ConnectionState represents the connection state of a client.
type ConnectionState int

const (
	Handshake     ConnectionState = iota // Default state before any packets is received.
	Status                               // Client declares Status intent during handshake.
	Login                                // Client declares Login intent during handshake.
	Configuration                        // Client acknowledged login and is now configuring the game.
	Play                                 // Client (re-)finished configuration.
)
