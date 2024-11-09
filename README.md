# Go Algorithms Project

This repository contains various algorithm implementations in Go. Each file represents specific algorithms, including descriptions and usage examples.

## Sorting Algorithms

### 1. Merge Sort - `arrays.go`
- **Description**: Merge Sort is a divide-and-conquer algorithm that recursively splits the array into halves, sorts each half, and merges them back together.
- **Explanation**: The `MergeSort` function divides the array into left and right subarrays, recursively sorts them, and merges them using the `merge` helper function. This approach achieves efficient sorting for large datasets.
- **Time Complexity**: O(n log n)

### 2. Quick Sort - `arrays.go`
- **Description**: Quick Sort is a divide-and-conquer algorithm that partitions the array around a pivot element and recursively sorts each partition.
- **Explanation**: The `QuickSort` function selects a pivot, partitions the array into left and right arrays, and recursively sorts each part. This implementation performs well on average but can degrade to O(n^2) in the worst case.
- **Time Complexity**: O(n log n) on average, O(n^2) worst case

### 3. Heap Sort - `arrays.go`
- **Description**: Heap Sort is a comparison-based sorting technique based on a binary heap data structure. It repeatedly extracts the maximum element to produce a sorted array.
- **Explanation**: The `HeapSort` function builds a max-heap from the input array. It then repeatedly swaps the root (maximum value) with the last element, reduces the heap size, and heapifies the root. This in-place sorting is efficient for scenarios requiring minimal extra memory.
- **Time Complexity**: O(n log n)

### 4. Bubble Sort - `arrays.go`
- **Description**: Bubble Sort is a simple comparison-based sorting algorithm that repeatedly steps through the list, swapping adjacent elements if they are in the wrong order.
- **Explanation**: The `BubbleSort` function iterates through the array multiple times, swapping elements when needed. While easy to implement, it is inefficient for large datasets due to its O(n^2) complexity.
- **Time Complexity**: O(n^2)

## Pattern Matching

### KMP Pattern Matching - `strings.go`
- **Description**: The Knuth-Morris-Pratt (KMP) algorithm is a pattern-matching algorithm that efficiently searches for occurrences of a pattern within a text.
- **Explanation**: The `search` function uses the KMP algorithm to find all occurrences of a pattern in a text. It first builds an LPS (Longest Prefix Suffix) array, which allows it to skip unnecessary comparisons, significantly speeding up the search process for repeated patterns.
- **Time Complexity**: O(n + m), where n is the length of the text and m is the length of the pattern.

## Utility Functions - `arrays.go` and `utils.go`
- **Slice**: Creates a subarray from start to end indices.
- **Max and Min**: Utility functions for finding the maximum or minimum of two integers.

## Time Complexity Summary

| Algorithm         | Time Complexity          | Space Complexity |
|-------------------|--------------------------|-------------------|
| Merge Sort        | O(n log n)               | O(n)             |
| Quick Sort        | O(n log n) (avg), O(n^2) (worst) | O(log n) |
| Heap Sort         | O(n log n)               | O(1)             |
| Bubble Sort       | O(n^2)                   | O(1)             |
| KMP Pattern Search| O(n + m)                 | O(m)             |

## Running the Code

To execute the code, use the following commands:

### Run the sorting algorithms in `arrays.go`:

```bash
go run arrays.go
