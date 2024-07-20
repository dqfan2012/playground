// Package set implements the set data structure.
package set

import (
	"sync"
)

// Set represents the set data structure.
// Mutex ensures the implementation of Set is thread-safe.
type Set[T comparable] struct {
	// items is a map where the key is the value we want to store in the set.
	// Using an empty struct as the value type minimizes memory usage.
	items map[T]struct{}
	mu    sync.Mutex
}

// New creates a new set.
func New[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]struct{}),
	}
}

// Add adds a value to the set if the value isn't present.
func (s *Set[T]) Add(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.items[value]

	if !exists {
		s.items[value] = struct{}{}
	}
}

// Remove removes a value from the set if it is present.
func (s *Set[T]) Remove(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.items[value]

	if exists {
		delete(s.items, value)
	}
}

// Contains checks if the set contains the given value, returning true if it does and false otherwise.
func (s *Set[T]) Contains(value T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.items[value]

	return exists
}

// Union returns a new set that contains all unique elements from this set and the provided set.
func (s *Set[T]) Union(setB *Set[T]) *Set[T] {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Create a new set
	newSet := New[T]()

	// Loop through the items in the first set and add them to the new set.
	for k, _ := range s.items {
		newSet.Add(k)
	}

	// Loop through the items in the second set and add them to the new set.
	for k, _ := range setB.items {
		newSet.Add(k)
	}

	// return the new set.
	return newSet
}

// Intersection returns a new set that contains only the elements present in the current set and
// the provided set.
func (s *Set[T]) Intersection(setB *Set[T]) *Set[T] {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Create a new set
	newSet := New[T]()

	for k, _ := range s.items {
		if _, exists := setB.items[k]; exists {
			newSet.Add(k)
		}
	}

	// return the new set.
	return newSet
}

// Difference returns a new set that contains elements unique to the current set and the provided set.
// Items present in both sets are not included in the newly returned set.
func (s *Set[T]) Difference(setB *Set[T]) *Set[T] {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Create a new set
	newSet := New[T]()

	// Add the unique items from the current set to the new set.
	for k, _ := range s.items {
		if _, exists := setB.items[k]; !exists {
			newSet.Add(k)
		}
	}

	// Add the unique items in setB to the new set.
	for k, _ := range setB.items {
		if _, exists := s.items[k]; !exists {
			newSet.Add(k)
		}
	}

	// return the new set.
	return newSet
}

// IsSubset checks if the provided set is a subset of the current set.
// Returns true if all the items in the provided set are present in the current set,
// otherwise returns false.
func (s *Set[T]) IsSubset(setB *Set[T]) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Add the unique items in setB to the new set.
	for k, _ := range setB.items {
		if _, exists := s.items[k]; !exists {
			return false
		}
	}

	// setB is a subset of the current set.
	return true
}

// Clear clears all the items from the set.
func (s *Set[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items = make(map[T]struct{})
}

// IsEmpty returns true if the set is empty. Otherwise, false.
func (s *Set[T]) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.items) == 0
}

// Len returns number of items in the set.
func (s *Set[T]) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.items)
}
