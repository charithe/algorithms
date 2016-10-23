package ds

import "fmt"

var (
	ErrQueueFull  = fmt.Errorf("Queue is full")
	ErrQueueEmpty = fmt.Errorf("Queue is empty")
)

type CircularQueue struct {
	capacity int
	size     int
	head     int
	tail     int
	data     []interface{}
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		capacity: capacity,
		size:     0,
		head:     0,
		tail:     0,
		data:     make([]interface{}, capacity),
	}
}

func (q *CircularQueue) Enqueue(item interface{}) error {
	if q.size >= q.capacity {
		return ErrQueueFull
	}

	q.data[q.tail] = item
	q.tail = (q.tail + 1) % q.capacity
	q.size += 1
	return nil
}

func (q *CircularQueue) Dequeue() (interface{}, error) {
	if q.size == 0 {
		return nil, ErrQueueEmpty
	}

	item := q.data[q.head]
	q.head = (q.head + 1) % q.capacity
	q.size -= 1
	return item, nil
}
