package main

import (
	"fmt"
	"practice/Play-with-Algorithms/02-Sorting-Basic/Optional-02-Bubble-Sort/src/helper"
	"practice/Play-with-Algorithms/02-Sorting-Basic/Optional-02-Bubble-Sort/src/InsertionSort"
	"practice/Play-with-Algorithms/02-Sorting-Basic/Optional-02-Bubble-Sort/src/SelectionSort"
	"practice/Play-with-Algorithms/02-Sorting-Basic/Optional-02-Bubble-Sort/src/BubbleSort"
	"practice/Play-with-Algorithms/02-Sorting-Basic/Optional-02-Bubble-Sort/src/BubbleSort2"
	"practice/Play-with-Algorithms/02-Sorting-Basic/Optional-02-Bubble-Sort/src/BubbleSort3"
)

func main() {
	N := 20000

	// 测试1 一般测试
	fmt.Println("Test for random array, size = ", N, ", random range [0, ", N, "]")

	arr2 := make([]int, N)
	arr3 := make([]int, N)
	arr4 := make([]int, N)
	arr5 := make([]int, N)
	arr1 := helper.GenerateRandomArray(N, 0, N)
    copy(arr2, arr1)
    copy(arr3, arr1)
    copy(arr4, arr1)
    copy(arr5, arr1)

	helper.TestSort(InsertionSort.Sort, arr1)
	helper.TestSort(SelectionSort.Sort, arr2)
	helper.TestSort(BubbleSort.Sort, arr3)
	helper.TestSort(BubbleSort2.Sort, arr4)
	helper.TestSort(BubbleSort3.Sort, arr5)

	fmt.Println()

	// 测试2 有序性更强的测试
	fmt.Println("Test for more ordered random array, size = ", N, " ,random range [0,3]")
	arr1 = helper.GenerateRandomArray(N, 0, 3)
	copy(arr2, arr1)
	copy(arr3, arr1)

	helper.TestSort(InsertionSort.Sort, arr1)
	helper.TestSort(SelectionSort.Sort, arr2)
	helper.TestSort(BubbleSort.Sort, arr3)
	helper.TestSort(BubbleSort2.Sort, arr4)
	helper.TestSort(BubbleSort3.Sort, arr5)

	fmt.Println()

	// 测试3 测试近乎有序的数组
	swapTimes := 100
	fmt.Println("Test for nearly ordered array, size = ", N, " , swap time = ", swapTimes)

	arr1 = helper.GenerateNearlyOrderedArray(N, swapTimes)
	copy(arr2, arr1)
	copy(arr3, arr1)

	helper.TestSort(InsertionSort.Sort, arr1)
	helper.TestSort(SelectionSort.Sort, arr2)
	helper.TestSort(BubbleSort.Sort, arr3)
	helper.TestSort(BubbleSort2.Sort, arr4)
	helper.TestSort(BubbleSort3.Sort, arr5)
}
