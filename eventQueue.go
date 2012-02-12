package baukasten

import (
	"container/list"
	"sync"
)

type Event interface{}

type EventQueue interface {
	Push(Event) Event
	Pull() Event
	Wait() Event
}

type ListEventQueue struct {
	list          *list.List
	lock          *sync.Cond
	limit         int  // bound on size
	discardNewest bool // default is to discard oldest when the limit is reached
}

// Create an empty linked-list based event queue.
// limit specifies the number of events to buffer before discarding (0 indicates no limit).
// discardNewest specifies whether older or newer events should discarded when the queue is full.
func NewEventQueue(limit int, discardNewest bool) *ListEventQueue {
	if limit < 0 {
		limit = 0
	}
	return &ListEventQueue{new(list.List), sync.NewCond(new(sync.Mutex)), limit, discardNewest}
}

// Get the value of discardNewest.
func (q *ListEventQueue) DiscardNewest() bool {
	return q.discardNewest
}

// Set the value of discardNewest, returning the old value.
func (q *ListEventQueue) SetDiscardNewest(newest bool) bool {
	q.lock.L.Lock()
	defer q.lock.L.Unlock()
	previous := q.discardNewest
	q.discardNewest = newest
	return previous
}

// Get the value of limit
func (q *ListEventQueue) Limit() int {
	q.lock.L.Lock()
	defer q.lock.L.Unlock()
	return q.limit
}

// Set the value of limit, returning the old limit value and any discarded events.
func (q *ListEventQueue) SetLimit(limit int) (oldLimit int, discarded []Event) {
	oldLimit = q.limit
	q.lock.L.Lock()
	defer q.lock.L.Unlock()
	q.limit = limit
	if limit <= 0 {
		return
	}
	discarded = q.drainTo(limit)
	return
}

func (q *ListEventQueue) pull(wait bool) Event {
	q.lock.L.Lock()
	defer q.lock.L.Unlock()
	lst := q.list
	if wait {
		if lst.Len() == 0 {
			return nil
		}
	} else {
		for lst.Len() == 0 {
			q.lock.Wait()
		}
	}
	return lst.Remove(lst.Front())
}

// Pulls the oldest event off the queue, returning that event, or nil if the queue is empty.
func (q *ListEventQueue) Pull() Event {
	return q.pull(false)
}

// Pulls the oldest event off the queue, returning that event.
// This method will block until an event is available.
func (q *ListEventQueue) Wait() Event {
	return q.pull(true)
}

// Push an event on the end of the queue, returning the event that was
// discarded to accommodate the push, if any, or nil if there was space.
func (q *ListEventQueue) Push(e Event) (discarded Event) {
	q.lock.L.Lock()
	defer q.lock.Signal()
	defer q.lock.L.Unlock()
	lst := q.list
	if q.limit > 0 && lst.Len() >= q.limit {
		switch {
		case lst.Len() > q.limit:
			panic("ListEventQueue has inexplicably overflowed")
		case q.discardNewest:
			discarded = lst.Remove(lst.Back())
		default:
			discarded = lst.Remove(lst.Front())
		}
	}
	q.list.PushBack(e)
	return
}

func (q *ListEventQueue) clear() *list.List {
	lst := q.list
	q.list = list.New()
	return lst
}

// Atomically empties the queue, releasing all queued events.
func (q *ListEventQueue) Clear() {
	q.lock.L.Lock()
	q.clear()
	q.lock.L.Unlock()
}

// Atomically empties the queue, returning a slice with the discarded events in normal order.
func (q *ListEventQueue) Drain() []Event {
	q.lock.L.Lock()
	lst := q.clear()
	q.lock.L.Unlock()
	length := lst.Len()
	result := make([]Event, length)
	for i := 0; i < length; i++ {
		result[i] = lst.Remove(lst.Front())
	}
	return result
}

func (q *ListEventQueue) drainTo(length int) []Event {
	lst := q.list
	currlen := lst.Len()
	if currlen <= length {
		return nil
	}
	size := currlen - length
	discarded := make([]Event, size)
	if q.discardNewest {
		for i := size - 1; i >= 0; i-- {
			discarded[i] = lst.Remove(lst.Back())
		}
	} else {
		for i := 0; i < size; i++ {
			discarded[i] = lst.Remove(lst.Front())
		}
	}
	return discarded
}

// Atomically drain down to the specified length, returning any discarded events in normal order.
func (q *ListEventQueue) DrainTo(length int) []Event {
	if length <= 0 {
		return q.Drain()
	}
	q.lock.L.Lock()
	defer q.lock.L.Unlock()
	return q.drainTo(length)
}

// Continually fill an EventQueue from a channel of events until the channel is closed.
// This does not spawn its own goroutines, and thus blocks unless invoked as a goroutine.
func ChanToQueue(q EventQueue, ch <-chan Event) {
	for event := range ch {
		q.Push(event) // do not intercept the return value
	}
}

// Continually send events from an EventQueue to a channel, until the channel is closed.
// This does not spawn its own goroutines, and thus blocks unless invoked as a goroutine.
func QueueToChan(q EventQueue, ch chan<- Event) {
	defer func() { recover() }() // don't panic on closed chan
	for {
		ch <- q.Wait()
	}
}
