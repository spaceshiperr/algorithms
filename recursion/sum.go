package recursion

import "algorithms/common"

func SumArray[T common.Numeric](array []T) T {
	if len(array) == 0 {
		return 0
	} else {
		return array[0] + SumArray(array[1:])
	}
}
