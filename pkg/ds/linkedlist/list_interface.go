// Package linkedlist implements the Linked List Data Structure.
package linkedlist

import "errors"

// LinkedList interface defines the operations for a linked list.
type LinkedList[T comparable] interface {
	DeleteAtPosition(position int) (any, bool, error)
	DeleteHead() (any, bool)
	DeleteTail() (any, bool)
	IsPresent(data any) bool
	InsertAtPosition(any, int) (bool, error)
	InsertHead(data any) bool
	InsertTail(data any) bool
	SetHeadIfEmpty(newNode *ListNode[T]) bool
	GetHead() *ListNode[T]
	GetTail() *ListNode[T]
}

// ListHelper interface defines additional operations for a linked list.
type ListHelper[T comparable] interface {
	ClearList()
	IsEmpty() bool
	Len() int
}

// ErrInvalidPosition is an error indicating an invalid position in the linked list.
var ErrInvalidPosition = errors.New("invalid position")
