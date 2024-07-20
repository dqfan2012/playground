// Package stack implements the stack data structure.
package stack

import "errors"

// ErrEmptyStack is an error indicating the stack is empty when attempting to
// peek at or remove the element on the top of the stack.
var ErrEmptyStack = errors.New("empty stack")

// Stacker defines the operations a stack should implement.
type Stacker[T comparable] interface {
	Push(T)
	Pop() (*T, error)
	Peek() (*T, error)
	IsEmpty()
	Len()
}
