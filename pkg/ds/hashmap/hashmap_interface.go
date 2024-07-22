// Package hashmap implements the hashmap data structure.
package hashmap

// Mapper defines all the operations for hashmaps.
type Mapper[T comparable] interface {
	Insert(string, T)
	Delete(string)
	Get(string) (*T, bool)
}
