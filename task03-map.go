package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	keysSlice := make([]int, 0, len(input))
	valueSlice := make([]string, 0, len(input))
	for key := range input {
		keysSlice = append(keysSlice, key)
	}
	sort.SliceStable(keysSlice, func(i, j int) bool { return keysSlice[i] < keysSlice[j] })
	for key := range keysSlice {
		valueSlice = append(valueSlice, input[key])
	}
	return valueSlice
}
