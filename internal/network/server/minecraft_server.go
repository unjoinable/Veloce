package server

import (
	"Veloce/internal/entity/player"
	"Veloce/internal/event"
	"Veloce/internal/network"
	"Veloce/internal/protocol"
	"Veloce/internal/scheduler"
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

	packetRegistry *network.PacketRegistry
	scheduler      *scheduler.Scheduler
	ticker         *scheduler.Ticker
	eventNode      *event.Node
}

func NewMinecraftServer() *MinecraftServer {
	registry := network.NewPacketRegistry()
	schedule := scheduler.NewScheduler()

	return &MinecraftServer{
		running:        false,
		playPlayers:    make(map[uuid.UUID]*player.Player),
		packetRegistry: registry,
		scheduler:      schedule,
		ticker:         scheduler.NewTicker(schedule),
		eventNode:      event.NewNode(),
	}
}

func (s *MinecraftServer) Init() {
	protocol.RegisterAllPackets(s.packetRegistry)
}

func (s *MinecraftServer) Start(address string) {
	tcpServer := NewTCPServer(address, s.packetRegistry)

	if err := tcpServer.Start(); err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}

	s.running = true
	s.ticker.Start()
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

func (s *MinecraftServer) GetEventNode() *event.Node {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.eventNode
}

func (s *MinecraftServer) Shutdown() {
	s.running = false
	s.ticker.Shutdown()
	s.scheduler.Shutdown()
}
