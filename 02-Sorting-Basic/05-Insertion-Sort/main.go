package main

import (
	"github.com/donng/Play-with-Data-Algorithms/02-Sorting-Basic/05-Insertion-Sort/helper"
	"github.com/donng/Play-with-Data-Algorithms/02-Sorting-Basic/05-Insertion-Sort/insertionsort"
	"github.com/donng/Play-with-Data-Algorithms/02-Sorting-Basic/05-Insertion-Sort/selectionsort"
)

func main() {
	n := 10000
	arr := helper.GenerateRandomArray(n, 0, n)
	arr2 := make([]int, n)
	copy(arr2, arr)

	helper.TestSort(insertionsort.Sort, arr)
	helper.TestSort(selectionsort.Sort, arr2)
}
