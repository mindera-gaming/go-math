package earclipping

import (
	"container/list"

	. "github.com/mindera-gaming/go-math/geometry"
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
// Returns the set of indices of the vertices of the calculated triangles.
func Triangulate(vertices []vector.Vector2, options TriangulationOptions) (triangles []int, err error) {
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
		if !IsSimplePolygon(vertices) {
			err = ErrNotSimplePolygon
			return
		}
	}
	if !options.SkipColinearEdgesValidation {
		if ContainsColinearEdges(vertices) {
			err = ErrColinearEdges
			return
		}
	}

	if !options.SkipWindingOrderValidation {
		var order WindingOrder
		_, order = ComputePolygonArea(vertices)
		if order == Invalid {
			err = ErrInvalidWindingOrder
			return
		}
		if order == CounterClockwise {
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

			// checks if the test vertex is convex (valid)
			if currentToPreviousVector.Cross(currentToNextVector) <= 0 {
				// test vertex is reflex (not valid)
				continue
			}

			isEar := true

			// checks if the test ear contains any polygon vertices
			for testEarIndex := indexList.Front(); testEarIndex != nil; testEarIndex = testEarIndex.Next() {
				earIndexValue := testEarIndex.Value.(int)

				if earIndexValue == current || earIndexValue == previous || earIndexValue == next {
					continue
				}

				p := vertices[earIndexValue]
				if IsPointInTriangle(p, previousVector, currentVector, nextVector) {
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
