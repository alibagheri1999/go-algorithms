package sorts

func HeapSort(array []int) []int {
	n := len(array)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(array, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		heapify(array, i, 0)
	}
	return array
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}
