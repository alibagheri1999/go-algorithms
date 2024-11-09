package sorts

// Slice function to create a subarray from start to end indices
func Slice(array []int, start, end int) []int {
	var result []int

	if start < 0 {
		start = len(array) + start
	}
	if end < 0 {
		end = len(array) + end
	}
	start = max(0, start)
	end = min(len(array), end)

	for i := start; i < end; i++ {
		result = append(result, array[i])
	}
	return result
}
