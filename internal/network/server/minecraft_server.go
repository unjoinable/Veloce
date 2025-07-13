package server

import (
	"Veloce/internal/entity/player"
	"Veloce/internal/handler"
	"Veloce/internal/interfaces"
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

	packetRegistry   *network.PacketRegistry
	packetDispatcher *handler.PacketDispatcher
	scheduler        scheduler.Scheduler
	ticker           *scheduler.Ticker // Manages game ticks
}

func NewMinecraftServer() *MinecraftServer {
	registry := network.NewPacketRegistry()
	sched := scheduler.NewScheduler()
	return &MinecraftServer{
		running:          false,
		playPlayers:      make(map[uuid.UUID]*player.Player),
		packetRegistry:   registry,
		packetDispatcher: handler.NewPacketDispatcher(registry),
		scheduler:        sched,
		ticker:           scheduler.NewTicker(sched), // Initialize Ticker
	}
}

func (s *MinecraftServer) Init() {
	protocol.RegisterAllPackets(s.packetRegistry)

	// Register Handlers
	s.packetDispatcher.RegisterHandler(interfaces.Handshake, 0x00, &handler.HandshakePacketHandler{})

	s.packetDispatcher.RegisterHandler(interfaces.Status, 0x00, &handler.StatusRequestPacketHandler{})
	s.packetDispatcher.RegisterHandler(interfaces.Status, 0x01, &handler.PingRequestPacketHandler{})

	s.packetDispatcher.RegisterHandler(interfaces.Login, 0x00, &handler.LoginStartPacketHandler{})
	s.packetDispatcher.RegisterHandler(interfaces.Login, 0x03, &handler.LoginAcknowledgedPacketHandler{})

	s.packetDispatcher.RegisterHandler(interfaces.Configuration, 0x00, &handler.ClientInformationPacketHandler{})
	s.packetDispatcher.RegisterHandler(interfaces.Configuration, 0x02, &handler.PluginMessagePacketHandler{})
	s.packetDispatcher.RegisterHandler(interfaces.Configuration, 0x03, &handler.AcknowledgeFinishConfigurationPacketHandler{})
	s.packetDispatcher.RegisterHandler(interfaces.Configuration, 0x07, &handler.ServerBoundKnownPacksPacketHandler{})

	s.packetDispatcher.RegisterHandler(interfaces.Play, 0x0B, &handler.ClientTickEndPacketHandler{})
	s.packetDispatcher.RegisterHandler(interfaces.Play, 0x1C, &handler.MovePlayerPosPacketHandler{})
	s.packetDispatcher.RegisterHandler(interfaces.Play, 0x1D, &handler.MovePlayerPosRotPacketHandler{})
}

func (s *MinecraftServer) Start(address string) {
	tcpServer := NewTCPServer(address, s.packetDispatcher)

	if err := tcpServer.Start(); err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}

	s.running = true
	s.ticker.Start() // Start the game loop
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

func (s *MinecraftServer) Shutdown() {
	s.running = false
	s.ticker.Shutdown()
	s.scheduler.Shutdown()
}
