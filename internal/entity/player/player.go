package player

import (
	"Veloce/internal/interfaces"
	"Veloce/internal/objects/coordinate"
	"github.com/google/uuid"
	"sync"
)

type Player struct {
	//Internals
	remoteAddr  string
	pc          interfaces.PlayerConnection
	gameProfile GameProfile
	mu          sync.RWMutex

	// Player Fields
	uuid        uuid.UUID
	displayName string
	gameMode    GameMode
	position    coordinate.Position
	velocity    coordinate.Vector
}

func (p *Player) SendPacket(packet interfaces.ClientboundPacket) error {
	p.mu.RLock()
	defer p.mu.RUnlock()
	err := p.pc.SendPacket(packet)
	if err != nil {
		return err
	}
	return nil
}

func (p *Player) SetGameMode(gameMode GameMode) {
	p.mu.Lock()
	p.gameMode = gameMode
	p.mu.Unlock()
}

func (p *Player) GetGameMode() GameMode {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.gameMode
}

func (p *Player) SetVelocity(velocity coordinate.Vector) {
	p.mu.Lock()
	p.velocity = velocity
	p.mu.Unlock()
}

func (p *Player) GetVelocity() coordinate.Vector {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.velocity
}

func (p *Player) GetUUID() uuid.UUID {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.uuid
}
