package entity

type GameMode byte

const (
	Survival GameMode = iota
	Creative
	Adventure
	Spectator
)

func (gm GameMode) ID() byte {
	return byte(gm)
}
