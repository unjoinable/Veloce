package server

import (
	"Veloce/internal/network"
	"Veloce/internal/network/buffer"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

// TCPServer represents a simplified TCP server
type TCPServer struct {
	listener    net.Listener
	addr        string
	running     bool
	connections sync.Map
}

// NewTCPServer creates a new simplified TCP server
func NewTCPServer(addr string) *TCPServer {
	return &TCPServer{
		addr: addr,
	}
}

func (s *TCPServer) Start() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	s.listener = listener
	s.running = true
	fmt.Printf("TCP Server started on %s\n", s.addr)

	for s.running {
		conn, err := listener.Accept()
		if err != nil {
			if s.running {
				fmt.Printf("Error accepting connection: %v\n", err)
			}
			continue
		}

		go s.handleConnection(conn)
	}
	return nil
}

func (s *TCPServer) Shutdown() error {
	s.running = false

	if s.listener != nil {
	_:
		s.listener.Close()
	}

	s.connections.Range(func(key, value interface{}) bool {
		if pc, ok := value.(*network.PlayerConnection); ok {
		_:
			pc.Close()
		}
		return true
	})

	return nil
}

func (s *TCPServer) readRawPacket(conn net.Conn) ([]byte, error) {
	if err := conn.SetReadDeadline(time.Now().Add(30 * time.Second)); err != nil {
		return nil, fmt.Errorf("failed to set read deadline: %w", err)
	}

	var (
		length     int
		multiplier = 1
		value      = 0
		buf        [1]byte
	)

	for i := 0; i < 5; i++ {
		if _, err := conn.Read(buf[:]); err != nil {
			return nil, fmt.Errorf("reading VarInt: %w", err)
		}
		b := buf[0]
		value |= int(b&0x7F) * multiplier
		if b&0x80 == 0 {
			length = value
			break
		}
		multiplier <<= 7
	}

	if length <= 0 || length > 2097151 {
		return nil, fmt.Errorf("invalid packet length: %d", length)
	}

	packetData := make([]byte, length)
	if _, err := io.ReadFull(conn, packetData); err != nil {
		return nil, fmt.Errorf("reading packet data: %w", err)
	}

	return packetData, nil
}

func (s *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	pc := network.NewPlayerConnection(conn)
	connID := conn.RemoteAddr().String()
	s.connections.Store(connID, pc)
	defer s.connections.Delete(connID)

	for s.running {
		rawPacket, err := s.readRawPacket(conn)

		if err != nil {
			continue // Skip if invalid packet
		}

		buf := buffer.NewBuffer(rawPacket)

		if err := pc.HandlePacket(buf); err != nil {
			fmt.Printf("Skipping Packet! Error handling packet from %s: %v\n", conn.RemoteAddr(), err)
			continue // Skip if we haven't registered it
		}
	}
}
