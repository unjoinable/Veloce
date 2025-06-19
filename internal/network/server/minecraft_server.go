package server

import (
	"Veloce/internal/network/protocol"
	"Veloce/internal/objects/set"
	"sync"
)

type MinecraftServer struct {
	players set.PlayerSet
	running bool
	mu      sync.RWMutex
}

func init() {
	protocol.RegisterAllPackets()
}
