package recursion

import (
	"algorithms/common"
	"errors"
)

func BinarySearch[T common.Numeric](array []T, item T) (int, error) {
	var search func(array []T, item T, p int) (int, error)

	search = func(array []T, item T, p int) (int, error) {
		if len(array) == 0 {
			return -1, errors.New("not found")
		} else if array[0] == item {
			return p, nil
		} else {
			return search(array[1:], item, p+1)
		}
	}

	return search(array, item, 0)
}
