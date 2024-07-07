package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedListInitialization(t *testing.T) {
	linkedList := NewEmptyDoublyLinkedList()
	assert.Equal(t, 0, linkedList.Size())
	assert.True(t, linkedList.IsEmpty())

	linkedListWithHead := NewDoublyLinkedListWithHead(10)
	assert.Equal(t, 1, linkedListWithHead.Size())
	assert.Equal(t, 10, linkedListWithHead.head.data)
	assert.Equal(t, 10, linkedListWithHead.tail.data)
}

func TestDoublyLinkedListInsertHead(t *testing.T) {
	linkedList := NewEmptyDoublyLinkedList()
	linkedList.InsertHead(10)
	linkedList.InsertHead(20)

	assert.Equal(t, 2, linkedList.Size())
	assert.Equal(t, 20, linkedList.head.data)
	assert.Equal(t, 10, linkedList.head.next.data)
}

func TestDoublyLinkedListInsertTail(t *testing.T) {
	linkedList := NewEmptyDoublyLinkedList()
	linkedList.InsertTail(10)
	linkedList.InsertTail(20)

	assert.Equal(t, 2, linkedList.Size())
	assert.Equal(t, 10, linkedList.head.data)
	assert.Equal(t, 20, linkedList.head.next.data)
}

func TestDoublyLinkedListInsertAtPosition(t *testing.T) {
	linkedList := NewEmptyDoublyLinkedList()
	linkedList.InsertHead(10)
	linkedList.InsertHead(20)
	linkedList.InsertAtPosition(15, 1)

	assert.Equal(t, 3, linkedList.Size())
	assert.Equal(t, 20, linkedList.head.data)
	assert.Equal(t, 15, linkedList.head.next.data)
	assert.Equal(t, 10, linkedList.head.next.next.data)
}

func TestDoublyLinkedListDeleteHead(t *testing.T) {
	linkedList := NewEmptyDoublyLinkedList()
	linkedList.InsertHead(10)
	linkedList.InsertHead(20)
	value, ok := linkedList.DeleteHead()

	assert.True(t, ok)
	assert.Equal(t, 20, value)
	assert.Equal(t, 1, linkedList.Size())
	assert.Equal(t, 10, linkedList.head.data)
}

func TestDoublyLinkedListDeleteTail(t *testing.T) {
	linkedList := NewEmptyDoublyLinkedList()
	linkedList.InsertHead(10)
	linkedList.InsertHead(20)
	value, ok := linkedList.DeleteTail()

	assert.True(t, ok)
	assert.Equal(t, 10, value)
	assert.Equal(t, 1, linkedList.Size())
	assert.Equal(t, 20, linkedList.head.data)
	assert.Nil(t, linkedList.head.next) // Tail should be nil after deletion
}

func TestDoublyLinkedListDeleteAtPosition(t *testing.T) {
	linkedList := NewEmptyDoublyLinkedList()
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
}

func TestDoublyLinkedListIsEmpty(t *testing.T) {
	linkedList := NewEmptyDoublyLinkedList()
	assert.True(t, linkedList.IsEmpty())

	linkedList.InsertHead(10)
	assert.False(t, linkedList.IsEmpty())
}

func TestDoublyLinkedListIsValuePresent(t *testing.T) {
	linkedList := NewEmptyDoublyLinkedList()
	linkedList.InsertHead(10)
	linkedList.InsertHead(20)

	assert.True(t, linkedList.IsValuePresent(10))
	assert.False(t, linkedList.IsValuePresent(30))
}

func TestDoublyLinkedListSize(t *testing.T) {
	linkedList := NewEmptyDoublyLinkedList()
	assert.Equal(t, 0, linkedList.Size())

	linkedList.InsertHead(10)
	linkedList.InsertHead(20)
	assert.Equal(t, 2, linkedList.Size())
}
