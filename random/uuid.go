package random

import (
	"github.com/google/uuid"
)

// GenerateRandomUUID generates and returns a Universally Unique Identifier
func GenerateRandomUUID() (uuid.UUID, error) {
	return uuid.NewRandom()
}
