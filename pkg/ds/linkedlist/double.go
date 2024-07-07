// Package linkedlist implements the Linked List Data Structure.
package linkedlist

import "fmt"

// Double represents a doubly linked list.
type Double[T comparable] struct {
	Head, Tail *ListNode[T]
	len        int
}

// NewDoubleEmpty creates a new linked list without a head or tail.
func NewDoubleEmpty[T comparable]() *Double[T] {
	return &Double[T]{}
}

// NewDoubleWithHead creates a new linked list with the head element set.
func NewDoubleWithHead[T comparable](data T) *Double[T] {
	head := NewListNodeWithData[T](data)

	return &Double[T]{
		Head: head,
		Tail: head,
		len:  1,
	}
}

// ClearList clears the Linked List.
func (list *Double[T]) ClearList() {
	list.Head = nil
	list.Tail = nil
	list.len = 0
}

// DeleteAtPosition deletes the node at the specified position.
func (list *Double[T]) DeleteAtPosition(position int) (any, bool, error) {
	if list.isInvalidPosition(position) {
		return nil, false, fmt.Errorf("%w: %d", ErrInvalidPosition, position)
	}

	if list.IsEmpty() {
		return 0, false, nil
	}

	var value any
	if position == 0 {
		value = list.Head.Data
		list.Head = list.Head.Next

		if list.Head != nil {
			list.Head.Prev = nil
		} else {
			list.Tail = nil
		}

		list.len--

		return value, true, nil
	}

	current := list.Head
	for i := 0; i < position-1; i++ {
		current = current.Next
	}

	value = current.Next.Data
	current.Next = current.Next.Next

	if current.Next != nil {
		current.Next.Prev = current
	} else {
		list.Tail = current
	}

	list.len--

	return value, true, nil
}

// DeleteHead deletes the head node.
func (list *Double[T]) DeleteHead() (any, bool) {
	if list.IsEmpty() {
		return 0, false
	}

	value := list.Head.Data
	list.Head = list.Head.Next

	if list.Head != nil {
		list.Head.Prev = nil
	} else {
		list.Tail = nil
	}

	list.len--

	return value, true
}

// DeleteTail deletes the tail node.
func (list *Double[T]) DeleteTail() (any, bool) {
	if list.IsEmpty() {
		return 0, false
	}

	if list.len == 1 {
		value := list.Head.Data
		list.ClearList()

		return value, true
	}

	value := list.Tail.Data
	current := list.Tail.Prev
	current.Next = nil
	list.Tail = current
	list.len--

	return value, true
}

// IsEmpty returns whether the linked list is empty.
func (list *Double[T]) IsEmpty() bool {
	return list.Len() == 0
}

// isInvalidPosition checks to see if a list position is invalid.
func (list *Double[T]) isInvalidPosition(position int) bool {
	return position < 0 || position > list.Len()
}

// InsertAtPosition inserts a node at the specified position.
func (list *Double[T]) InsertAtPosition(position int, data T) (bool, error) {
	if list.isInvalidPosition(position) {
		return false, fmt.Errorf("%w: %d", ErrInvalidPosition, position)
	}

	newNode := NewListNodeWithData(data)
	if list.SetHeadIfEmpty(newNode) {
		return true, nil
	}

	current := list.Head
	for i := 0; i < position-1; i++ {
		current = current.Next
	}

	newNode.Next = current.Next
	newNode.Prev = current
	current.Next = newNode
	list.len++

	return true, nil
}

// InsertHead inserts a node at the head of the list.
func (list *Double[T]) InsertHead(data T) bool {
	newNode := NewListNodeWithData(data)
	if list.SetHeadIfEmpty(newNode) {
		return true
	}

	newNode.Next = list.Head
	list.Head.Prev = newNode
	list.Head = newNode
	list.len++

	return true
}

// InsertTail inserts a node at the tail of the list.
func (list *Double[T]) InsertTail(data T) bool {
	newNode := NewListNodeWithData(data)
	if list.SetHeadIfEmpty(newNode) {
		return true
	}

	list.Tail.Next = newNode
	newNode.Prev = list.Tail
	list.Tail = newNode
	list.len++

	return true
}

// IsValuePresent checks if a value is present in the list.
func (list *Double[T]) IsValuePresent(data T) bool {
	if !list.IsEmpty() {
		temp := list.Head

		for temp != nil {
			if temp.Data == data {
				return true
			}

			temp = temp.Next
		}
	}

	return false
}

// SetHeadIfEmpty sets the head if the list is empty or position is invalid.
func (list *Double[T]) SetHeadIfEmpty(newHead *ListNode[T]) bool {
	if list.IsEmpty() {
		list.Head = newHead
		list.Tail = newHead
		list.len++

		return true
	}

	return false
}

// GetHead retrieves the head node of the linked list.
func (list *Double[T]) GetHead() *ListNode[T] {
	return list.Head
}

// GetTail retrieves the tail node of the linked list.
func (list *Double[T]) GetTail() *ListNode[T] {
	return list.Tail
}

// Len retrieves the length of the linked list.
func (list *Double[T]) Len() int {
	return list.len
}
