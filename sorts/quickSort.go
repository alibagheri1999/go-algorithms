package sorts

func QuickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	pivot := array[0]
	left, right := []int{}, []int{}

	for _, v := range array[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)
	return append(append(left, pivot), right...)
}
