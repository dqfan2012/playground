// Package linkedlist implements the Linked List Data Structure.
package linkedlist

import "fmt"

// SinglyLinkedList represents a singly linked list.
type SinglyLinkedList struct {
	Head   *ListNode
	Tail   *ListNode
	Length int
}

// NewEmptySinglyLinkedList creates an empty linked list.
func NewEmptySinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// NewSinglyLinkedListWithHead creates a linked list with an initial head node.
func NewSinglyLinkedListWithHead(data int) *SinglyLinkedList {
	newNode := NewListNodeWithData(data)

	return &SinglyLinkedList{
		head:   newNode,
		tail:   newNode,
		length: 1,
	}
}

// ClearList clears the linked list.
func (list *SinglyLinkedList) ClearList() {
	list.head = nil
	list.tail = nil
	list.length = 0
}

// DeleteAtPosition deletes the node at the specified position.
func (list *SinglyLinkedList) DeleteAtPosition(position int) (int, bool, error) {
	if list.isInvalidPosition(position) {
		// return an error instead
		return 0, false, fmt.Errorf("%w: %d", ErrInvalidPosition, position)
	}

	if list.IsEmpty() {
		return 0, false, nil
	}

	var value int
	if position == 0 {
		value = list.head.Data
		list.head = list.head.Next

		if list.head == nil {
			list.tail = nil
		}

		list.length--

		return value, true, nil
	}

	current := list.head
	for i := 0; i < position-1; i++ {
		current = current.Next
	}

	value = current.Next.Data
	current.Next = current.Next.Next

	if current.Next == nil {
		list.tail = current
	}

	list.length--

	return value, true, nil
}

// DeleteHead deletes the head node.
func (list *SinglyLinkedList) DeleteHead() (int, bool) {
	if list.IsEmpty() {
		return 0, false
	}

	value := list.head.data
	list.head = list.head.next
	list.length--

	return value, true
}

// DeleteTail deletes the tail node.
func (list *SinglyLinkedList) DeleteTail() (int, bool) {
	if list.IsEmpty() {
		return 0, false
	}

	if list.length == 1 {
		value := list.head.data
		list.ClearList()

		return value, true
	}

	current := list.head
	for current.next.next != nil {
		current = current.next
	}

	value := current.next.data
	current.next = nil
	list.length--

	return value, true
}

// IsEmpty checks if the list is empty.
func (list *SinglyLinkedList) IsEmpty() bool {
	return list.length == 0
}

// isInvalidPosition checks to see if the position is invalid.
func (list *SinglyLinkedList) isInvalidPosition(position int) bool {
	return position < 0 || position > list.length
}

// IsValuePresent checks if a value is present in the list.
func (list *SinglyLinkedList) IsValuePresent(data int) bool {
	if !list.IsEmpty() {
		temp := list.head

		for temp != nil {
			if temp.data == data {
				return true
			}

			temp = temp.next
		}
	}

	return false
}

// InsertAtPosition inserts a node at the specified position.
func (list *SinglyLinkedList) InsertAtPosition(data int, position int) {
	if list.isInvalidPosition(position) {
		// return error
		return
	}

	newNode := NewListNodeWithData(data)
	if list.SetHeadIfEmpty(newNode) {
		return
	}

	current := list.head
	for i := 0; i < position-1; i++ {
		current = current.next
	}

	newNode.next = current.next
	current.next = newNode
	list.length++
}

// InsertHead inserts a node at the head of the list.
func (list *SinglyLinkedList) InsertHead(data int) {
	newNode := NewListNodeWithData(data)
	if list.SetHeadIfEmpty(newNode) {
		return
	}

	newNode.next = list.head
	list.head = newNode
	list.length++
}

// InsertTail inserts a node at the tail of the list.
func (list *SinglyLinkedList) InsertTail(data int) {
	newNode := NewListNodeWithData(data)
	if list.SetHeadIfEmpty(newNode) {
		return
	}

	list.tail.next = newNode
	list.tail = newNode
	list.length++
}

// SetHeadIfEmpty sets the head if the list is empty or position is invalid.
func (list *SinglyLinkedList) SetHeadIfEmpty(newNode *ListNode) bool {
	if list.IsEmpty() {
		list.head = newNode
		list.tail = newNode
		list.length++

		return true
	}

	return false
}

// GetHead returns the head of the list.
func (list *SinglyLinkedList) GetHead() *ListNode {
	return list.head
}

// GetTail returns the tail of the list.
func (list *SinglyLinkedList) GetTail() *ListNode {
	return list.tail
}

// Size returns the size of the list.
func (list *SinglyLinkedList) Size() int {
	return list.length
}
