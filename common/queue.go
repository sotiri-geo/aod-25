package common

import "errors"

var (
	ErrDequeueEmptyQueue = errors.New("cannot dequeue an empty queue")
	ErrPeakEmptyQueue    = errors.New("cannot peak an empty queue")
)

type Queue[T any] struct {
	queue []T
	head  int
}

func NewQueue[T any](queue []T) *Queue[T] {
	return &Queue[T]{queue: queue}
}

func (q *Queue[T]) Len() int {
	return len(q.queue) - q.head
}

func (q *Queue[T]) Enqueue(element T) {
	q.queue = append(q.queue, element)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if len(q.queue) <= q.head {
		return zero, ErrDequeueEmptyQueue
	}
	// Zeroing is to drop references for GC when T contains references
	top := q.queue[q.head]
	q.queue[q.head] = zero
	// increment pointer for future references
	q.head++

	// When head is large, copy tail to new slice
	// avoids frequent GC from running
	// Compaction occurs as q.head * 2 >= len(queue)
	// i.e. compaction occurs once at least half the backing slice is consumed
	if 2*q.head >= len(q.queue) {
		q.queue = q.queue[q.head:]
		// Reset head back to 0
		q.head = 0
	}

	return top, nil
}

func (q *Queue[T]) IsEmpty() bool {
	// In the event head grows beyond the length of queue
	return q.Len() <= 0
}

func (q *Queue[T]) Peek() (T, error) {
	var top T
	if q.IsEmpty() {
		return top, ErrPeakEmptyQueue
	}
	top = q.queue[q.head]

	return top, nil
}
