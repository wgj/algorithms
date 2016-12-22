package queue

import "errors"

type queue struct {
	items []int
	index int
	cap   int
}

func (q *queue) enqueue(item int) {
	if len(q.items) == q.index {
		q.grow()
	}
	q.items[q.index] = item
}

func (q *queue) dequeue() (int, error) {
	if len(q.items) == q.index {
		return 0, errors.New("empty queue")
	}
	item := q.items[q.index]
	q.index++
	return item, nil
}

func (q *queue) grow() {
	var nitems int
	if len(q.items) == 0 {
		nitems = 1
	} else {
		nitems = len(q.items) * 2
	}
	newItems := make([]int, nitems)
	if len(q.items) > 0 {
		copy(newItems, q.items)
	}
	q.items = newItems
}
