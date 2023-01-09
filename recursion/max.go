package recursion

import "algorithms/common"

func FindMax[T common.Numeric](array []T) T {
	var findMax func(max T, arr []T) T

	findMax = func(max T, arr []T) T {
		if len(arr) == 2 {
			if arr[0] > arr[1] {
				return arr[0]
			} else {
				return arr[1]
			}
		} else {
			m := findMax(arr[0], arr[1:])
			if m > arr[0] {
				return m
			} else {
				return arr[0]
			}
		}
	}

	return findMax(array[0], array[1:])
}
