package network

import (
	buffer2 "Veloce/internal/network/buffer"
	"fmt"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

// TCPServer represents a simplified TCP server for Minecraft
type TCPServer struct {
	listener    net.Listener
	addr        string
	running     atomic.Bool
	connections sync.Map
}

// NewTCPServer creates a new simplified TCP server
func NewTCPServer(addr string) *TCPServer {
	return &TCPServer{
		addr: addr,
	}
}

// Start begins listening for connections
func (s *TCPServer) Start() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	s.listener = listener
	s.running.Store(true)
	fmt.Printf("TCP Server started on %s\n", s.addr)

	for s.running.Load() {
		conn, err := listener.Accept()
		if err != nil {
			if s.running.Load() {
				fmt.Printf("Error accepting connection: %v\n", err)
			}
			continue
		}

		go s.handleConnection(conn)
	}
	return nil
}

// handleConnection processes a single client connection
func (s *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	pc := NewPlayerConnection(conn)
	connID := fmt.Sprintf("%s-%d", conn.RemoteAddr(), time.Now().UnixNano())
	s.connections.Store(connID, pc)
	defer s.connections.Delete(connID)

	for s.running.Load() {
		packetData, err := s.readPacket(conn)
		if err != nil {
			break
		}
		buffer := buffer2.NewBuffer(packetData)
		if err := pc.HandlePacket(buffer); err != nil {
			fmt.Printf("Error handling packet from %s: %v\n", conn.RemoteAddr(), err)
			break
		}
	}
}

// readPacket reads a complete packet using Buffer class
func (s *TCPServer) readPacket(conn net.Conn) ([]byte, error) {
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))

	// Read VarInt length
	var lengthBytes []byte
	for i := 0; i < 5; i++ {
		b := make([]byte, 1)
		if _, err := conn.Read(b); err != nil {
			return nil, err
		}
		lengthBytes = append(lengthBytes, b[0])
		if (b[0] & 0x80) == 0 {
			break
		}
	}

	// Parse length using Buffer
	buf := buffer2.NewBuffer(lengthBytes)
	length, err := buf.ReadVarInt()
	if err != nil || length <= 0 || length > 2097151 {
		return nil, fmt.Errorf("invalid packet length: %d", length)
	}

	// Read packet data
	packetData := make([]byte, length)
	_, err = io.ReadFull(conn, packetData)
	return packetData, err
}

// Close gracefully shuts down the server
func (s *TCPServer) Close() error {
	if !s.running.CompareAndSwap(true, false) {
		return nil
	}

	if s.listener != nil {
		s.listener.Close()
	}

	s.connections.Range(func(key, value interface{}) bool {
		if pc, ok := value.(*PlayerConnection); ok {
			pc.Close()
		}
		return true
	})

	return nil
}
