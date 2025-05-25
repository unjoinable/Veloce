package clientbound

import (
	"Veloce/internal/network"
	"fmt"
)

const JSON_EXAMPLE = `{
    "version": {
        "name": "1.21.5",
        "protocol": 770
    },
    "players": {
        "max": 100,
        "online": 5,
        "sample": [
            {
                "name": "thinkofdeath",
                "id": "4566e69f-c907-48ee-8d71-d7ba5aa00d20"
            }
        ]
    },
    "description": {
        "text": "Hello, world!"
    },
    "favicon": "data:image/png;base64,<data>",
    "enforcesSecureChat": false
}`

// StatusResponsePacket represents a status response to a StatusRequestPacket
type StatusResponsePacket struct{}

func (p *StatusResponsePacket) ID() int32 {
	return 0x00
}

func (p *StatusResponsePacket) Write(buf *network.Buffer) {
	fmt.Println("Write StatusResponsePacket")
	buf.WriteString(JSON_EXAMPLE)
}

// PongPacket represents a pong response to a ping request
type PongPacket struct {
	Number int64
}

func (p *PongPacket) ID() int32 {
	return 0x01
}

func (p *PongPacket) Write(buf *network.Buffer) {
	fmt.Println("Write PongPacket")
	buf.WriteInt64(p.Number)
}
