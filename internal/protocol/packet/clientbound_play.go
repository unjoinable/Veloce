package packet

import (
	"Veloce/internal/interfaces"
	"Veloce/internal/objects/protocol"
)

type LoginPlayPacket struct {
	EntityID            int32
	IsHardcore          bool
	DimensionNames      []string // Identifiers
	MaxPlayers          protocol.VarInt
	ViewDistance        protocol.VarInt
	SimulationDistance  protocol.VarInt
	ReducedDebugInfo    bool
	EnableRespawnScreen bool
	DoLimitedCrafting   bool
	DimensionType       protocol.VarInt
	DimensionName       string // Identifier
	HashedSeed          int64
	GameMode            byte
	PreviousGameMode    byte
	IsDebug             bool
	IsFlat              bool
	HasDeathLocation    bool
	//DeathDimensionName  *string   // optional
	//DeathLocation       *Position // optional struct
	PortalCooldown     protocol.VarInt
	SeaLevel           protocol.VarInt
	EnforcesSecureChat bool
}

func (p *LoginPlayPacket) ID() int32 {
	return 0x2B
}

func (p *LoginPlayPacket) Write(buf *interfaces.Buffer) {
	buf.WriteInt32(p.EntityID)
	buf.WriteBool(p.IsHardcore)

	//prefixed array dimension field we are having 1 dimension rn
	buf.WriteVarInt(1)
	buf.WriteString("minecraft:overworld")

	buf.WriteVarInt(p.MaxPlayers)
	buf.WriteVarInt(p.ViewDistance)
	buf.WriteVarInt(p.SimulationDistance)
	buf.WriteBool(p.ReducedDebugInfo)
	buf.WriteBool(p.EnableRespawnScreen)
	buf.WriteBool(p.DoLimitedCrafting)
	buf.WriteVarInt(p.DimensionType)

	//dimenision name field
	buf.WriteString("minecraft:overworld")

	buf.WriteInt64(p.HashedSeed)
	buf.WriteByte(p.GameMode)
	buf.WriteByte(p.PreviousGameMode)
	buf.WriteBool(p.IsDebug)
	buf.WriteBool(p.IsFlat)
	buf.WriteBool(p.HasDeathLocation)
	//buf.WriteByte(0) // OPTIONAL
	//buf.WriteByte(0) //OPTIONAL
	buf.WriteVarInt(p.PortalCooldown)
	buf.WriteVarInt(p.SeaLevel)
	buf.WriteBool(p.EnforcesSecureChat)
}
