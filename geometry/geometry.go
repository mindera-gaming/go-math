package geometry

import (
	"math"

	"github.com/mindera-gaming/go-math/geometry/sweepline"
	vector "github.com/mindera-gaming/go-math/vector2"
)

// IsSimplePolygon determines whether the polygon is simple or complex.
func IsSimplePolygon(vertices []vector.Vector2) bool {
	if len(vertices) < 3 {
		return false
	}

	// firstly it will check for overlapping vertices
	for i := 1; i < len(vertices); i++ {
		if vertices[i] == vertices[i-1] {
			// there is an overlap
			return false
		}
	}

	// finding intersections using the sweep-line algorithm
	segments := make([]sweepline.Segment, len(vertices))
	for i := 1; i < len(vertices); i++ {
		segments[i-1] = sweepline.NewSegment(vertices[i-1], vertices[i])
	}
	lastIndex := len(vertices) - 1
	segments[lastIndex] = sweepline.NewSegment(vertices[lastIndex], vertices[0])

	intersections := sweepline.FindIntersections(segments)

	return len(intersections) <= 0
}

// ContainsColinearEdges determines if the polygon contains collinear edges.
func ContainsColinearEdges(vertices []vector.Vector2) bool {
	if len(vertices) < 3 {
		return false
	}

	// determining if 3 points are collinear using the distance formula
	for i := 0; i < len(vertices)-2; i++ {
		a := vertices[i]
		b := vertices[i+1]
		c := vertices[i+2]

		abDist := a.Distance(b)
		cbDist := c.Distance(b)
		acDist := a.Distance(c)

		if (abDist + cbDist) == acDist {
			return true
		}
	}
	return false
}

// ComputePolygonArea determines the winding order of a polygon.
// Returns the area of the polygon and its winding order.
func ComputePolygonArea(vertices []vector.Vector2) (float64, WindingOrder) {
	if len(vertices) < 3 {
		return 0, Invalid
	}

	var area float64
	var windingOrder WindingOrder = Invalid

	for i := 1; i < len(vertices); i++ {
		start := vertices[i-1]
		end := vertices[i]
		area += (end.X - start.X) * (end.Y + start.Y)
	}
	// closing the polygon
	start := vertices[len(vertices)-1]
	end := vertices[0]
	area += (end.X - start.X) * (end.Y + start.Y)

	if area > 0 {
		windingOrder = Clockwise
	}
	if area < 0 {
		windingOrder = CounterClockwise
	}

	return math.Abs(area / 2), windingOrder
}

// IsPointInTriangle verifies if the point p is inside the triangle abc.
func IsPointInTriangle(p, a, b, c vector.Vector2) bool {
	ab := a.To(b)
	bc := b.To(c)
	ca := c.To(a)

	ap := a.To(p)
	bp := b.To(p)
	cp := c.To(p)

	cross1 := ab.Cross(ap)
	cross2 := bc.Cross(bp)
	cross3 := ca.Cross(cp)

	if cross1 > 0 || cross2 > 0 || cross3 > 0 {
		return false
	}

	return true
}
