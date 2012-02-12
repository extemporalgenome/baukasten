package baukasten

import "testing"

func TestQueueLimitOldest(t *testing.T) {
	testQueueLimit(t, false, 2, 0)
}

func TestQueueLimitNewest(t *testing.T) {
	testQueueLimit(t, true, 2, 1)
}

func testQueueLimit(t *testing.T, discardNewest bool, size, sentinel int) {
	q := NewEventQueue(size, discardNewest)
	var i int
	for i = 0; i < size; i++ {
		if discarded := q.Push(i); discarded != nil {
			t.Fatal("ListEventQueue unexpectedly dumped an item:", discarded)
		}
	}
	if discarded := q.Push(i); discarded != sentinel {
		t.Fatal("ListEventQueue dumped an unexpected item:", discarded)
	}
}

func sliceEqual(a, b []Event) bool {
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}
	return true
}

func testQueueDrainTo(t *testing.T, discardNewest bool, size int, match []Event) {
	q := NewEventQueue(size, discardNewest)
	var i int
	for i = 0; i < size; i++ {
		if discarded := q.Push(i); discarded != nil {
			t.Fatal("ListEventQueue unexpectedly dumped an item:", discarded)
		}
	}
	if discarded := q.DrainTo(size - len(match)); !sliceEqual(discarded, match) {
		t.Fatal("ListEventQueue did not drain as expected: expected", match, "got", discarded)
	}
}

func TestQueueDrainToOldest(t *testing.T) {
	testQueueDrainTo(t, false, 10, []Event{0, 1, 2})
}

func TestQueueDrainToNewest(t *testing.T) {
	testQueueDrainTo(t, true, 10, []Event{7, 8, 9})
}
