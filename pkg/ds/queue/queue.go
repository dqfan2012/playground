// Package queue implements the queue data structure.
package queue

import (
	"sync"
)

// Queue represents the queue data structure.
// sync.Mutex is used to ensure safe concurrent access to the queue.
type Queue[T comparable] struct {
	items []T
	mu    sync.Mutex
}

// New creates a new queue.
func New[T comparable]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue adds an item to the end of the queue.
func (q *Queue[T]) Enqueue(value T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.items = append(q.items, value)
}

// Dequeue removes and returns the item from the front of the queue.
// It returns an error if the queue is empty.
func (q *Queue[T]) Dequeue() (*T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return nil, ErrEmptyQueue
	}

	// Remove the first item from the front of the queue, since
	// Queues are FIFO data structures
	item := q.items[0]
	q.items = q.items[1:]

	return &item, nil
}

// IsEmpty returns true if the queue is empty, otherwise false.
func (q *Queue[T]) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.items) == 0
}

// Len returns the number of items in the queue.
func (q *Queue[T]) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.items)
}
