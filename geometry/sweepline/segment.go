package sweepline

import (
	vector "github.com/mindera-gaming/go-math/vector2"
)

// Segment defines a line segment
type Segment struct {
	A     vector.Vector2
	B     vector.Vector2
	Value float64
}

// NewSegment returns a new segment with the indicated points
func NewSegment(pointA, pointB vector.Vector2) Segment {
	s := Segment{
		A: pointA,
		B: pointB,
	}
	s.CalculateValue(s.First().X)
	return s
}

// First returns the first point of the segment
func (s Segment) First() vector.Vector2 {
	if s.A.X <= s.B.X {
		return s.A
	} else {
		return s.B
	}
}

// Second returns the Second point of the segment
func (s Segment) Second() vector.Vector2 {
	if s.A.X <= s.B.X {
		return s.B
	} else {
		return s.A
	}
}

// CalculateValue calculates the segment value
func (s *Segment) CalculateValue(value float64) {
	first := s.First()
	second := s.Second()
	s.Value = first.Y + (((second.Y - first.Y) / (second.X - first.X)) * (value - first.X))
}

// Equals checks whether two segments are equal or not
func (s Segment) Equals(segment Segment) bool {
	if s.Value != segment.Value {
		return false
	}
	if s.A.X != segment.A.X || s.A.Y != segment.A.Y {
		return false
	}
	if s.B.X != segment.B.X || s.B.Y != segment.B.Y {
		return false
	}
	return true
}
