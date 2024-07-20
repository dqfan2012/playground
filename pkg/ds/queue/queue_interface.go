// Package queue implements the queue data structure.
package queue

import "errors"

// ErrEmptyQueue is an error indicating that the queue is empty when attempting to
// dequeue an item from the queue.
var ErrEmptyQueue = errors.New("empty queue")

// Queueer defines the operations for a queue.
type Queuer[T comparable] interface {
	Enqueue(T)
	Dequeue() (*T, error)
	IsEmpty() bool
	Len() int
}
