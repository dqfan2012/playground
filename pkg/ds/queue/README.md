# Queue Package

This package provides implementations of both a classical queue and a priority queue in Go. It is designed to be thread-safe and efficient for concurrent use.

## Installation

To install this package, use `go get`:

```sh
go get github.com/dqfan2012/playground/pkg/ds/queue
```

## Complexities

**Classic Queue:**

- Enqueue(): $`O(1)`$
- Dequeue(): $`O(1)`$
- IsEmpty(): $`O(1)`$
- Len(): $`O(1)`$

Space Complexity: $`O(n)`$

**Priority Queue:**

- Enqueue(): $`O(n)`$ - Adding to the queue requires arranging items based on their priority. In the worst case, this is $`O(n)`$, where `n` is the number of items in the queue.
- Dequeue(): $`O(1)`$
- IsEmpty(): $`O(1)`$
- Len(): $`O(1)`$

Space Complexity: $`O(n)`$