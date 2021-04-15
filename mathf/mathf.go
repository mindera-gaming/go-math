package mathf

// Clamp restricts the given value to the range defined by the given minimum and maximum values. Returns the
// minimum if the given value is less than the minimum, or the maximum if it's greater.
func Clamp(v, minimum, maximum float64) float64 {
	if v > maximum {
		return maximum
	}
	if v < minimum {
		return minimum
	}

	return v
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
