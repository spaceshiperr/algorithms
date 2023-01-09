package main

import (
	"algorithms/search"
	"fmt"
)

func main() {
	//array := []int{1, 2, 3, 110, 4, 5, 9}

	//if result, err := search.BinarySearch(array, 2); err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println(result)
	//}

	//fmt.Println(sort.SelectionSort(array))

	//fmt.Println(recursion.SumArray(array))

	//fmt.Println(recursion.FindMax(array))

	//fmt.Println(recursion.BinarySearch(array, 9))

	var book search.IPhonebook
	book = search.NewPhonebook()

	book.Add("mary", "213")
	book.Add("alex", "123123")
	fmt.Printf(book.Get("alex2"))
}
