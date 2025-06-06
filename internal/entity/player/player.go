package player

import "Veloce/internal/network"

type Player struct {
	pc          *network.PlayerConnection
	gameProfile *GameProfile
}
