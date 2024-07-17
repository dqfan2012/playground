# LinkedList Package

## Overview

The `linkedlist` package provides implementations for both singly and doubly linked lists in Go. This package is designed to offer efficient and easy-to-use data structures for managing collections of elements with dynamic sizes.

## Features

- **Singly Linked List**: Efficient for appending and sequential traversal.
- **Doubly Linked List**: Allows traversal in both directions and efficient element removal.

## Installation

To install the `linkedlist` package, use the following command:

```bash
go get github.com/dqfan2012/playground/pkg/ds/linkedlist
```

## Complexities

**Singly Linked List**

- DeleteAtPosition(position int): $`O(n)`$
- DeleteHead(): $`O(1)`$
- DeleteTail(): $`O(1)`$
- InsertAtPosition(any, int): $`O(n)`$
- InsertHead(data any): $`O(1)`$
- InsertTail(data any): $`O(1)`$
- GetHead(): $`O(1)`$
- GetTail(): $`O(1)`$

Space Complexity: $`O(n)`$

**Doubly Linked List**

- DeleteAtPosition(position int): $`O(n)`$
- DeleteHead(): $`O(1)`$
- DeleteTail(): $`O(1)`$
- InsertAtPosition(any, int): $`O(n)`$
- InsertHead(data any): $`O(1)`$
- InsertTail(data any): $`O(1)`$
- GetHead(): $`O(1)`$
- GetTail(): $`O(1)`$

Space Complexity: $`O(n)`$

**ListHelper Interface**

- SetHeadIfEmpty(newNode *ListNode[T]): $`O(1)`$
- ClearList(): $`O(1)`$
- IsEmpty(): $`O(1)`$
- IsPresent(data any): $`O(n)`$
- Len(): $`O(1)`$