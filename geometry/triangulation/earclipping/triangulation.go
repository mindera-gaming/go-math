package earclipping

import (
	"container/list"
	"math"

	"github.com/mindera-gaming/go-math/geometry/sweepline"
	vector "github.com/mindera-gaming/go-math/vector2"
)

const (
	// No specific reason for choosing this value.
	// Just to limit the maximum number of vertices.
	maxVertices = 1000
)

// TriangulationOptions is the structure that defines the triangulation options.
type TriangulationOptions struct {
	SkipSimplePolygonValidation bool // set to true to skip the simple polygon verification
	SkipColinearEdgesValidation bool // set to true to skip the collinear edge verification
	SkipWindingOrderValidation  bool // set to true to skip the winding order verification
}

// Triangulate decomposes a simple polygon into a set of triangles.
//
// Receives the set of vertices of a polygon and the triagulation options.
// Returns the set of indices of the vertices of the calculated triangles and the area of the polygon.
func Triangulate(vertices []vector.Vector2, options TriangulationOptions) (triangles []int, area float64, err error) {
	if vertices == nil {
		err = ErrNilVertices
		return
	}
	if len(vertices) < 3 {
		err = ErrInsufficientVertices
		return
	}
	if len(vertices) > maxVertices {
		err = ErrExceededVertices
		return
	}
	if !options.SkipSimplePolygonValidation {
		if !isSimplePolygon(vertices) {
			err = ErrNotSimplePolygon
			return
		}
	}
	if !options.SkipColinearEdgesValidation {
		if containsColinearEdges(vertices) {
			err = ErrColinearEdges
			return
		}
	}

	if !options.SkipWindingOrderValidation {
		var order windingOrder
		area, order = computePolygonArea(vertices)
		if order == invalid {
			err = ErrInvalidWindingOrder
			return
		}
		if order == counterClockwise {
			reverseVertices(&vertices)
		}
	}

	indexList := list.New()
	for i := 0; i < len(vertices); i++ {
		indexList.PushBack(i)
	}

	totalTriangleCount := len(vertices) - 2
	totalTriangleIndexCount := totalTriangleCount * 3

	triangles = make([]int, totalTriangleIndexCount)
	triangleIndexCount := 0

	for indexList.Len() > 3 {
		for testIndex := indexList.Front(); testIndex != nil; testIndex = testIndex.Next() {
			current := testIndex.Value.(int)
			previous := getPreviousElement(*indexList, *testIndex).Value.(int)
			next := getNextElement(*indexList, *testIndex).Value.(int)

			currentVector := vertices[current]
			previousVector := vertices[previous]
			nextVector := vertices[next]

			currentToPreviousVector := currentVector.To(previousVector)
			currentToNextVector := currentVector.To(nextVector)

			// checks if the test vertex is convex
			if currentToPreviousVector.Cross(currentToNextVector) <= 0 {
				continue
			}

			isEar := true

			// checks if the test ear contains any polygon vertices
			for j := 0; j < len(vertices); j++ {
				if j == current || j == previous || j == next {
					continue
				}

				p := vertices[j]
				if isPointInTriangle(p, previousVector, currentVector, nextVector) {
					isEar = false
					break
				}
			}

			if isEar {
				triangles[triangleIndexCount] = previous
				triangles[triangleIndexCount+1] = current
				triangles[triangleIndexCount+2] = next
				triangleIndexCount += 3

				// remove the ear found
				indexList.Remove(testIndex)
				break
			}
		}
	}

	index := indexList.Front()
	triangles[triangleIndexCount] = index.Value.(int)
	triangles[triangleIndexCount+1] = index.Next().Value.(int)
	triangles[triangleIndexCount+2] = index.Next().Next().Value.(int)

	return
}

// isSimplePolygon determines whether the polygon is simple or complex.
func isSimplePolygon(vertices []vector.Vector2) bool {
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

// containsColinearEdges determines if the polygon contains collinear edges.
func containsColinearEdges(vertices []vector.Vector2) bool {
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

// computePolygonArea determines the winding order of a polygon.
// Returns the area of the polygon and its winding order.
func computePolygonArea(vertices []vector.Vector2) (float64, windingOrder) {
	var area float64
	var windingOrder windingOrder = invalid

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
		windingOrder = clockwise
	}
	if area < 0 {
		windingOrder = counterClockwise
	}

	return math.Abs(area / 2), windingOrder
}

// isPointInTriangle verifies if the point p is inside the triangle abc.
func isPointInTriangle(p, a, b, c vector.Vector2) bool {
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
