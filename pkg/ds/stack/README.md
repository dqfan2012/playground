# Stack Package

This package provides an implementation of a classical stack in Go. It is designed to be thread-safe and efficient for concurrent use.

## Installation

To install this package, use `go get`:

```sh
go get github.com/dqfan2012/playground/pkg/ds/stack
```

## Complexities

Time Complexities:

- Push(): $`O(1)`$
- Pop(): $`O(1)`$
- Peek(): $`O(1)`$
- IsEmpty(): $`O(1)`$
- Len(): $`O(1)`$

Space Complexity: $`O(n)`$, where `n` is the number of elements in the stack.
