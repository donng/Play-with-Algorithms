package main

import (
	"practice/Play-with-Algorithms/02-Sorting-Basic/06-Insertion-Sort-Advance/src/InsertionSort"
	"practice/Play-with-Algorithms/02-Sorting-Basic/06-Insertion-Sort-Advance/src/SelectionSort"
	"practice/Play-with-Algorithms/02-Sorting-Basic/06-Insertion-Sort-Advance/src/helper"
)

func main() {
	n := 10000
	arr := helper.GenerateNearlyOrderedArray(n, 100)
	arr2 := make([]int, n)
	copy(arr2, arr)

	helper.TestSort(InsertionSort.Sort, arr)
	helper.TestSort(SelectionSort.Sort, arr2)
}
