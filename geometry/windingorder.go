package geometry

// WindingOrder defines an enumerator representing the winding order of the vertices.
type WindingOrder int

const (
	Invalid WindingOrder = iota - 1
	Clockwise
	CounterClockwise
)
