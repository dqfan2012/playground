// Package set implements the set data structure.
package set

// Setter defines all the operations for a Set.
type Setter[T comparable] interface {
	Add(T)
	Remove(T)
	Contains(T)
	Union(*Set[T]) *Set[T]
	Intersection(*Set[T]) *Set[T]
	Difference(*Set[T]) *Set[T]
	IsSubset(*Set[T]) bool
	Clear()
	IsEmpty() bool
	Len() int
}
