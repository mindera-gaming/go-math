package vector2

import "math"

// Point represents a 2D point, vector or position
type Point struct {
	X, Y float64
}

// Add sums the given point to this point
func (p Point) Add(other Point) Point {
	p.X += other.X
	p.Y += other.Y

	return p
}

func (p *Point) Reset() {
	p.X = 0
	p.Y = 0
}

// Slope
func (p Point) Slope(other Point) float64 {
	return (p.Y - other.Y) / (p.X - other.X)
}

// Distance returns the distance between this point and the given point
func (p Point) Distance(other Point) float64 {
	return math.Sqrt((math.Pow((other.X - p.X), 2)) + (math.Pow((other.Y - p.Y), 2)))
}
