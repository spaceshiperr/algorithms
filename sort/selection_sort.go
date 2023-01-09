package sort

import (
	"algorithms/common"
)

func SelectionSort[T common.Numeric](array []T) []T {
	result := make([]T, 0, len(array))

	for range array {
		min, i := findSmallest(array)

		result = append(result, min)
		array = append(array[:i], array[i+1:]...)
	}

	return result
}

func findSmallest[T common.Numeric](array []T) (T, int) {
	min, min_i := array[0], 0

	for i, value := range array {
		if value < min {
			min, min_i = value, i
		}
	}

	return min, min_i
}
