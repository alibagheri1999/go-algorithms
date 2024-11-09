package main

import (
	"log"
	"main/algorithms/sorts"
	"time"
)

func main() {
	array := []int{1, 5, 4, 8, 5}
	log.Println("Slice of array:", sorts.Slice(array, 0, len(array)))
	log.Println("Slice of array from index 1 to 3:", sorts.Slice(array, 1, 3))

	// Testing merge sort
	sortedArray := sorts.MergeSort(array)
	log.Println("Original array:", array, "Sorted array (Merge Sort):", sortedArray)

	// Testing quick sort
	array2 := []int{4, 2, 7, 1, 9, 3}
	start := time.Now()
	sortedArray2 := sorts.QuickSort(array2)
	log.Printf("Time taken (Quick Sort): %v\n", time.Since(start))
	log.Println("Original array2:", array2, "Sorted array2 (Quick Sort):", sortedArray2)

	// Testing heap sort
	array3 := []int{4, 2, 7, 1, 9, 3}
	start = time.Now()
	sortedArray3 := sorts.HeapSort(array3)
	log.Printf("Time taken (Heap Sort): %v\n", time.Since(start))
	log.Println("Original array3:", array3, "Sorted array3 (Heap Sort):", sortedArray3)

	// Testing bubble sort
	array4 := []int{4, 2, 7, 1, 9, 3}
	start = time.Now()
	sortedArray4 := sorts.BubbleSort(array4)
	log.Printf("Time taken (Bubble Sort): %v\n", time.Since(start))
	log.Println("Original array4:", array4, "Sorted array4 (Bubble Sort):", sortedArray4)

	log.Println("Sorting example [2, 0, 2, 1, 1, 0] (Merge Sort):", sorts.MergeSort([]int{2, 0, 2, 1, 1, 0}))
}
