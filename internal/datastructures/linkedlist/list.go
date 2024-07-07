// Package linkedlist implements the Linked List Data Structure.
package linkedlist

import "errors"

// LinkedList interface defines the operations for a singly linked list.
type LinkedList interface {
	ClearList()
	DeleteAtPosition(position int) (int, bool, error)
	DeleteHead() (int, bool)
	DeleteTail() (int, bool)
	IsEmpty() bool
	IsPresent(data int) bool
	InsertAtPosition(data int, position int)
	InsertHead(data int)
	InsertTail(data int)
	SetHeadIfEmpty(newNode *ListNode) bool
	GetHead() *ListNode
	GetTail() *ListNode
	Size() int
}

// ErrInvalidPosition is an error indicating an invalid position in the linked list.
var ErrInvalidPosition = errors.New("invalid position")
