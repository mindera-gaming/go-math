package sweepline

import (
	"github.com/mindera-gaming/go-data-structure/comparator"
	"github.com/mindera-gaming/go-data-structure/navigableset"
	"github.com/mindera-gaming/go-data-structure/queue"
	"github.com/mindera-gaming/go-math/vector2"
)

// Using a Sweep-Line Algorithm for segment intersection, based on the pseudo-code found here:
// Text: https://web.archive.org/web/20141211224415/http://www.lems.brown.edu/~wq/projects/cs252.html
// Video: https://youtu.be/I9EsN2DTnN8

var queueComparator = func(v1, v2 interface{}) comparator.Result {
	value1 := v1.(event).Value
	value2 := v2.(event).Value
	if value1 < value2 {
		return comparator.Less
	}
	if value1 > value2 {
		return comparator.Greater
	}
	return comparator.Equal
}

var navigableComparator = func(v1, v2 interface{}) comparator.Result {
	value1 := v1.(Segment).Value
	value2 := v2.(Segment).Value
	if value1 > value2 {
		return comparator.Less
	}
	if value1 < value2 {
		return comparator.Greater
	}
	return comparator.Equal
}

// FindIntersections finds and returns the intersections
func FindIntersections(inputData []Segment) (intersections []vector2.Vector2) {
	// initialising the event queue

	eventQueue := queue.New(queueComparator)
	for _, segment := range inputData {
		eventQueue.Add(NewEventSingleSegment(segment.First(), segment, start))
		eventQueue.Add(NewEventSingleSegment(segment.Second(), segment, end))
	}

	// initialising the sweep-line status
	status, _ := navigableset.New(navigableComparator)

	// runs the queue until it is empty
	for eventQueue.Len() != 0 {
		e := eventQueue.Poll().(event)
		l := e.Value

		// handling an event
		switch e.Type {
		case start:
			for _, s := range e.Segments {
				recalculate(&status, l)
				status.Add(s)

				lower := status.Lower(s)
				higher := status.Higher(s)

				if lower != nil {
					reportIntersection(&eventQueue, lower.(Segment), s, l)
				}
				if higher != nil {
					reportIntersection(&eventQueue, higher.(Segment), s, l)
				}
				if lower != nil && higher != nil {
					removeFuture(&eventQueue, lower.(Segment), higher.(Segment))
				}
			}
		case end:
			for _, s := range e.Segments {
				lower := status.Lower(s)
				higher := status.Higher(s)

				if lower != nil && higher != nil {
					reportIntersection(&eventQueue, lower.(Segment), higher.(Segment), l)
				}
				status.Remove(s)
			}
		case intersection:
			s1 := e.Segments[0]
			s2 := e.Segments[1]

			swap(&status, s1, s2)

			if s1.Value < s2.Value {
				s1Higher := status.Higher(s1)
				s2Lower := status.Lower(s2)

				if s1Higher != nil {
					reportIntersection(&eventQueue, s1Higher.(Segment), s1, l)
					removeFuture(&eventQueue, s1Higher.(Segment), s2)
				}
				if s2Lower != nil {
					reportIntersection(&eventQueue, s2Lower.(Segment), s2, l)
					removeFuture(&eventQueue, s2Lower.(Segment), s1)
				}
			} else {
				s2Higher := status.Higher(s2)
				s1Lower := status.Lower(s1)

				if s2Higher != nil {
					reportIntersection(&eventQueue, s2Higher.(Segment), s2, l)
					removeFuture(&eventQueue, s2Higher.(Segment), s1)
				}
				if s1Lower != nil {
					reportIntersection(&eventQueue, s1Lower.(Segment), s1, l)
					removeFuture(&eventQueue, s1Lower.(Segment), s2)
				}
			}

			// adds the new intersection
			intersections = append(intersections, e.Point)
		}
	}
	return
}

// reportIntersection reports an intersection point and its involved segments
func reportIntersection(queue *queue.Queue, s1, s2 Segment, l float64) bool {
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
				queue.Add(NewEvent(
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
func removeFuture(queue *queue.Queue, s1, s2 Segment) bool {
	for e := queue.List.Front(); e != nil; e = e.Next() {
		event := e.Value.(event)
		if event.Type == intersection {
			if (event.Segments[0] == s1 && event.Segments[1] == s2) || (event.Segments[0] == s2 && event.Segments[1] == s1) {
				queue.Remove(e)
				return true
			}
		}
	}
	return false
}

// swap switches two segments of the status list
func swap(status *navigableset.NavigableSet, s1, s2 Segment) {
	status.Remove(s1)
	status.Remove(s2)

	tempValue := s1.Value
	s1.Value = s2.Value
	s2.Value = tempValue

	status.Add(s1)
	status.Add(s2)
}

// recalculate recalculates the value of all segments in the status list
func recalculate(status *navigableset.NavigableSet, line float64) {
	for e := status.List.Front(); e != nil; e = e.Next() {
		value := e.Value.(Segment)
		value.CalculateValue(line)
		e.Value = value
	}
}
