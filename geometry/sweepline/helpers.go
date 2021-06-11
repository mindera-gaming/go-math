package sweepline

import "container/list"

// insertEvent inserts a new event to the queue
func insertEvent(list *list.List, event Event) {
	for e := list.Front(); e != nil; e = e.Next() {
		value := e.Value.(Event).Value
		if value == event.Value {
			list.InsertBefore(event, e)
			return
		}
		if value > event.Value {
			list.InsertBefore(event, e)
			return
		}
	}
	list.PushBack(event)
}

// eventPoll returns and removes the head of the queue
func eventPoll(list *list.List) (element *list.Element) {
	element = list.Front()
	list.Remove(element)
	return
}

// insertStatusSegment inserts a new segment to the list in descending order
func insertStatusSegment(list *list.List, segment Segment) {
	for e := list.Front(); e != nil; e = e.Next() {
		value := e.Value.(Segment).Value
		if value == segment.Value {
			return
		}
		if value < segment.Value {
			list.InsertBefore(segment, e)
			return
		}
	}
	list.PushBack(segment)
}

// removeStatusSegment removes a segment from the list by its value
func removeStatusSegment(list *list.List, segment Segment) {
	for e := list.Front(); e != nil; e = e.Next() {
		if e.Value.(Segment).Equals(segment) {
			list.Remove(e)
			return
		}
	}
}

// lowerStatusSegment returns the previous element to the element with the indicated value
func lowerStatusSegment(list *list.List, segment Segment) *list.Element {
	for e := list.Front(); e != nil; e = e.Next() {
		if e.Value.(Segment).Equals(segment) {
			return e.Prev()
		}
	}
	return nil
}

// higherStatusSegment returns the next element to the element with the indicated value
func higherStatusSegment(list *list.List, segment Segment) *list.Element {
	for e := list.Front(); e != nil; e = e.Next() {
		if e.Value.(Segment).Equals(segment) {
			return e.Next()
		}
	}
	return nil
}
