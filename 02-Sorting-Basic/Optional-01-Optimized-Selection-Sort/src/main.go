package main

import (
	"fmt"
	"practice/Play-with-Algorithms/02-Sorting-Basic/Optional-01-Optimized-Selection-Sort/src/helper"
	"practice/Play-with-Algorithms/02-Sorting-Basic/Optional-01-Optimized-Selection-Sort/src/InsertionSort"
	"practice/Play-with-Algorithms/02-Sorting-Basic/Optional-01-Optimized-Selection-Sort/src/SelectionSort"
	"practice/Play-with-Algorithms/02-Sorting-Basic/Optional-01-Optimized-Selection-Sort/src/SelectionSort2"
)

func main() {
	N := 20000

	// 测试1 一般测试
	fmt.Println("Test for random array, size = ", N, ", random range [0, ", N, "]")

	arr2 := make([]int, N)
	arr3 := make([]int, N)
	arr1 := helper.GenerateRandomArray(N, 0, N)
    copy(arr2, arr1)
    copy(arr3, arr1)

	helper.TestSort(InsertionSort.Sort, arr1)
	helper.TestSort(SelectionSort.Sort, arr2)
	helper.TestSort(SelectionSort2.Sort, arr3)

	fmt.Println()

	// 测试2 有序性更强的测试
	fmt.Println("Test for more ordered random array, size = ", N, " ,random range [0,3]")
	arr1 = helper.GenerateRandomArray(N, 0, 3)
	copy(arr2, arr1)
	copy(arr3, arr1)

	helper.TestSort(InsertionSort.Sort, arr1)
	helper.TestSort(SelectionSort.Sort, arr2)
	helper.TestSort(SelectionSort2.Sort, arr3)

	fmt.Println()

	// 测试3 测试近乎有序的数组
	swapTimes := 100
	fmt.Println("Test for nearly ordered array, size = ", N, " , swap time = ", swapTimes)

	arr1 = helper.GenerateNearlyOrderedArray(N, swapTimes)
	copy(arr2, arr1)
	copy(arr3, arr1)

	helper.TestSort(InsertionSort.Sort, arr1)
	helper.TestSort(SelectionSort.Sort, arr2)
	helper.TestSort(SelectionSort2.Sort, arr3)
}
