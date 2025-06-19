package player

import (
	"Veloce/internal/objects/optional"
	"github.com/google/uuid"
)

type Property struct {
	Name      string
	Value     string
	Signature optional.Optional[string]
}

type GameProfile struct {
	UUID       uuid.UUID
	Name       string
	Properties []Property
}
