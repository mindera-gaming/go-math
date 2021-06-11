package earclipping

import (
	"container/list"

	vector "github.com/mindera-gaming/go-math/vector2"
)

// reverseVertices reverses the order of an array of vertices.
func reverseVertices(array *[]vector.Vector2) {
	for i, j := 0, len(*array)-1; i < j; i, j = i+1, j-1 {
		(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
	}
}

// getPreviousElement returns the previous list element in a circular way.
func getPreviousElement(list list.List, current list.Element) list.Element {
	element := current.Prev()
	if element == nil {
		element = list.Back()
	}
	return *element
}

// getNextElement returns the next list element in a circular way.
func getNextElement(list list.List, current list.Element) list.Element {
	element := current.Next()
	if element == nil {
		element = list.Front()
	}
	return *element
}
