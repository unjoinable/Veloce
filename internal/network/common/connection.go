package common

import (
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

// SendRaw to send raw bytes
func (pc *PlayerConnection) SendRaw(data []byte) error {
	pc.mu.RLock()
	conn := pc.conn
	pc.mu.RUnlock()

	fmt.Println("sending raw data")

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

	buf := NewBuffer(nil)
	p.Write(buf)

	header := NewBuffer(nil)
	header.WriteVarInt(int32(buf.Len() + 1))
	header.WriteVarInt(p.ID())

	// Combine: header + payload
	final := NewBuffer(nil)
	final.Write(header.Bytes())
	final.Write(buf.Bytes())

	_, err := conn.Write(final.Bytes())
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
