package entity

type GameMode byte

const (
	Survival  GameMode = 0
	Creative  GameMode = 1
	Adventure GameMode = 2
	Spectator GameMode = 3
)

func (gm GameMode) ID() byte {
	return byte(gm)
}
