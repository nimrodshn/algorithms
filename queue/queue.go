package queue

import "errors"

var (
	// ErrQueueUnderflow is an error indicating the stack is empty.
	ErrQueueUnderflow = errors.New("Queue Underflow")

	// ErrQueueOverflow is an error indicating the stack is full.
	ErrQueueOverflow = errors.New("Queue Overflow")
)

// Queue is an implementation of a queue.
type Queue struct {
	length int
	array  []int
	top    int
}

func New(n int) *Queue {
	return &Queue{
		length: n,
		array:  make([]int, n),
		top:    -1,
	}
}

// Insert inserts a new element to the queue.
func (q *Queue) Insert(x int) error {
	if q.full() {
		return ErrQueueOverflow
	}
	k := q.array[0]
	for j := 1; j <= q.top+1; j++ {
		m := q.array[j]
		q.array[j] = k
		k = m
	}
	q.array[0] = x
	q.top++
	return nil
}

// Remove removes an element from the queue.
func (q *Queue) Remove() (int, error) {
	if q.empty() {
		return -1, ErrQueueUnderflow
	}
	res := q.array[q.top]
	q.top--
	return res, nil
}

// Size returns the size of the queue.
func (q *Queue) Size() int {
	return q.top + 1
}

func (q *Queue) empty() bool {
	return q.top == -1
}

func (q *Queue) full() bool {
	return q.top == (q.length - 1)
}
