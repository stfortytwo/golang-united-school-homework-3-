package homework

func reverse(input []int64) (result []int64) {
	reversedSlice := make([]int64, 0, cap(input))
	for i := len(input) - 1; i >= 0; i-- {
		reversedSlice = append(reversedSlice, input[i])
	}
	return reversedSlice
}
