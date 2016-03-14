package service

import (
	. "github.com/microbusinesses/Micro-Businesses-Core/system"
)

// Default implementation of UUID generator service
type UUIDGeneratorServiceImpl struct {
}

// Generates random UUID value.
// Either the new random generated UUID or an error if something goes wrong.
func (UUIDGeneratorServiceImpl) GenerateRandomUUID() (UUID, error) {
	return RandomUUID()
}
