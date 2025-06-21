package server

import (
	"Veloce/internal/entity/player"
	"Veloce/internal/network/protocol"
	"github.com/google/uuid"
	"log"
	"sync"
)

const (
	VersionName         = "1.21.5"
	ProtocolVersion     = 770
	DataVersion         = 4325
	ResourcePackVersion = 55
	DataPackVersion     = 71
)

type MinecraftServer struct {
	tcpServer TCPServer
	running   bool
	mu        sync.RWMutex

	playPlayers map[uuid.UUID]*player.Player
	brandName   string
}

func NewMinecraftServer() *MinecraftServer {
	return &MinecraftServer{
		running:     false,
		playPlayers: make(map[uuid.UUID]*player.Player),
	}
}

func (s *MinecraftServer) Init() {
	protocol.RegisterAllPackets()
}

func (s *MinecraftServer) Start(address string) {
	tcpServer := NewTCPServer(address)

	if err := tcpServer.Start(); err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}

	s.running = true
}

func (s *MinecraftServer) SetBrand(brand string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.brandName = brand
}

func (s *MinecraftServer) GetBrand() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.brandName
}
