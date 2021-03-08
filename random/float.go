package random

import "math/rand"

// GenerateRandomFloat64InRange generates and returns a number between [min: inclusive] and [max: exclusive]
func GenerateRandomFloat64InRange(min, max float64) float64 {
	return (max-min)*rand.Float64() + min
}
