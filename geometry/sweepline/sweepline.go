package sweepline

import (
	"container/list"

	"github.com/mindera-gaming/go-math/vector2"
)

// Using a Sweep-Line Algorithm for segment intersection, based on the pseudo-code found here:
// Text: https://web.archive.org/web/20141211224415/http://www.lems.brown.edu/~wq/projects/cs252.html
// Video: https://youtu.be/I9EsN2DTnN8

// FindIntersections finds and returns the intersections
func FindIntersections(inputData []Segment) (intersections []vector2.Vector2) {
	// initialising the event queue
	eventQueue := list.New()
	for _, segment := range inputData {
		insertEvent(eventQueue, NewEventSingleSegment(segment.First(), segment, start))
		insertEvent(eventQueue, NewEventSingleSegment(segment.Second(), segment, end))
	}

	// initialising the sweep-line status
	status := list.New()

	// runs the queue until it is empty
	for eventQueue.Len() != 0 {
		e := eventPoll(eventQueue).Value.(event)
		l := e.Value

		// handling an event
		switch e.Type {
		case start:
			for _, s := range e.Segments {
				recalculate(status, l)
				insertStatusSegment(status, s)

				lower := lowerStatusSegment(status, s)
				higher := higherStatusSegment(status, s)

				if lower != nil {
					reportIntersection(eventQueue, lower.Value.(Segment), s, l)
				}
				if higher != nil {
					reportIntersection(eventQueue, higher.Value.(Segment), s, l)
				}
				if lower != nil && higher != nil {
					removeFuture(eventQueue, lower.Value.(Segment), higher.Value.(Segment))
				}
			}
		case end:
			for _, s := range e.Segments {
				lower := lowerStatusSegment(status, s)
				higher := higherStatusSegment(status, s)

				if lower != nil && higher != nil {
					reportIntersection(eventQueue, lower.Value.(Segment), higher.Value.(Segment), l)
				}
				removeStatusSegment(status, s)
			}
		case intersection:
			s1 := e.Segments[0]
			s2 := e.Segments[1]

			swap(status, s1, s2)

			if s1.Value < s2.Value {
				s1Higher := higherStatusSegment(status, s1)
				s2Lower := lowerStatusSegment(status, s2)

				if s1Higher != nil {
					reportIntersection(eventQueue, s1Higher.Value.(Segment), s1, l)
					removeFuture(eventQueue, s1Higher.Value.(Segment), s2)
				}
				if s2Lower != nil {
					reportIntersection(eventQueue, s2Lower.Value.(Segment), s2, l)
					removeFuture(eventQueue, s2Lower.Value.(Segment), s1)
				}
			} else {
				s2Higher := higherStatusSegment(status, s2)
				s1Lower := lowerStatusSegment(status, s1)

				if s2Higher != nil {
					reportIntersection(eventQueue, s2Higher.Value.(Segment), s2, l)
					removeFuture(eventQueue, s2Higher.Value.(Segment), s1)
				}
				if s1Lower != nil {
					reportIntersection(eventQueue, s1Lower.Value.(Segment), s1, l)
					removeFuture(eventQueue, s1Lower.Value.(Segment), s2)
				}
			}

			// adds the new intersection
			intersections = append(intersections, e.Point)
		}
	}
	return
}

// reportIntersection reports an intersection point and its involved segments
func reportIntersection(queue *list.List, s1, s2 Segment, l float64) bool {
	x1 := s1.First().X
	y1 := s1.First().Y

	x2 := s1.Second().X
	y2 := s1.Second().Y

	x3 := s2.First().X
	y3 := s2.First().Y

	x4 := s2.Second().X
	y4 := s2.Second().Y

	r := (x2-x1)*(y4-y3) - (y2-y1)*(x4-x3)
	if r != 0 {
		t := ((x3-x1)*(y4-y3) - (y3-y1)*(x4-x3)) / r
		u := ((x3-x1)*(y2-y1) - (y3-y1)*(x2-x1)) / r

		if t >= 0 && t <= 1 && u >= 0 && u <= 1 {
			xC := x1 + t*(x2-x1)
			yC := y1 + t*(y2-y1)

			if xC > l {
				insertEvent(
					queue,
					NewEvent(
						vector2.Vector2{X: xC, Y: yC},
						[]Segment{s1, s2},
						intersection,
					))
				return true
			}
		}
	}
	return false
}

// removeFuture removes future segments from the queue
func removeFuture(queue *list.List, s1, s2 Segment) bool {
	for e := queue.Front(); e != nil; e = e.Next() {
		event := e.Value.(event)
		if event.Type == 2 {
			if (event.Segments[0] == s1 && event.Segments[1] == s2) || (event.Segments[0] == s2 && event.Segments[1] == s1) {
				queue.Remove(e)
				return true
			}
		}
	}
	return false
}

// swap switches two segments of the status list
func swap(status *list.List, s1, s2 Segment) {
	removeStatusSegment(status, s1)
	removeStatusSegment(status, s2)

	tempValue := s1.Value
	s1.Value = s2.Value
	s2.Value = tempValue

	insertStatusSegment(status, s1)
	insertStatusSegment(status, s2)
}

// recalculate recalculates the value of all segments in the status list
func recalculate(status *list.List, line float64) {
	for e := status.Front(); e != nil; e = e.Next() {
		value := e.Value.(Segment)
		value.CalculateValue(line)
		e.Value = value
	}
}
