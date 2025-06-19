package configuration

import (
	"Veloce/internal/network"
	"Veloce/internal/network/buffer"
	"Veloce/internal/network/protocol/clientbound/play"
	"fmt"
)

type AcknowledgeFinishConfigurationPacket struct { /*No Fields*/
}

func (p *AcknowledgeFinishConfigurationPacket) ID() int32 {
	return 0x03
}

func (p *AcknowledgeFinishConfigurationPacket) Read(*buffer.Buffer) {
	fmt.Println("Read AcknowledgeFinishConfigurationPacket")
}

func (p *AcknowledgeFinishConfigurationPacket) Handle(pc *network.PlayerConnection) {
	fmt.Println("Handle AcknowledgeFinishConfigurationPacket")
	pc.SetState(network.Play)

	packet := &play.LoginPlayPacket{
		EntityID:            23,
		IsHardcore:          false,
		DimensionNames:      nil,
		MaxPlayers:          10,
		ViewDistance:        1,
		SimulationDistance:  1,
		ReducedDebugInfo:    false,
		EnableRespawnScreen: false,
		DoLimitedCrafting:   false,
		DimensionType:       0,
		DimensionName:       "minecraft:overworld",
		HashedSeed:          7432018730923847123,
		GameMode:            0,
		PreviousGameMode:    0,
		IsDebug:             false,
		IsFlat:              false,
		HasDeathLocation:    false,
		PortalCooldown:      0,
		SeaLevel:            0,
		EnforcesSecureChat:  false,
	}
	pc.SendPacket(packet)
}
