package server

import (
	"Veloce/internal/network"
	"Veloce/internal/network/buffer"
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

const (
	ReadTimeout     = 30 * time.Second
	WriteTimeout    = 30 * time.Second
	MaxPacketLength = 2097151 // 2^21 - 1 (Minecraft protocol limit)
	VarIntMaxBytes  = 5
)

// ConnectionHandler defines the interface for handling connection events
type ConnectionHandler interface {
	HandleConnect(playerConn *network.PlayerConnection)
	HandleDisconnect(playerConn *network.PlayerConnection)
	HandlePacket(playerConn *network.PlayerConnection, data []byte) error
}

// TCPServer handles low-level TCP networking operations
type TCPServer struct {
	addr        string
	listener    net.Listener
	running     atomic.Bool
	connections sync.Map // map[string]*network.PlayerConnection
	handler     ConnectionHandler

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

// NewTCPServer creates a new TCP server instance
func NewTCPServer(addr string, handler ConnectionHandler) *TCPServer {
	ctx, cancel := context.WithCancel(context.Background())
	return &TCPServer{
		addr:    addr,
		handler: handler,
		ctx:     ctx,
		cancel:  cancel,
	}
}

// Start begins listening for TCP connections
func (s *TCPServer) Start() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to start TCP server: %w", err)
	}

	s.listener = listener
	s.running.Store(true)

	fmt.Printf("TCP Server listening on %s\n", s.addr)

	s.wg.Add(1)
	go s.acceptConnections()

	return nil
}

