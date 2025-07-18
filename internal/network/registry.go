package network

import (
	"Veloce/internal/network/common"
	"sync"
)

type PacketRegistry struct {
	mu          sync.RWMutex
	serverBound map[common.ConnectionState]map[int32]func() common.ServerboundPacket
}

func NewPacketRegistry() *PacketRegistry {
	return &PacketRegistry{
		serverBound: make(map[common.ConnectionState]map[int32]func() common.ServerboundPacket),
	}
}

// RegisterServerBound registers a packet type we can receive from clients
func (r *PacketRegistry) RegisterServerBound(state common.ConnectionState, id int32, factory func() common.ServerboundPacket) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.serverBound[state] == nil {
		r.serverBound[state] = make(map[int32]func() common.ServerboundPacket)
	}
	r.serverBound[state][id] = factory
}

// CreateServerBound creates a new server-bound packet instance
func (r *PacketRegistry) CreateServerBound(state common.ConnectionState, id int32) (common.ServerboundPacket, bool) {
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
func (r *PacketRegistry) GetServerBoundPacket(state common.ConnectionState, id int32) (common.ServerboundPacket, bool) {
	return r.CreateServerBound(state, id)
}
