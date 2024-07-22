// Package hashmap implements the HashMap data structure.
package hashmap

import "sync"

// MapNode implements a node in the HashMap.
type MapNode[T comparable] struct {
	key   string
	value T
	next  *MapNode[T]
}

// HashMap implements a thread-safe HashMap.
type HashMap[T comparable] struct {
	buckets []*MapNode[T]
	len     int
	mu      sync.Mutex
}

// NewHashMap creates and returns a new HashMap with a specified number of buckets.
func NewHashMap[T comparable](len int) *HashMap[T] {
	return &HashMap[T]{
		buckets: make([]*MapNode[T], len),
		len:     len,
	}
}

// hash converts a key string into a hash, which is an index for the buckets array.
func (h *HashMap[T]) hash(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}

	return hash % h.len
}

// Insert adds a key-value pair into the HashMap. If the key already exists, its value is updated.
func (h *HashMap[T]) Insert(key string, value T) {
	h.mu.Lock()
	defer h.mu.Unlock()

	index := h.hash(key)
	newNode := &MapNode[T]{
		key:   key,
		value: value,
	}

	if h.buckets[index] == nil {
		h.buckets[index] = newNode
	} else {
		current := h.buckets[index]
		for current.next != nil && current.key != key {
			current = current.next
		}
		if current.key == key {
			current.value = value
		} else {
			current.next = newNode
		}
	}
}

// Get retrieves the value associated with a given key from the HashMap.
// It returns the value and a boolean indicating whether the key was found.
func (h *HashMap[T]) Get(key string) (*T, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()

	index := h.hash(key)
	current := h.buckets[index]

	for current != nil {
		if current.key == key {
			return &current.value, true
		}
		current = current.next
	}

	return nil, false
}

// Delete removes the key-value pair associated with a given key from the HashMap.
func (h *HashMap[T]) Delete(key string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	index := h.hash(key)
	current := h.buckets[index]
	var prev *MapNode[T]

	for current != nil {
		if current.key == key {
			if prev == nil {
				h.buckets[index] = current.next
			} else {
				prev.next = current.next
			}
			return
		}
		prev = current
		current = current.next
	}
}
