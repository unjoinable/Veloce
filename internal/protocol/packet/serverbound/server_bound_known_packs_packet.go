package serverbound

import (
	"Veloce/internal/interfaces"
	"Veloce/internal/protocol/packet/clientbound"
	"fmt"
	"os"
	"path/filepath"
)

type ServerBoundKnownPacksPacket struct {
	//TODO
}

func (p *ServerBoundKnownPacksPacket) ID() int32 {
	return 0x07
}

func (p *ServerBoundKnownPacksPacket) Read(*interfaces.Buffer) {
	// Nothing to read
}

func (p *ServerBoundKnownPacksPacket) Handle(pc *interfaces.PlayerConnection) {
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
	packet := &clientbound.FinishConfigurationPacket{}
	_ = pc.SendPacket(packet)
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
