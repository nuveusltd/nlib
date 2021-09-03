package nlib

import (
	"strings"

	"github.com/google/uuid"
)

type Randomizer struct {
}

func (r *Randomizer) NewStringUUID() string {
	return uuid.New().String()
}

func (r *Randomizer) NewStringUUIDWoHypens() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
