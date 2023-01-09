package search

import (
	"algorithms/common"
	"errors"
)

func BinarySearch[T common.Numeric](array []T, item T) (int, error) {
	left, right := 0, len(array)-1
	mid := (left + right) / 2

	for left <= right {
		if array[mid] == item {
			return mid, nil
		} else if array[mid] < item {
			left = mid + 1
		} else {
			right = mid - 1
		}

		mid = (left + right) / 2
	}

	return -1, errors.New("not found")
}
