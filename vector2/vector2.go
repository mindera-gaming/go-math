// Package vector2 provides a struct for handling 2D vectors, points and positions
package vector2

import (
	"fmt"
	"math"

	"github.com/mindera-gaming/go-math/mathf"
)

const (
	// X represents Vector2's X property JSON tag
	X = "x"
	// Y represents Vector2's Y property JSON tag
	Y = "y"

	jsonFormat = `{"` + X + `":%g,"` + Y + `":%g}`
)

// Vector2 represents a 2D vector, point or position
type Vector2 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Reset resets the vector to zero
func (v *Vector2) Reset() {
	v.X = 0
	v.Y = 0
}

// One returns a vector with a value of 1 in both fields, X and Y. Provides a shorthand for Vector2{X: 1, Y: 1}
func One() Vector2 {
	return Vector2{X: 1, Y: 1}
}

// Right returns a unit vector pointing to the right in world space. Provides a shorthand for Vector2{X: 1}
func Right() Vector2 {
	return Vector2{X: 1}
}

// Left returns a unit vector pointing to the left in world space. Provides a shorthand for Vector2{X: -1}
func Left() Vector2 {
	return Vector2{X: -1}
}

// Up returns a unit vector pointing up in world space. Provides a shorthand for Vector2{Y: 1}
func Up() Vector2 {
	return Vector2{Y: 1}
}

// Down returns a unit vector pointing down in world space. Provides a shorthand for Vector2{Y: -1}
func Down() Vector2 {
	return Vector2{Y: -1}
}

// Dot returns the dot product between this vector and the given vector. For normalized vectors, Dot returns 1 if they point
// in exactly the same direction and -1 if they point in completely opposite directions. For perpendicular vectors,
// their dot product will be 0.
func (v Vector2) Dot(other Vector2) float64 {
	return v.X*other.X + v.Y*other.Y
}

// Cross returns the cross product between this vector and the given vector. As the cross product is not defined in
// 2 dimensions, this calculation treats both as 3D vectors, with a 0 in the Z component. The returned value represents
// the Z component of the resulting cross product vector. In essence:
//   (x1, y1, 0) x (x2, y2, 0) = (0, 0, x1*y2 - y1*x2) => x1*y2 - y1*x2
func (v Vector2) Cross(other Vector2) float64 {
	return v.X*other.Y - v.Y*other.X
}

// Cross returns the cross product between this vector and the given scalar. As the cross product is not defined in
// 2 dimensions, this calculation treats both as 3D vectors, the first with a 0 in the Z component and the other a 0
// in the X and Y components. In essence:
//   (x1, y1, 0) x (0, 0, z2) = (z2 * y1, -z2 * x1, 0)
func (v Vector2) CrossScalar(scalar float64) Vector2 {
	return Vector2{X: scalar * v.Y, Y: -scalar * v.X}
}

// Cross returns the cross product between the given scalar and this vector. As the cross product is not defined in
// 2 dimensions, this calculation treats both as 3D vectors, the other with a 0 in the X and Y components and this one
// with a 0 in the Z component. In essence:
//   (0, 0, z2) x (x1, y1, 0) = (-z2 * y1, z2 * x1, 0)
func (v Vector2) CrossScalarVector(scalar float64) Vector2 {
	return Vector2{X: -scalar * v.Y, Y: scalar * v.X}
}

// To creates a Vector2 from this Vector2 to the given Vector2.
func (v Vector2) To(other Vector2) Vector2 {
	return other.Sub(v)
}

// IsZero returns true if both components of this vector have an absolute value below the float64 eps
func (v Vector2) IsZero() bool {
	return math.Abs(v.X) < mathf.Epsilon64 && math.Abs(v.Y) < mathf.Epsilon64
}

// Left returns a vector orthogonal to this one, rotated 90º to the left. The magnitude is preserved
func (v Vector2) Left() Vector2 {
	return Vector2{X: -v.Y, Y: v.X}
}

// Right returns a vector orthogonal to this one, rotated 90º to the right. The magnitude is preserved
func (v Vector2) Right() Vector2 {
	return Vector2{X: v.Y, Y: -v.X}
}

