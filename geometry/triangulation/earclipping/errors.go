package earclipping

import (
	"errors"
	"fmt"
)

var (
	ErrNilVertices          = errors.New("The vertex list is nil.")
	ErrInsufficientVertices = errors.New("The vertex list must have at least 3 vertices.")
	ErrExceededVertices     = fmt.Errorf("The max vertex list length is %d.", maxVertices)
	ErrNotSimplePolygon     = errors.New("The vertex list does not define a simple polygon.")
	ErrColinearEdges        = errors.New("The vertex list contains colinear edges.")
	ErrInvalidWindingOrder  = errors.New("The vertex list does not contain a valid polygon.")
)
