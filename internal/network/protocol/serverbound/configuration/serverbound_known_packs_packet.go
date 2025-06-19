package configuration

import (
	"Veloce/internal/network"
	"Veloce/internal/network/buffer"
	"Veloce/internal/network/protocol/clientbound/configuration"
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

func (p *ServerBoundKnownPacksPacket) Read(buf *buffer.Buffer) {
	fmt.Println("Read ServerBoundKnownPacksPacket")
}

func (p *ServerBoundKnownPacksPacket) Handle(pc *network.PlayerConnection) {
	fmt.Println("ServerBoundKnownPacksPacket Handle")

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
	packet := &configuration.FinishConfigurationPacket{}
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
