package main

import (
	"fmt"
	"practice/Play-with-Algorithms/03-Sorting-Advance/05-Quick-Sort/src/helper"
	"practice/Play-with-Algorithms/03-Sorting-Advance/05-Quick-Sort/src/MergeSort2"
	"practice/Play-with-Algorithms/03-Sorting-Advance/05-Quick-Sort/src/QuickSort"
)

// 比较Merge Sort和Quick Sort两种排序算法的性能效率
// 两种排序算法虽然都是O(nlogn)级别的, 但是Quick Sort算法有常数级的优势
// Quick Sort要比Merge Sort快, 即使我们对Merge Sort进行了优化
func main() {
	N := 1000000;

	// 测试1 一般测试
	fmt.Println("Test for random array, size =", N, ", random range [0, ", N, "]")

	arr1 := helper.GenerateRandomArray(N, 0, N)
	arr2 := make([]int, N)
	copy(arr2, arr1)

	helper.TestSort(MergeSort2.MergeSort, arr1)
	helper.TestSort(QuickSort.Sort, arr2)

	// 测试2 测试近乎有序的数组
	// 但是对于近乎有序的数组, 我们的快速排序算法退化成了O(n^2)级别的算法
	// 思考一下为什么对于近乎有序的数组, 快排退化成了O(n^2)的算法? :)
	swapTimes := 100

	fmt.Println("Test for nearly ordered array, size= ", N, " swap time = ", swapTimes)

	arr1 = helper.GenerateNearlyOrderedArray(N, swapTimes)
	copy(arr2, arr1)

	helper.TestSort(MergeSort2.MergeSort, arr1)
	helper.TestSort(QuickSort.Sort, arr2)
}
