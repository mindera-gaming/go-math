package random

import (
	"math/big"
	"math/rand"

	"github.com/google/uuid"
)

// GenerateRandomIntInRange generates and returns an integer between [min: inclusive] and [max: exclusive]
func GenerateRandomIntInRange(min, max int) int {
	return rand.Intn(max-min) + min
}

// GenerateRandomPositiveInt64 generates and returns a positive random integer
func GenerateRandomPositiveInt64() (int64, error) {
	r, err := uuid.NewRandom()
	if err != nil {
		return 0, err
	}
	// shaving off the negative bit
	r[0] &= 0x7f
	// an int64 can be represented by 8 bytes
	return bytesToInt64(r[:8]), nil
}

func bytesToInt64(bytes []byte) int64 {
	return (&big.Int{}).SetBytes(bytes).Int64()
}
