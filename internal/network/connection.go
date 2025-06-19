package network

import (
	"Veloce/internal/network/buffer"
	"fmt"
	"net"
	"sync"
	"time"
)

// PlayerConnection represents a client connection
type PlayerConnection struct {
	conn  net.Conn
	state ConnectionState
	mu    sync.RWMutex
}

// NewPlayerConnection creates a new player connection
func NewPlayerConnection(conn net.Conn) *PlayerConnection {
	return &PlayerConnection{
		conn:  conn,
		state: Handshake,
	}
}

// HandlePacket processes incoming packets
func (pc *PlayerConnection) HandlePacket(buf *buffer.Buffer) error {
	pc.mu.RLock()
	currentState := pc.state
	pc.mu.RUnlock()

	packetID, err := buf.ReadVarInt()
	if err != nil {
		return fmt.Errorf("failed to read packet ID: %w", err)
	}

	packet, ok := GetServerBoundPacket(currentState, packetID)
	if !ok {
		return fmt.Errorf("unknown packet ID %d for state %v", packetID, currentState)
	}

	if packet == nil {
		return fmt.Errorf("packet not implemented: ID %d, State %v", packetID, currentState)
	}

	packet.Read(buf)
	packet.Handle(pc)
	return nil
}

// SendRaw to send raw bytes
func (pc *PlayerConnection) SendRaw(data []byte) error {
	pc.mu.RLock()
	conn := pc.conn
	pc.mu.RUnlock()

	fmt.Println("sending raw data")
	fmt.Println(data)

	if conn == nil {
		return fmt.Errorf("connection is closed")
	}

	if err := conn.SetWriteDeadline(time.Now().Add(30 * time.Second)); err != nil {
		return err
	}

	_, err := conn.Write(data)
	return err
}

// SendPacket sends a packet to the client
func (pc *PlayerConnection) SendPacket(p ClientboundPacket) error {
	pc.mu.RLock()
	conn := pc.conn
	pc.mu.RUnlock()

	if conn == nil {
		return fmt.Errorf("connection is closed")
	}

	if err := conn.SetWriteDeadline(time.Now().Add(30 * time.Second)); err != nil {
		return err
	}

	buf := buffer.NewBuffer(nil)
	p.Write(buf)

	buffer := buffer.NewBuffer(nil)
	buffer.WriteVarInt(int32(buf.Len() + 1))
	buffer.WriteVarInt(p.ID())
	buffer.WriteBytes(buf.Data())

	_, err := conn.Write(buffer.Data())
	return err
}

// SetState updates the connection state
func (pc *PlayerConnection) SetState(s ConnectionState) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.state = s
}

// GetState returns the current connection state
func (pc *PlayerConnection) GetState() ConnectionState {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	return pc.state
}

// Close closes the player connection
func (pc *PlayerConnection) Close() error {
	pc.mu.Lock()
	defer pc.mu.Unlock()

	if pc.conn != nil {
		err := pc.conn.Close()
		pc.conn = nil
		return err
	}
	return nil
}

// GetRemoteAddr returns the remote address of the connection
func (pc *PlayerConnection) GetRemoteAddr() net.Addr {
	pc.mu.RLock()
	defer pc.mu.RUnlock()

	if pc.conn != nil {
		return pc.conn.RemoteAddr()
	}
	return nil
}

// GetConn returns the underlying net.Conn for direct access
func (pc *PlayerConnection) GetConn() net.Conn {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	return pc.conn
}

// IsConnected checks if the connection is still active
func (pc *PlayerConnection) IsConnected() bool {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	return pc.conn != nil
}
