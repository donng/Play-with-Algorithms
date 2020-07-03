package main

import (
	"fmt"
	"github.com/donng/Play-with-Data-Algorithms/02-Sorting-Basic/Optional-03-Shell-Sort/bubblesort"
	"github.com/donng/Play-with-Data-Algorithms/02-Sorting-Basic/Optional-03-Shell-Sort/bubblesort2"
	"github.com/donng/Play-with-Data-Algorithms/02-Sorting-Basic/Optional-03-Shell-Sort/bubblesort3"
	"github.com/donng/Play-with-Data-Algorithms/02-Sorting-Basic/Optional-03-Shell-Sort/helper"
	"github.com/donng/Play-with-Data-Algorithms/02-Sorting-Basic/Optional-03-Shell-Sort/insertionsort"
	"github.com/donng/Play-with-Data-Algorithms/02-Sorting-Basic/Optional-03-Shell-Sort/selectionsort"
	"github.com/donng/Play-with-Data-Algorithms/02-Sorting-Basic/Optional-03-Shell-Sort/shellsort"
)

func main() {
	N := 20000

	// 测试1 一般测试
	fmt.Println("Test for random array, size = ", N, ", random range [0, ", N, "]")

	arr2 := make([]int, N)
	arr3 := make([]int, N)
	arr4 := make([]int, N)
	arr5 := make([]int, N)
	arr6 := make([]int, N)
	arr1 := helper.GenerateRandomArray(N, 0, N)
	copy(arr2, arr1)
	copy(arr3, arr1)
	copy(arr4, arr1)
	copy(arr5, arr1)
	copy(arr6, arr1)

	helper.TestSort(insertionsort.Sort, arr1)
	helper.TestSort(selectionsort.Sort, arr2)
	helper.TestSort(bubblesort.Sort, arr3)
	helper.TestSort(bubblesort2.Sort, arr4)
	helper.TestSort(bubblesort3.Sort, arr5)
	helper.TestSort(shellsort.Sort, arr6)

	fmt.Println()

	// 测试2 有序性更强的测试
	fmt.Println("Test for more ordered random array, size = ", N, " ,random range [0,3]")
	arr1 = helper.GenerateRandomArray(N, 0, 3)
	copy(arr2, arr1)
	copy(arr3, arr1)

	helper.TestSort(insertionsort.Sort, arr1)
	helper.TestSort(selectionsort.Sort, arr2)
	helper.TestSort(bubblesort.Sort, arr3)
	helper.TestSort(bubblesort2.Sort, arr4)
	helper.TestSort(bubblesort3.Sort, arr5)
	helper.TestSort(shellsort.Sort, arr6)

	fmt.Println()

	// 测试3 测试近乎有序的数组
	swapTimes := 100
	fmt.Println("Test for nearly ordered array, size = ", N, " , swap time = ", swapTimes)

	arr1 = helper.GenerateNearlyOrderedArray(N, swapTimes)
	copy(arr2, arr1)
	copy(arr3, arr1)

	helper.TestSort(insertionsort.Sort, arr1)
	helper.TestSort(selectionsort.Sort, arr2)
	helper.TestSort(bubblesort.Sort, arr3)
	helper.TestSort(bubblesort2.Sort, arr4)
	helper.TestSort(bubblesort3.Sort, arr5)
}
