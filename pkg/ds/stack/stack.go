// Package stack implements the stack data structure.
package stack

import (
	"sync"
)

// Stack represents the stack data structure. Stacks are LIFO data structures.
type Stack[T comparable] struct {
	items []T
	mu    sync.Mutex
}

// New creates a new stack.
func New[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an item to the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items = append(s.items, value)
}

// Pop removes the item from the top of the stack and returns it.
func (s *Stack[T]) Pop() (*T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.items) == 0 {
		return nil, ErrEmptyStack
	}

	// Remove the item from the top of the stack
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	// Return the item that was at the top of the stack.
	return &item, nil
}

// Peek returns the item at the top of the stack without removing it from the stack.
func (s *Stack[T]) Peek() (*T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.items) == 0 {
		return nil, ErrEmptyStack
	}

	// Retrieve the item from the top of the stack without removing it
	item := s.items[len(s.items)-1]

	// Return the item that's on the top of the stack
	return &item, nil
}

// IsEmpty returns true if the stack is empty. Otherwise returns false.
func (s *Stack[T]) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.items) == 0
}

// Len returns the length of the stack.
func (s *Stack[T]) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.items)
}
