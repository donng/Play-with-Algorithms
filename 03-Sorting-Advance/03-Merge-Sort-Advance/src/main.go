package main

import (
	"practice/Play-with-Algorithms/03-Sorting-Advance/03-Merge-Sort-Advance/src/helper"
	"practice/Play-with-Algorithms/03-Sorting-Advance/03-Merge-Sort-Advance/src/InsertionSort"
	"practice/Play-with-Algorithms/03-Sorting-Advance/03-Merge-Sort-Advance/src/MergeSort"
	"fmt"
	"practice/Play-with-Algorithms/03-Sorting-Advance/03-Merge-Sort-Advance/src/MergeSort2"
)

// 比较InsertionSort和MergeSort两种排序算法的性能效率
// 整体而言, MergeSort的性能最优, 对于近乎有序的数组的特殊情况, 见测试2的详细注释
func main() {
	N := 500000

	// 测试1 一般测试
	fmt.Println("Test for random array, size =", N, ", random range [0, ", N, "]")

	arr1 := helper.GenerateRandomArray(N, 0, N)
	arr2 := make([]int, N)
	arr3 := make([]int, N)
	copy(arr2, arr1)
	copy(arr3, arr1)

	helper.TestSort(InsertionSort.Sort, arr1)
	helper.TestSort(MergeSort.MergeSort, arr2)
	helper.TestSort(MergeSort2.MergeSort, arr3)

	// 测试2 测试近乎有序的数组
	swapTimes := 10

	fmt.Println("Test for nearly ordered array, size= ", N, " swap time = ", swapTimes)

	arr1 = helper.GenerateNearlyOrderedArray(N, swapTimes)
	copy(arr2, arr1)
	copy(arr3, arr1)

	helper.TestSort(InsertionSort.Sort, arr1)
	helper.TestSort(MergeSort.MergeSort, arr2)
	helper.TestSort(MergeSort2.MergeSort, arr3)
}
