package network

import (
	"sync"

	"Veloce/internal/interfaces"
)

type PacketRegistry struct {
	mu          sync.RWMutex
	serverBound map[interfaces.ConnectionState]map[int32]func() interfaces.ServerboundPacket
}

func NewPacketRegistry() *PacketRegistry {
	return &PacketRegistry{
		serverBound: make(map[interfaces.ConnectionState]map[int32]func() interfaces.ServerboundPacket),
	}
}

// RegisterServerBound registers a packet type we can receive from clients
func (r *PacketRegistry) RegisterServerBound(state interfaces.ConnectionState, id int32, factory func() interfaces.ServerboundPacket) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.serverBound[state] == nil {
		r.serverBound[state] = make(map[int32]func() interfaces.ServerboundPacket)
	}
	r.serverBound[state][id] = factory
}

// CreateServerBound creates a new server-bound packet instance
func (r *PacketRegistry) CreateServerBound(state interfaces.ConnectionState, id int32) (interfaces.ServerboundPacket, bool) {
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

// GetServerBoundPacket creates a server-bound packet from the registry
func (r *PacketRegistry) GetServerBoundPacket(state interfaces.ConnectionState, id int32) (interfaces.ServerboundPacket, bool) {
	return r.CreateServerBound(state, id)
}
