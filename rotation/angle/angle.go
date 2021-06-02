// Package angle provides an Angle structure representing an angle in radians for handling 2D rotations
package angle

import (
	"math"

	"github.com/mindera-gaming/go-math/mathf"
	"github.com/mindera-gaming/go-math/rotation"
	vector "github.com/mindera-gaming/go-math/vector2"
)

var reference = vector.Right()

// Angle represents a 2D rotation in radians
type Angle float64

// Dot returns the dot product between this and the given Angle
func (a Angle) Dot(other Angle) float64 {
	i, j := a.Rotate(reference), other.Rotate(reference)

	return i.Dot(j)
}

// Euler returns the euler angle (in degrees) for this Angle
func (a Angle) Euler() float64 {
	return float64(a * rotation.RadiansToEuler)
}

// Inverse returns the inverse of the Angle normalized to the range [-π, π]
func (a Angle) Inverse() Angle {
	return (-a).Normalized()
}

// Normalized returns the Angle normalized to the range [-π, π]
func (a Angle) Normalized() Angle {
	return Angle(mathf.NormalizeAngle(float64(a)))
}

// Rotate returns the rotated vector
func (a Angle) Rotate(v vector.Vector2) vector.Vector2 {
	angle := float64(a)
	cos := math.Cos(angle)
	sin := math.Sin(angle)

	return vector.Vector2{
		X: v.X*cos - v.Y*sin,
		Y: v.X*sin + v.Y*cos,
	}
}
