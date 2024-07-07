package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListInitialization(t *testing.T) {
	linkedList := NewEmptySinglyLinkedList()
	assert.Equal(t, 0, linkedList.length)
	assert.Nil(t, linkedList.head)

	linkedListWithHead := NewSinglyLinkedListWithHead(10)
	assert.Equal(t, 1, linkedListWithHead.length)
	assert.Equal(t, 10, linkedListWithHead.head.data)
}

func TestInsertHead(t *testing.T) {
	linkedList := NewEmptySinglyLinkedList()
	linkedList.InsertHead(10)
	linkedList.InsertHead(20)

	assert.Equal(t, 2, linkedList.length)
	assert.Equal(t, 20, linkedList.head.data)
	assert.Equal(t, 10, linkedList.head.next.data)
}

func TestInsertTail(t *testing.T) {
	linkedList := NewEmptySinglyLinkedList()
	linkedList.InsertTail(10)
	linkedList.InsertTail(20)

	assert.Equal(t, 2, linkedList.length)
	assert.Equal(t, 10, linkedList.head.data)
	assert.Equal(t, 20, linkedList.head.next.data)
}

func TestInsertAtPosition(t *testing.T) {
	linkedList := NewEmptySinglyLinkedList()
	linkedList.InsertHead(10)
	linkedList.InsertHead(20)
	linkedList.InsertAtPosition(15, 1)

	assert.Equal(t, 3, linkedList.length)
	assert.Equal(t, 20, linkedList.head.data)
	assert.Equal(t, 15, linkedList.head.next.data)
	assert.Equal(t, 10, linkedList.head.next.next.data)
}

func TestDeleteHead(t *testing.T) {
	linkedList := NewEmptySinglyLinkedList()
	linkedList.InsertHead(10)
	linkedList.InsertHead(20)
	value, ok := linkedList.DeleteHead()

	assert.True(t, ok)
	assert.Equal(t, 20, value)
	assert.Equal(t, 1, linkedList.length)
	assert.Equal(t, 10, linkedList.head.data)
}

func TestDeleteTail(t *testing.T) {
	// Test deleting tail from a list with multiple nodes
	linkedList := NewEmptySinglyLinkedList()
	linkedList.InsertHead(10)
	linkedList.InsertHead(20)
	value, ok := linkedList.DeleteTail()

	assert.True(t, ok)
	assert.Equal(t, 10, value)
	assert.Equal(t, 1, linkedList.Size())
	assert.Equal(t, 20, linkedList.head.data)
	assert.Nil(t, linkedList.head.next) // Tail should be nil after deletion

	// Test deleting tail from a list with only one node
	linkedList = NewEmptySinglyLinkedList()
	linkedList.InsertHead(10)
	value, ok = linkedList.DeleteTail()

	assert.True(t, ok)
	assert.Equal(t, 10, value)
	assert.Equal(t, 0, linkedList.Size())
	assert.Nil(t, linkedList.head) // Head should be nil after deletion
	assert.Nil(t, linkedList.tail) // Tail should be nil after deletion
}

func TestDeleteAtPosition(t *testing.T) {
	linkedList := NewEmptySinglyLinkedList()
	linkedList.InsertHead(10)
	linkedList.InsertHead(20)
	linkedList.InsertHead(30)

	// Delete node at position 1 (20)
	value, ok, _ := linkedList.DeleteAtPosition(1)
	assert.True(t, ok)
	assert.Equal(t, 20, value)
	assert.Equal(t, 2, linkedList.Size())
	assert.Equal(t, 30, linkedList.head.data)
	assert.Equal(t, 10, linkedList.head.next.data)

	// Delete node at position 0 (30)
	value, ok, _ = linkedList.DeleteAtPosition(0)
	assert.True(t, ok)
	assert.Equal(t, 30, value)
	assert.Equal(t, 1, linkedList.Size())
	assert.Equal(t, 10, linkedList.head.data)
	assert.Nil(t, linkedList.head.next)

	// Delete node at position 0 (10)
	value, ok, _ = linkedList.DeleteAtPosition(0)
	assert.True(t, ok)
	assert.Equal(t, 10, value)
	assert.Equal(t, 0, linkedList.Size())
	assert.Nil(t, linkedList.head)

	// Test deleting tail from a list with only one node
	linkedList.InsertHead(40)
	value, ok, _ = linkedList.DeleteAtPosition(0)
	assert.True(t, ok)
	assert.Equal(t, 40, value)
	assert.Equal(t, 0, linkedList.Size())
	assert.Nil(t, linkedList.head)
}

func TestIsEmpty(t *testing.T) {
	linkedList := NewEmptySinglyLinkedList()
	assert.True(t, linkedList.IsEmpty())

	linkedList.InsertHead(10)
	assert.False(t, linkedList.IsEmpty())
}

func TestIsValuePresent(t *testing.T) {
	linkedList := NewEmptySinglyLinkedList()
	linkedList.InsertHead(10)
	linkedList.InsertHead(20)

	assert.True(t, linkedList.IsValuePresent(10))
	assert.False(t, linkedList.IsValuePresent(30))
}

func TestSize(t *testing.T) {
	linkedList := NewEmptySinglyLinkedList()
	assert.Equal(t, 0, linkedList.Size())

	linkedList.InsertHead(10)
	linkedList.InsertHead(20)
	assert.Equal(t, 2, linkedList.Size())
}
