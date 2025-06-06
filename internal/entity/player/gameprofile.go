package player

import (
	"github.com/google/uuid"
)

type Property struct {
	Name  string
	Value string
}

type GameProfile struct {
	UUID       uuid.UUID
	Name       string
	Properties []Property
}
