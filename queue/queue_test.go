package queue

import "testing"

func TestQueueStruct(t *testing.T) {
	q := new(queue)
	nitems := len(q.items)
	witems := 0
	index := q.index
	windex := 0
	if nitems != witems {
		t.Errorf("len(queue.items) == %d, want %d", nitems, witems)
	}
	if index != windex {
		t.Errorf("queue.index == %d, want %d", index, windex)
	}
}

func TestEnqueue(t *testing.T) {
	q := new(queue)
	var enqueuetests = []struct {
		in int
	}{
		{0},
	}
	for _, tt := range enqueuetests {
		windex := q.index
		q.enqueue(tt.in)
		newIndex := q.index
		if newIndex != windex {
			t.Errorf("q.index == %d, want %d", newIndex, windex)
		}
	}
}

func TestDequeue(t *testing.T) {
	q := new(queue)
	_, err := q.dequeue()
	if err == nil {
		t.Error("q.dequeue on empty queue, want error")
	}
	var dequeuetests = []struct {
		in, out int
	}{
		{0, 0},
	}
	for _, tt := range dequeuetests {
		q.enqueue(tt.in)
	}
	for _, tt := range dequeuetests {
		got, err := q.dequeue()
		if err != nil {
			t.Errorf("q.dequeue: %s", err)
		}
		if got != tt.out {
			t.Errorf("q.dequeue() == %d, want %d", got, tt.out)
		}
	}
}
