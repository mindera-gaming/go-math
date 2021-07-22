// Package mathf provides a set of utility constants and functions for manipulating floats
package mathf

import "math"

const (
	// Resolved value: 2.220446049250313e-16
	//
	// Epsilon64 contains the 64-bit machine epsilon (eps) and gives the distance (gap) between 1 and the next
	// largest 64-bit floating point number. It represents the upper bound on the relative error caused by rounding in
	// floating point arithmetic (addition, subtraction, multiplication and division).
	//
	// The IEEE expression for:
	//	- 7/3 is 1.0010101010101010101010101010101010101010101010101011*2^1
	//	- 4/3 is 1.0101010101010101010101010101010101010101010101010101*2^0
	// As such,
	//	  7/3 - 4/3 =
	//	= 1.0000000000000000000000000000000000000000000000000001*2^0 =
	//	= 1 + eps
	//
	// Sourced from Problem 3 at http://rstudio-pubs-static.s3.amazonaws.com/13303_daf1916bee714161ac78d3318de808a9.html
	Epsilon64 = float64(7)/3 - 4./3 - 1
	// InvPi contains the 1 / π
	InvPi = 0.31830988618379067153776752674502872406891929148091289749533469
)

// Clamp restricts the given value to the range defined by the given minimum and maximum values. Returns the
// minimum if the given value is less than the minimum, or the maximum if it's greater.
func Clamp(v, minimum, maximum float64) float64 {
	// special cases
	switch {
	case math.IsInf(minimum, 1): // if minimum is positive infinity
		return math.Inf(1)
	case math.IsInf(maximum, -1): // if maximum is negative infinity
		return math.Inf(-1)
	case math.IsNaN(v) || math.IsNaN(minimum) || math.IsNaN(maximum):
		return math.NaN()
	case minimum == 0 && minimum == maximum: // in the case minimum and maximum both equal -0 and/or +0
		// if maximum == -0 or v > +0
		if math.Signbit(maximum) || v > maximum {
			return maximum
		}

		return minimum
	}

	if v > maximum {
		return maximum
	}
	if v < minimum {
		return minimum
	}

	return v
}

// Lerp linearly interpolates between two float64, a and b, by amount t. The parameter t is clamped to the
// range [0, 1]. The returned float64 will represent a float64 some fraction t of the way between a and b.
func Lerp(a, b, t float64) float64 {
	return LerpUnclamped(a, b, Clamp(t, 0, 1))
}

// LerpUnclamped linearly interpolates between two float64, a and b, by amount t. The returned float64
// will represent a float64 some fraction t of the way between a and b, if t is in the range [0, 1].
func LerpUnclamped(a, b, t float64) float64 {
	// special cases
	switch {
	case math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(t, 1):
		return math.Inf(1)
	case math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(t, -1):
		return math.Inf(-1)
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(t):
		return math.NaN()
	}

	return t*(b-a) + a
}

// NormalizeAngle normalizes an angle in radians to the range [-π, π]
func NormalizeAngle(angle float64) float64 {
	// special cases
	if math.IsInf(angle, 0) {
		return angle
	} else if math.IsNaN(angle) {
		return math.NaN()
	}

	return math.Mod(angle, math.Pi) - float64(int64(angle*InvPi)%2)*math.Pi
}

// Min returns the smallest number
func Min(a, b float64) float64 {
	if a <= b {
		return a
	}
	return b
}

// Max returns the largest number
func Max(a, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}

// QuadraticFormula provides the solution(s) to a quadratic equation.
// Given a general quadratic equation of the form `ax² + bx + c = 0`
// with x representing an unknown, a, b and c representing constants.
// This function returns `x1` (+) and `x2` (-)
func QuadraticFormula(a, b, c float64) (x1, x2 float64) {
	discriminant := math.Sqrt((b * b) - (4 * a * c))
	x1 = (-b + discriminant) / (2 * a)
	x2 = (-b - discriminant) / (2 * a)
	return
}
