package network

import (
	"sync"
)

type PacketRegistry struct {
	mu          sync.RWMutex
	serverBound map[ConnectionState]map[int32]func() ServerboundPacket
}

func NewPacketRegistry() *PacketRegistry {
	return &PacketRegistry{
		serverBound: make(map[ConnectionState]map[int32]func() ServerboundPacket),
	}
}

// RegisterServerBound registers a packet type we can receive from clients
func (r *PacketRegistry) RegisterServerBound(state ConnectionState, id int32, factory func() ServerboundPacket) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.serverBound[state] == nil {
		r.serverBound[state] = make(map[int32]func() ServerboundPacket)
	}
	r.serverBound[state][id] = factory
}

// CreateServerBound creates a new server-bound packet instance
func (r *PacketRegistry) CreateServerBound(state ConnectionState, id int32) (ServerboundPacket, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	statePackets, ok := r.serverBound[state]
	if !ok {
		return nil, false
	}
	factory, ok := statePackets[id]
	if !ok {
		return nil, false
	}
	return factory(), true
}

var GlobalPacketRegistry *PacketRegistry

func init() {
	GlobalPacketRegistry = NewPacketRegistry()
}

// GetGlobalPacketRegistry returns the global packet registry instance
func GetGlobalPacketRegistry() *PacketRegistry {
	return GlobalPacketRegistry
}

// GetServerBoundPacket creates a server-bound packet from the global registry
func GetServerBoundPacket(state ConnectionState, id int32) (ServerboundPacket, bool) {
	return GlobalPacketRegistry.CreateServerBound(state, id)
}
