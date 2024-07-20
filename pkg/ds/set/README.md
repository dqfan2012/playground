# Set Package

This package provides an implementation of a set data structure in Go. It is designed to be thread-safe and efficient for concurrent use.

## Installation

To install this package, use `go get`:

```sh
go get github.com/dqfan2012/playground/pkg/ds/set
```

## Complexities

Time Complexities

- Add(): $`O(1)`$
- Remove(): $`O(1)`$
- Contains(): $`O(1)`$
- Union(): $`O(n)`$
- Intersection(): $`O(n)`$
- Difference(): $`O(n)`$
- IsSubset(): $`O(n)`$
- Clear(): $`O(1)`$
- IsEmpty(): $`O(1)`$
- Len(): $`O(1)`$

Space Complexity: $`O(n)`$, where `n` is the number of elements in the set.
