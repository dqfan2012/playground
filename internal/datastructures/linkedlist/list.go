package linkedlist

// List interface defines the operations for a singly linked list.
type LinkedList interface {
	ClearList()
	DeleteAtPosition(int) (int, bool, error)
	DeleteHead() (int, bool)
	DeleteTail() (int, bool)
	IsEmpty() bool
	IsPresent(int) bool
	InsertAtPosition(int, int)
	InsertHead(int)
	InsertTail(int)
	SetHeadIfEmpty(*ListNode) bool
	Size() int
}