// acceptConnections handles the main accept loop
func (s *TCPServer) acceptConnections() {
	defer s.wg.Done()

	for s.isRunning() {
		conn, err := s.listener.Accept()
		if err != nil {
			if s.isShuttingDown() {
				return
			}
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		s.wg.Add(1)
		go s.handleConnection(conn)
	}
}

// handleConnection manages a single client connection
func (s *TCPServer) handleConnection(rawConn net.Conn) {
	defer s.wg.Done()
	defer rawConn.Close()

	// Create player connection wrapper
	playerConn := network.NewPlayerConnection(rawConn)
	connID := s.generateConnectionID(rawConn)

	s.registerConnection(connID, playerConn)
	defer s.unregisterConnection(connID)

	// Notify handler of new connection
	if s.handler != nil {
		s.handler.HandleConnect(playerConn)
	}

	defer func() {
		if s.handler != nil {
			s.handler.HandleDisconnect(playerConn)
		}
	}()

	s.processPackets(playerConn)
}

// processPackets handles the packet reading loop for a connection
func (s *TCPServer) processPackets(playerConn *network.PlayerConnection) {
	for s.isRunning() && playerConn.IsConnected() {
		if s.isShuttingDown() {
			return
		}

		packetData, err := s.readPacket(playerConn)
		if err != nil {
			if !s.isConnectionClosed(err) {
				fmt.Printf("Error reading packet from %s: %v\n",
					playerConn.GetRemoteAddr(), err)
			}
			return
		}

		if err := s.handlePacket(playerConn, packetData); err != nil {
			fmt.Printf("Error handling packet from %s: %v\n",
				playerConn.GetRemoteAddr(), err)
			return
		}
	}
}

// handlePacket processes a single packet through the handler
func (s *TCPServer) handlePacket(playerConn *network.PlayerConnection, data []byte) error {
	if s.handler == nil {
		return nil
	}
	return s.handler.HandlePacket(playerConn, data)
}

// readPacket reads a complete length-prefixed packet
func (s *TCPServer) readPacket(playerConn *network.PlayerConnection) ([]byte, error) {
	// Get underlying connection for direct reading
	conn := playerConn.GetConn()
	if conn == nil {
		return nil, fmt.Errorf("connection is closed")
	}

	conn.SetReadDeadline(time.Now().Add(ReadTimeout))

	length, err := s.readPacketLength(conn)
	if err != nil {
		return nil, err
	}

	if err := s.validatePacketLength(length); err != nil {
		return nil, err
	}

	return s.readPacketData(conn, length)
}

// readPacketLength reads and parses the VarInt length prefix
func (s *TCPServer) readPacketLength(conn net.Conn) (int, error) {
	var lengthBytes []byte

	for i := 0; i < VarIntMaxBytes; i++ {
		b := make([]byte, 1)
		if _, err := conn.Read(b); err != nil {
			return 0, err
		}
		lengthBytes = append(lengthBytes, b[0])

		// Check if this is the last byte (MSB not set)
		if (b[0] & 0x80) == 0 {
			break
		}
	}

	buf := buffer.NewBuffer(lengthBytes)
	length, err := buf.ReadVarInt()
	if err != nil {
		return 0, fmt.Errorf("failed to parse packet length: %w", err)
	}

	return length, nil
}

// validatePacketLength ensures the packet length is within acceptable bounds
func (s *TCPServer) validatePacketLength(length int) error {
	if length <= 0 || length > MaxPacketLength {
		return fmt.Errorf("invalid packet length: %d", length)
	}
	return nil
}

// readPacketData reads the packet payload
func (s *TCPServer) readPacketData(conn net.Conn, length int) ([]byte, error) {
	data := make([]byte, length)
	_, err := io.ReadFull(conn, data)
	if err != nil {
		return nil, fmt.Errorf("failed to read packet data: %w", err)
	}
	return data, nil
}

// Connection management methods
func (s *TCPServer) generateConnectionID(conn net.Conn) string {
	return fmt.Sprintf("%s-%d", conn.RemoteAddr(), time.Now().UnixNano())
}

func (s *TCPServer) registerConnection(id string, playerConn *network.PlayerConnection) {
	s.connections.Store(id, playerConn)
}

func (s *TCPServer) unregisterConnection(id string) {
	s.connections.Delete(id)
}

// GetConnectionCount returns the number of active connections
func (s *TCPServer) GetConnectionCount() int {
	count := 0
	s.connections.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}

// BroadcastPacket sends a packet to all connected players
func (s *TCPServer) BroadcastPacket(packet network.ClientboundPacket) {
	s.connections.Range(func(key, value interface{}) bool {
		if playerConn, ok := value.(*network.PlayerConnection); ok {
			// Send in goroutine to avoid blocking
			go func(pc *network.PlayerConnection) {
				if err := pc.SendPacket(packet); err != nil {
					fmt.Printf("Error broadcasting packet to %s: %v\n",
						pc.GetRemoteAddr(), err)
				}
			}(playerConn)
		}
		return true
	})
}

// Helper methods for cleaner conditionals
func (s *TCPServer) isRunning() bool {
	return s.running.Load()
}

func (s *TCPServer) isShuttingDown() bool {
	select {
	case <-s.ctx.Done():
		return true
	default:
		return false
	}
}

// IsRunning returns whether the server is currently running
func (s *TCPServer) IsRunning() bool {
	return s.isRunning()
}

// isConnectionClosed checks if error indicates a closed connection
func (s *TCPServer) isConnectionClosed(err error) bool {
	if err == io.EOF {
		return true
	}
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return false // Timeout is not a connection close
	}
	return false
}

// Close gracefully shuts down the TCP server
func (s *TCPServer) Close() error {
	if !s.running.CompareAndSwap(true, false) {
		return nil // Already closed
	}

	fmt.Println("Shutting down TCP server...")

	s.cancel()
	s.closeListener()
	s.closeAllConnections()
	s.wg.Wait()

	fmt.Println("TCP server shutdown complete")
	return nil
}

// closeListener closes the main listener
func (s *TCPServer) closeListener() {
	if s.listener != nil {
		s.listener.Close()
	}
}

// closeAllConnections closes all active connections
func (s *TCPServer) closeAllConnections() {
	s.connections.Range(func(key, value interface{}) bool {
		if playerConn, ok := value.(*network.PlayerConnection); ok {
			playerConn.Close()
		}
		return true
	})
}
