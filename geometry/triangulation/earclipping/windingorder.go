package earclipping

// windingOrder defines an enumerator representing the winding order of the vertices.
type windingOrder int

const (
	invalid windingOrder = iota - 1
	clockwise
	counterClockwise
)
