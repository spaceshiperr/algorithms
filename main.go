package main

import (
	"algorithms/search"
	"fmt"
)

func main() {
	m := 8
	p := []int{0, 1, 2, 5, 6}
	w := []int{0, 2, 3, 4, 5}

	profit := search.GetKnapsackMaxProfit(m, w, p)
	fmt.Print(profit)
}
