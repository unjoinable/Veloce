package set

import (
	"Veloce/internal/entity/player"
	"github.com/google/uuid"
)

type PlayerSet struct {
	players map[uuid.UUID]*player.Player
}

func NewPlayerSet() *PlayerSet {
	return &PlayerSet{
		players: make(map[uuid.UUID]*player.Player),
	}
}

func (ps *PlayerSet) Add(p *player.Player) {
	ps.players[p.GetUUID()] = p
}

func (ps *PlayerSet) Remove(p *player.Player) {
	delete(ps.players, p.GetUUID())
}

func (ps *PlayerSet) Contains(p *player.Player) bool {
	_, exists := ps.players[p.GetUUID()]
	return exists
}

func (ps *PlayerSet) Get(id uuid.UUID) (*player.Player, bool) {
	player2, ok := ps.players[id]
	return player2, ok
}

func (ps *PlayerSet) Len() int {
	return len(ps.players)
}

func (ps *PlayerSet) Values() []*player.Player {
	vals := make([]*player.Player, 0, len(ps.players))
	for _, player2 := range ps.players {
		vals = append(vals, player2)
	}
	return vals
}

func (ps *PlayerSet) Clear() {
	ps.players = make(map[uuid.UUID]*player.Player)
}

func (ps *PlayerSet) IDs() []uuid.UUID {
	ids := make([]uuid.UUID, 0, len(ps.players))
	for id := range ps.players {
		ids = append(ids, id)
	}
	return ids
}
