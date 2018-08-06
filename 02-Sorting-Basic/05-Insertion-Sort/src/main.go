package main

import (
	"practice/Play-with-Algorithms/02-Sorting-Basic/05-Insertion-Sort/src/helper"
	"practice/Play-with-Algorithms/02-Sorting-Basic/05-Insertion-Sort/src/InsertionSort"
	"practice/Play-with-Algorithms/02-Sorting-Basic/05-Insertion-Sort/src/SelectionSort"
)

func main() {
	n := 10000
	arr := helper.GenerateRandomArray(n, 0, n)
	arr2 := make([]int, n)
	copy(arr2, arr)

	helper.TestSort(InsertionSort.Sort, arr)
	helper.TestSort(SelectionSort.Sort, arr2)
}