// Inverse returns a vector pointing in the exact opposite direction to this one. The magnitude is preserved
func (v Vector2) Inverse() Vector2 {
	return Vector2{X: -v.X, Y: -v.Y}
}

// DistanceSqr returns the squared distance between this vector and the other one
func (v Vector2) DistanceSqr(other Vector2) float64 {
	x := other.X - v.X
	y := other.Y - v.Y
	return x*x + y*y
}

// Distance returns the distance between this vector and the other one
func (v Vector2) Distance(other Vector2) float64 {
	return math.Sqrt(v.DistanceSqr(other))
}

// Slope returns a number that describes both the direction and slope of a line formed by this point and the given point
func (v Vector2) Slope(other Vector2) float64 {
	return (v.Y - other.Y) / (v.X - other.X)
}

// MagnitudeSqr returns the squared length of this vector
func (v Vector2) MagnitudeSqr() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Magnitude returns the length of this vector
func (v Vector2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalized returns this vector with a magnitude of 1
func (v Vector2) Normalized() Vector2 {
	mag := v.Magnitude()
	if mag < mathf.Epsilon64 {
		return Vector2{X: math.Inf(0), Y: math.Inf(0)}
	}
	mag = 1 / mag
	return Vector2{X: v.X * mag, Y: v.Y * mag}
}

// ProjectOnto orthogonally projects this vector onto another vector. The resulting vector will share the same
// direction as the other vector (but may be pointing in the opposite sense). This is equivalent to taking the
// component of this vector that is parallel to the other
func (v Vector2) ProjectOnto(other Vector2) Vector2 {
	mag2 := other.MagnitudeSqr()
	if mag2 < mathf.Epsilon64 {
		return Vector2{}
	}
	ratio := v.Dot(other) / mag2
	return Vector2{X: other.X * ratio, Y: other.Y * ratio}
}

// Add adds the other vector to this one component-wise
func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{X: v.X + other.X, Y: v.Y + other.Y}
}

// AddScalar adds the given scalar to both components of this vector
func (v Vector2) AddScalar(value float64) Vector2 {
	return Vector2{X: v.X + value, Y: v.Y + value}
}

// Sub subtracts the other vector to this one component-wise
func (v Vector2) Sub(other Vector2) Vector2 {
	return Vector2{X: v.X - other.X, Y: v.Y - other.Y}
}

// SubScalar subtracts the given scalar to both components of this vector
func (v Vector2) SubScalar(value float64) Vector2 {
	return Vector2{X: v.X - value, Y: v.Y - value}
}

// Scale multiplies the other vector by this one component-wise
func (v Vector2) Scale(other Vector2) Vector2 {
	return Vector2{X: v.X * other.X, Y: v.Y * other.Y}
}

// Mul multiplies the given scalar by both components of this vector
func (v Vector2) Mul(value float64) Vector2 {
	return Vector2{X: v.X * value, Y: v.Y * value}
}

// Div divides both components of this vector by the given scalar
func (v Vector2) Div(value float64) Vector2 {
	value = 1 / value
	return Vector2{X: v.X * value, Y: v.Y * value}
}

// Lerp linearly interpolates between two Vector2, a and b, by amount t. The parameter t is clamped to the
// range [0, 1]. If a and b represent two points, the returned vector will represent a point some fraction t of the way
// along the line segment described by a and b.
func Lerp(a, b Vector2, t float64) Vector2 {
	return LerpUnclamped(a, b, mathf.Clamp(t, 0, 1))
}

// LerpUnclamped linearly interpolates between two Vector2, a and b, by amount t. If a and b represent two
// points, the returned vector will represent a point some fraction t of the way along the line described by a and b.
func LerpUnclamped(a, b Vector2, t float64) Vector2 {
	return a.To(b).Mul(t).Add(a)
}

// Reflect reflects a vector off the plane defined by a normal. The `normal` vector defines a plane (a
// plane's normal is the vector that is perpendicular to its surface). The `direction` vector is treated as a
// directional arrow coming in to the plane. The returned value is a vector of equal magnitude to `direction` but with
// its direction reflected.
func Reflect(direction Vector2, normal Vector2) Vector2 {
	return direction.Sub(normal.Mul(2 * direction.Dot(normal)))
}

func (v Vector2) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(jsonFormat, v.X, v.Y)), nil
}
