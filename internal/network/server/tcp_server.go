package server

import (
	"Veloce/internal/interfaces"
	"Veloce/internal/network"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

const (
	MaxPacketLength = 2097151 // Maximum allowed packet length (2^21 - 1)
)

// TCPServer represents a simplified TCP server
type TCPServer struct {
	listener       net.Listener
	addr           string
	running        bool
	connections    sync.Map
	packetRegistry *network.PacketRegistry
}

// NewTCPServer creates a new simplified TCP server
func NewTCPServer(addr string, packetRegistry *network.PacketRegistry) *TCPServer {
	return &TCPServer{
		addr:           addr,
		packetRegistry: packetRegistry,
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
		s.listener.Close()
	}

	s.connections.Range(func(key, value interface{}) bool {
		if pc, ok := value.(*interfaces.PlayerConnection); ok {
			pc.Close()
		}
		return true
	})

	return nil
}

func (s *TCPServer) readPacket(conn net.Conn) (*interfaces.Buffer, error) {
	if err := conn.SetReadDeadline(time.Now().Add(30 * time.Second)); err != nil {
		return nil, fmt.Errorf("failed to set read deadline: %w", err)
	}

	// Create a temporary buffer to read the VarInt length
	tempBuf := make([]byte, 5) // VarInt can be at most 5 bytes
	bytesRead := 0

	// Read VarInt byte by byte
	for i := 0; i < 5; i++ {
		if _, err := conn.Read(tempBuf[bytesRead : bytesRead+1]); err != nil {
			return nil, fmt.Errorf("reading VarInt: %w", err)
		}
		bytesRead++

		// Check if this is the last byte of the VarInt (MSB is 0)
		if tempBuf[bytesRead-1]&0x80 == 0 {
			break
		}
	}

	// Create buffer from the VarInt bytes and read the length
	varintBuf := interfaces.NewBuffer(tempBuf[:bytesRead])
	length, err := varintBuf.ReadVarInt()
	if err != nil {
		return nil, fmt.Errorf("failed to read packet length: %w", err)
	}

	if length <= 0 || length > MaxPacketLength {
		return nil, fmt.Errorf("invalid packet length: %d", length)
	}

	// Read the actual packet data
	packetData := make([]byte, length)
	if _, err := io.ReadFull(conn, packetData); err != nil {
		return nil, fmt.Errorf("reading packet data: %w", err)
	}

	// Return a buffer containing the packet data
	return interfaces.NewBuffer(packetData), nil
}

func (s *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	pc := interfaces.NewPlayerConnection(conn)
	connID := conn.RemoteAddr().String()
	s.connections.Store(connID, pc)
	defer s.connections.Delete(connID)

	for s.running {
		packetBuf, err := s.readPacket(conn)
		if err != nil {
			fmt.Printf("Error reading packet from %s: %v\n", conn.RemoteAddr(), err)
			continue
		}

		currentState := pc.GetState()
		packetId, err := packetBuf.ReadVarInt()

		if err != nil {
			fmt.Printf("failed to read packet ID: %s", err)
			continue
		}

		packet, _ := s.packetRegistry.GetServerBoundPacket(currentState, packetId)
		packet.Handle(pc)
	}
}
