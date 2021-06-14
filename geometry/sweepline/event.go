package sweepline

import vector "github.com/mindera-gaming/go-math/vector2"

// eventType defines an enumerator that represents the type of the event
type eventType byte

const (
	start eventType = iota
	end
	intersection
)

// event defines an event point
type event struct {
	Point    vector.Vector2
	Segments []Segment
	Value    float64
	Type     eventType
}

// NewEventSingleSegment returns a new event with a single segment
func NewEventSingleSegment(p vector.Vector2, s Segment, t eventType) (e event) {
	e.Point = p
	e.Segments = append(e.Segments, s)
	e.Value = p.X
	e.Type = t
	return
}

// NewEvent returns a new event with a set of segments
func NewEvent(p vector.Vector2, s []Segment, t eventType) event {
	return event{
		Point:    p,
		Segments: s,
		Value:    p.X,
		Type:     t,
	}
}

// AddSegment adds a new segment to the event
func (e *event) AddSegment(s Segment) {
	e.Segments = append(e.Segments, s)
}
