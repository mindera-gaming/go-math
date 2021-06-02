// Package matrix provides a Matrix structure representing internally a 2x2 matrix for handling 2D rotations
package matrix

import (
	"math"

	"github.com/mindera-gaming/go-math/mathf"
	"github.com/mindera-gaming/go-math/rotation"
	vector "github.com/mindera-gaming/go-math/vector2"
)

var reference = vector.Right()

// Matrix represents a 2x2 matrix specifically for handling 2D rotations. As rotations matrices are of the type
//   [ a, b]
//   [-b, a]
// it has been reduced to an array of floats of size 2, [a, b].
type Matrix [2]float64

// Identity returns the identity matrix
//   [1, 0]
//   [0, 1]
// Provides a shorthand for Matrix{1, 0}
func Identity() Matrix {
	return Matrix{1, 0}
}

// Determinant returns the determinant of this matrix. It is used for calculating the inverse matrix to
// this one and can be used to determine how much the rotation matrix diverges from a orthonormalized matrix
func (m Matrix) Determinant() float64 {
	return m[0]*m[0] + m[1]*m[1]
}

// Inverse returns the inverse of this matrix
func (m Matrix) Inverse() Matrix {
	invDeterminant := 1 / m.Determinant()
	m[0] *= invDeterminant
	m[1] *= -invDeterminant

	return m
}

// Euler returns the angle of this rotation matrix in euler units
func (m Matrix) Euler() float64 {
	return m.Radians() * rotation.RadiansToEuler
}

// Radians returns the angle of this rotation matrix in radians
func (m Matrix) Radians() float64 {
	return math.Atan2(-m[1], m[0])
}

// Trace returns the trace of this rotation matrix. A normalized matrix has a trace of 1
func (m Matrix) Trace() float64 {
	return 2 * m[0]
}

// Transpose returns the transpose of this matrix. For a rotation matrix representing an angle α,
// its transpose will represent the angle -α
func (m Matrix) Transpose() Matrix {
	m[1] = -m[1]

	return m
}

// Orthonormalized returns this rotation matrix orthonormalized. An orthonormalized rotation matrix will not scale
// vectors when rotating them
func (m Matrix) Orthonormalized() Matrix {
	invMagnitude := 1 / math.Sqrt(m[0]*m[0]+m[1]*m[1])
	m[0] *= invMagnitude
	m[1] *= invMagnitude

	return m
}

// Angle returns the angle in radians between this rotation matrix and the other given one
func (m Matrix) Angle(other Matrix) float64 {
	return m.Mul(other.Transpose()).Radians()
}

// Dot returns the dot product between this rotation matrix and the other one. For orthonormalized rotation matrices,
// Dot returns 1 if they represent the same rotation and -1 if they represent completely opposite rotations
func (m Matrix) Dot(other Matrix) float64 {
	return reference.X*(m[0]-m[1]+other[0]-other[1]) + reference.Y*(m[0]+m[1]+other[0]+other[1])
}

// Add adds the other rotation matrix to this one
func (m Matrix) Add(other Matrix) Matrix {
	m[0] += other[0]
	m[1] += other[1]

	return m
}

// Add subtracts the other rotation matrix from this one
func (m Matrix) Sub(other Matrix) Matrix {
	m[0] -= other[0]
	m[1] -= other[1]

	return m
}

// Mul multiplies the other rotation matrix by this one
func (m Matrix) Mul(other Matrix) (res Matrix) {
	res[0] = m[0]*other[0] - m[1]*other[1]
	res[1] = m[0]*other[1] + m[1]*other[0]

	return
}

// MulScalar multiplies this rotation matrix by the given scalar
func (m Matrix) MulScalar(scalar float64) Matrix {
	m[0] *= scalar
	m[1] *= scalar

	return m
}

// Rotate rotates the given vector by this rotation matrix
func (m Matrix) Rotate(v vector.Vector2) vector.Vector2 {
	return vector.Vector2{
		X: v.X*m[0] + v.Y*m[1],
		Y: v.Y*m[0] - v.X*m[1],
	}
}

// RotateAround rotates the given vector around the given pivot by this rotation matrix
func (m Matrix) RotateAround(point, pivot vector.Vector2) vector.Vector2 {
	point = point.Sub(pivot)
	return vector.Vector2{
		X: point.X*m[0] + point.Y*m[1] + pivot.X,
		Y: point.Y*m[0] - point.X*m[1] + pivot.Y,
	}
}

// Lerp linearly interpolates between two Matrix, a and b, by amount t. The parameter t is clamped to the
// range [0, 1]. If a and b represent two rotations, the returned Matrix will represent a rotation some linear
// fraction t of the way along the rotation between a and b.
func Lerp(a, b Matrix, t float64) Matrix {
	return LerpUnclamped(a, b, mathf.Clamp(t, 0, 1))
}

// LerpUnclamped linearly interpolates between two Matrix, a and b, by amount t. If a and b represent two
// rotations, the returned Matrix will represent a rotation some linear fraction t of the way along the
// rotation between a and b.
func LerpUnclamped(a, b Matrix, t float64) Matrix {
	return b.Sub(a).MulScalar(t).Add(a)
}

// FromEuler returns a rotation matrix representing the given angle in euler units
func FromEuler(angle float64) Matrix {
	return FromRadians(angle * rotation.EulerToRadians)
}

// FromRadians returns a rotation matrix representing the given angle in radians
func FromRadians(angle float64) Matrix {
	return Matrix{math.Cos(angle), -math.Sin(angle)}
}

// LookRotation returns a rotation matrix representing the direction the given vector points to
func LookRotation(v vector.Vector2) Matrix {
	return FromToRotation(reference, v)
}

// FromToRotation returns a rotation matrix for rotating the vector `from` to the vector `to`
func FromToRotation(from, to vector.Vector2) Matrix {
	return Matrix{from.Dot(to), -from.Cross(to)}.Orthonormalized()
}
