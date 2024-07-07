package linkedlist

import "fmt"

// DoublyLinkedList represents a singly linked list.
type DoublyLinkedList struct {
	head   *ListNode
	tail   *ListNode
	length int
}

// NewEmptyDoublyLinkedList creates an empty linked list.
func NewEmptyDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// NewDoublyLinkedListWithHead creates a linked list with an initial head node.
func NewDoublyLinkedListWithHead(data int) *DoublyLinkedList {
	newNode := NewListNodeWithData(data)
	return &DoublyLinkedList{
		head:   newNode,
		tail:   newNode,
		length: 1,
	}
}

// ClearList clears the linked list.
func (list *DoublyLinkedList) ClearList() {
	list.head = nil
	list.tail = nil
	list.length = 0
}

// DeleteAtPosition deletes the node at the specified position.
func (list *DoublyLinkedList) DeleteAtPosition(position int) (int, bool, error) {
	if list.isInvalidPosition(position) {
		return 0, false, fmt.Errorf("invalid position")
	}

	if list.IsEmpty() {
		return 0, false, nil
	}

	var value int
	if position == 0 {
		value = list.head.data
		list.head = list.head.next
		if list.head != nil {
			list.head.prev = nil
		} else {
			list.tail = nil
		}

		list.length--

		return value, true, nil
	}

	current := list.head
	for i := 0; i < position-1; i++ {
		current = current.next
	}

	value = current.next.data
	current.next = current.next.next
	if current.next != nil {
		current.next.prev = current
	} else {
		list.tail = current
	}

	list.length--

	return value, true, nil
}

// DeleteHead deletes the head node.
func (list *DoublyLinkedList) DeleteHead() (int, bool) {
	if list.IsEmpty() {
		return 0, false
	}

	value := list.head.data
	list.head = list.head.next
	if list.head != nil {
		list.head.prev = nil
	} else {
		list.tail = nil
	}

	list.length--

	return value, true
}

// DeleteTail deletes the tail node.
func (list *DoublyLinkedList) DeleteTail() (int, bool) {
	if list.IsEmpty() {
		return 0, false
	}

	if list.length == 1 {
		value := list.head.data
		list.ClearList()

		return value, true
	}

	value := list.tail.data
	current := list.tail.prev
	current.next = nil
	list.tail = current
	list.length--

	return value, true
}

// IsEmpty checks if the list is empty.
func (list *DoublyLinkedList) IsEmpty() bool {
	return list.length == 0
}

// isInvalidPosition checks to see if the position is invalid.
func (list *DoublyLinkedList) isInvalidPosition(position int) bool {
	return position < 0 || position > list.length
}

// IsValuePresent checks if a value is present in the list.
func (list *DoublyLinkedList) IsValuePresent(data int) bool {
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
func (list *DoublyLinkedList) InsertAtPosition(data int, position int) {
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
	newNode.prev = current
	current.next = newNode
	list.length++
}

// InsertHead inserts a node at the head of the list.
func (list *DoublyLinkedList) InsertHead(data int) {
	newNode := NewListNodeWithData(data)
	if list.SetHeadIfEmpty(newNode) {
		return
	}

	newNode.next = list.head
	list.head.prev = newNode
	list.head = newNode
	list.length++
}

// InsertTail inserts a node at the tail of the list.
func (list *DoublyLinkedList) InsertTail(data int) {
	newNode := NewListNodeWithData(data)
	if list.SetHeadIfEmpty(newNode) {
		return
	}

	list.tail.next = newNode
	newNode.prev = list.tail
	list.tail = newNode
	list.length++
}

// SetHeadIfEmpty sets the head if the list is empty or position is invalid.
func (list *DoublyLinkedList) SetHeadIfEmpty(node *ListNode) bool {
	if list.IsEmpty() {
		list.head = node
		list.tail = node
		list.length++
		return true
	}

	return false
}

// Size returns the size of the list.
func (list *DoublyLinkedList) Size() int {
	return list.length
}
