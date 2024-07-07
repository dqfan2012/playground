// Package linkedlist implements the Linked List Data Structure.
package linkedlist

import "fmt"

// Single represents a singly linked list.
type Single[T comparable] struct {
	Head, Tail *ListNode[T]
	len        int
}

// NewSingleEmpty creates a new linked list without a head or tail.
func NewSingleEmpty[T comparable]() *Single[T] {
	return &Single[T]{}
}

// NewSingleWithHead creates a new linked list with the head element set.
func NewSingleWithHead[T comparable](data T) *Single[T] {
	head := NewListNodeWithData[T](data)

	return &Single[T]{
		Head: head,
		Tail: head,
		len:  1,
	}
}

// ClearList clears the Linked List.
func (list *Single[T]) ClearList() {
	list.Head = nil
	list.Tail = nil
	list.len = 0
}

// DeleteAtPosition deletes a list node at a specific position.
func (list *Single[T]) DeleteAtPosition(position int) (any, bool, error) {
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

		if list.Head == nil {
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

	if current.Next == nil {
		list.Tail = current
	}

	list.len--

	return value, true, nil
}

// DeleteHead deletes the head node.
func (list *Single[T]) DeleteHead() (any, bool) {
	if list.IsEmpty() {
		return 0, false
	}

	value := list.Head.Data
	list.Head = list.Head.Next
	list.len--

	return value, true
}

// DeleteTail should delete the tail node from the list.
func (list *Single[T]) DeleteTail() (any, bool) {
	if list.IsEmpty() {
		return 0, false
	}

	if list.len == 1 {
		value := list.Head.Data
		list.ClearList()

		return value, true
	}

	current := list.Head
	for current.Next.Next != nil {
		current = current.Next
	}

	value := current.Next.Data
	current.Next = nil
	list.len--

	return value, true
}

// IsEmpty returns whether the linked list is empty.
func (list *Single[T]) IsEmpty() bool {
	return list.Len() == 0
}

// isInvalidPosition checks to see if a list position is invalid.
func (list *Single[T]) isInvalidPosition(position int) bool {
	return position < 0 || position > list.Len()
}

// InsertAtPosition will insert a node into the list at the given position.
func (list *Single[T]) InsertAtPosition(position int, data T) (bool, error) {
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
	current.Next = newNode
	list.len++

	return true, nil
}

// InsertHead will insert a node at the beginning of the list.
func (list *Single[T]) InsertHead(data T) bool {
	newNode := NewListNodeWithData[T](data)
	if list.SetHeadIfEmpty(newNode) {
		return true
	}

	newNode.Next = list.Head
	list.Head = newNode
	list.len++

	return true
}

// InsertTail will insert a node at the end of the list.
func (list *Single[T]) InsertTail(data T) bool {
	newNode := NewListNodeWithData(data)
	if list.SetHeadIfEmpty(newNode) {
		return true
	}

	list.Tail.Next = newNode
	list.Tail = newNode
	list.len++

	return true
}

// IsValuePresent checks to see if the value is present in the list.
func (list *Single[T]) IsValuePresent(data T) bool {
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
func (list *Single[T]) SetHeadIfEmpty(newHead *ListNode[T]) bool {
	if list.IsEmpty() {
		list.Head = newHead
		list.Tail = newHead
		list.len++

		return true
	}

	return false
}

// GetHead retrieves the head node of the linked list.
func (list *Single[T]) GetHead() *ListNode[T] {
	return list.Head
}

// GetTail retrieves the tail node of the linked list.
func (list *Single[T]) GetTail() *ListNode[T] {
	return list.Tail
}

// Len retrieves the length of the linked list.
func (list *Single[T]) Len() int {
	return list.len
}
