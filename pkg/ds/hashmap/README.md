# HashMap Package

## Introduction

This package is a custom implementation of a hashmap in Go. It is intended for educational purposes to help understand the underlying mechanics of hashmaps, such as hash functions, collision handling, and basic operations like insertion, retrieval, and deletion.

## Why Implement a Hashmap?

While Go provides a built-in map type that is highly optimized and should be used for practical purposes, implementing a hashmap from scratch is a valuable exercise to:

- Gain a deeper understanding of data structures.
- Learn about hash functions and collision resolution techniques.
- Develop problem-solving and coding skills.

## Go's Built-in map

In Go, the map type is a hashmap. It provides efficient key-value pair management with average-case $`O(1)`$ complexity for insertion, deletion, and look-up operations. Here is a brief example of how to use Go’s built-in map:

```go
package main

import "fmt"

func main() {
    myMap := make(map[string]int)
    myMap["key1"] = 100
    myMap["key2"] = 200

    value, exists := myMap["key1"]
    if exists {
        fmt.Println("key1:", value)
    } else {
        fmt.Println("key1 does not exist")
    }

    delete(myMap, "key1")
    _, exists = myMap["key1"]
    if !exists {
        fmt.Println("key1 successfully deleted")
    }
}
```

For most applications, you should use Go’s built-in map due to its performance and reliability.

## Custom Hashmap Implementation

### Features

- Hash Function: Converts keys into bucket indices.
- Buckets: An array of linked lists to handle collisions.
- Node Structure: Stores key-value pairs.
- Basic Operations: Insertion, retrieval, and deletion of key-value pairs.

### Limitations

- This implementation is for learning purposes and may not be as optimized as Go’s built-in map.
- It does not include advanced features like dynamic resizing or more sophisticated collision resolution techniques.

### Complexities

- Insert: $`O(1)`$ on average, $`O(n)`$ in the worst case due to collisions.
- Get: $`O(1)`$ on average, $`O(n)`$ in the worst case due to collisions.
- Delete: $`O(1)`$ on average, $`O(n)`$ in the worst case due to collisions.

Space Complexity: $`O(n)`$

## Conclusion

This project serves as an educational tool to better understand how hashmaps work. For production use, rely on Go’s built-in map type for its efficiency and ease of use.
