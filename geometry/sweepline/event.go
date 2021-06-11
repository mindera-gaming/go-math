package sweepline

import vector "github.com/mindera-gaming/go-math/vector2"

// Event defines an event point
type Event struct {
	Point    vector.Vector2
	Segments []Segment
	Value    float64
	Type     int
}

// NewEventSingleSegment returns a new event with a single segment
func NewEventSingleSegment(p vector.Vector2, s Segment, t int) (e Event) {
	e.Point = p
	e.Segments = append(e.Segments, s)
	e.Value = p.X
	e.Type = t
	return
}

// NewEvent returns a new event with a set of segments
func NewEvent(p vector.Vector2, s []Segment, t int) Event {
	return Event{
		Point:    p,
		Segments: s,
		Value:    p.X,
		Type:     t,
	}
}

// AddSegment adds a new segment to the event
func (e *Event) AddSegment(s Segment) {
	e.Segments = append(e.Segments, s)
}
