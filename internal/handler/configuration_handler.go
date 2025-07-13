package handler

import (
	"Veloce/internal/interfaces"
	"Veloce/internal/protocol/packet"
	"fmt"
	"os"
	"path/filepath"
)

type AcknowledgeFinishConfigurationPacketHandler struct{}

func (h *AcknowledgeFinishConfigurationPacketHandler) HandlePacket(pc interfaces.Connection, p interfaces.ServerboundPacket) {
	pc.SetState(interfaces.Play)

	packet := &packet.LoginPlayPacket{
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

type ClientInformationPacketHandler struct{}

func (h *ClientInformationPacketHandler) HandlePacket(pc interfaces.Connection, p interfaces.ServerboundPacket) {
}

type PluginMessagePacketHandler struct{}

func (h *PluginMessagePacketHandler) HandlePacket(pc interfaces.Connection, p interfaces.ServerboundPacket) {
}

type ServerBoundKnownPacksPacketHandler struct{}

func (h *ServerBoundKnownPacksPacketHandler) HandlePacket(pc interfaces.Connection, p interfaces.ServerboundPacket) {
	// Send raw contents of line_0.bin to line_19.bin
	for i := 0; i <= 19; i++ {
		data, err := ReadLineBinFile(i)
		if err != nil {
			fmt.Printf("failed to read line_%d.bin: %v\n", i, err)
			continue // skip this one and keep going
		}

		if err := pc.SendRaw(data); err != nil {
			fmt.Printf("failed to send raw data for line_%d.bin: %v\n", i, err)
			continue
		}
	}

	// After sending raw bin files, send the actual configuration packet
	packet := &packet.FinishConfigurationPacket{}
	pc.SendPacket(packet)
}

// ReadLineBinFile reads the content of "reg/line_<index>.bin" and returns it as []byte.
func ReadLineBinFile(index int) ([]byte, error) {
	filename := filepath.Join("reg", fmt.Sprintf("line_%d.bin", index))
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", filename, err)
	}
	return data, nil
}
