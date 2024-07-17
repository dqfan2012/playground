// Package queue implements the queue data structure.
package queue

import "sync"

// Item is an item that gets added to the queue.
type Item[T comparable] struct {
	Value    T
	Priority int
}

// PriorityQueue represents the queue data structure containing items with set priority.
// sync.Mutex is used to ensure safe concurrent access to the queue.
type PriorityQueue[T comparable] struct {
	items []*Item[T]
	mu    sync.Mutex
}

// NewPriorityQueue creates a new priority queue.
func NewPriorityQueue[T comparable]() *PriorityQueue[T] {
	return &PriorityQueue[T]{}
}

// Enqueue adds an item with a specific priority to the correct place in the queue.
func (pq *PriorityQueue[T]) Enqueue(value T, priority int) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	item := &Item[T]{Value: value, Priority: priority}

	// Find the correct position to insert the item based on priority.
	position := 0
	for ; position < len(pq.items); position++ {
		// Break the loop when finding the correct position.
		if pq.items[position].Priority > priority {
			break
		}
	}

	// If the position is set to the end
	if position == len(pq.items) {
		// append the item to the end of the queue.
		pq.items = append(pq.items, item)
	} else {
		// pq.items[:position+1] - Create a new slice from the beginning of pq.items
		// up to, but not including, position+1. Essentially, it includes elements
		// at indices 0 to position.
		// pq.items[position:]... - Takes the slice of all elements from position to
		// the end of pq.items and unpacks it so that each element is appended
		// individually.
		// append(pq.items[:position+1], pq.items[position:]...) - This combines the
		// two slices. The first part includes elements from the beginning to position
		// inclusive, and the second part includes elements from position to the end.
		// It creates a new slice where elements at and after position are moved one
		// position forward, leaving an empty space at position for the new item.
		pq.items = append(pq.items[:position+1], pq.items[position:]...)
		// Insert the new item at the position
		pq.items[position] = item
	}
}

// Dequeue removes and returns the item from the front of the queue.
// The item at the front of the queue should always have highest priority.
// It returns an error if the queue is empty.
func (pq *PriorityQueue[T]) Dequeue() (*Item[T], error) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	if len(pq.items) == 0 {
		return nil, ErrEmptyQueue
	}

	item := pq.items[0]
	pq.items = pq.items[1:]

	return item, nil
}

// IsEmpty returns true if the queue is empty, otherwise false.
func (pq *PriorityQueue[T]) IsEmpty() bool {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	return len(pq.items) == 0
}

// Len returns the number of items in the queue.
func (pq *PriorityQueue[T]) Len() int {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	return len(pq.items)
}
