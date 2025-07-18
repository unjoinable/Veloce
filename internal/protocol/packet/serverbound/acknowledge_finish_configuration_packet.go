package serverbound

import (
	common2 "Veloce/internal/network/common"
	"Veloce/internal/protocol/packet/clientbound"
)

type AcknowledgeFinishConfigurationPacket struct {
	// No Fields
}

func (p *AcknowledgeFinishConfigurationPacket) ID() int32 {
	return 0x03
}

func (p *AcknowledgeFinishConfigurationPacket) Read(*common2.Buffer) {
	// Nothing to read
}

func (p *AcknowledgeFinishConfigurationPacket) Handle(pc *common2.PlayerConnection) {
	pc.SetState(common2.Play)

	packet := &clientbound.LoginPlayPacket{
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
	_ = pc.SendPacket(packet)
}
