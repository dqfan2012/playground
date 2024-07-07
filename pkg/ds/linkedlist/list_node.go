// Package linkedlist implements the Linked List Data Structure.
package linkedlist

// ListNode represents a single element in a singly linked list.
// Each Node contains some data and a pointer to the next Node in the list.
type ListNode[T comparable] struct {
	Data       T            // The data stored in the node
	Prev, Next *ListNode[T] // Pointer to the previous and next node.
}

// NewEmptyListNode returns a new empty list node.
func NewEmptyListNode[T comparable]() *ListNode[T] {
	return &ListNode[T]{}
}

// NewListNodeWithData returns a new Node with its data value set.
func NewListNodeWithData[T comparable](data T) *ListNode[T] {
	return &ListNode[T]{
		Data: data,
		Prev: nil,
		Next: nil,
	}
}
