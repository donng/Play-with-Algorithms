package main

import (
	"github.com/donng/Play-with-Data-Algorithms/02-Sorting-Basic/06-Insertion-Sort-Advance/helper"
	"github.com/donng/Play-with-Data-Algorithms/02-Sorting-Basic/06-Insertion-Sort-Advance/insertionsort"
	"github.com/donng/Play-with-Data-Algorithms/02-Sorting-Basic/06-Insertion-Sort-Advance/selectionsort"
)

func main() {
	n := 10000
	arr := helper.GenerateNearlyOrderedArray(n, 100)
	arr2 := make([]int, n)
	copy(arr2, arr)

	helper.TestSort(insertionsort.Sort, arr)
	helper.TestSort(selectionsort.Sort, arr2)
}
