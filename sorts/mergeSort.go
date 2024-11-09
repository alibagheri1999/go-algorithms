package sorts

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	middle := len(array) / 2
	left := Slice(array, 0, middle)
	right := Slice(array, middle, len(array))
	sortedLeft := MergeSort(left)
	sortedRight := MergeSort(right)
	return merge(sortedLeft, sortedRight)
}

func merge(left, right []int) []int {
	var result []int
	lIdx, rIdx := 0, 0

	for lIdx < len(left) && rIdx < len(right) {
		if left[lIdx] < right[rIdx] {
			result = append(result, left[lIdx])
			lIdx++
		} else {
			result = append(result, right[rIdx])
			rIdx++
		}
	}
	result = append(result, Slice(left, lIdx, len(left))...)
	result = append(result, Slice(right, rIdx, len(right))...)
	return result
}
