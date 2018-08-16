package main

import (
	"practice/Play-with-Algorithms/03-Sorting-Advance/04-Merge-Sort-Bottom-Up/src/helper"
	"fmt"
	"practice/Play-with-Algorithms/03-Sorting-Advance/04-Merge-Sort-Bottom-Up/src/MergeSort2"
	"practice/Play-with-Algorithms/03-Sorting-Advance/04-Merge-Sort-Bottom-Up/src/MergeSortBU"
)

// 比较Merge Sort和Merge Sort Bottom Up两种排序算法的性能效率
// 整体而言, 两种算法的效率是差不多的。但是如果进行仔细测试, 自底向上的归并排序会略胜一筹。
// 更详细的测试, 可以参考课程的这个问题: http://coding.imooc.com/learn/questiondetail/3208.html
// 本章节的代码仓也会给出更详细的测试代码
func main() {
	N := 1000000

	// 测试1 一般测试
	fmt.Println("Test for random array, size =", N, ", random range [0, ", N, "]")

	arr1 := helper.GenerateRandomArray(N, 0, N)
	arr2 := make([]int, N)
	copy(arr2, arr1)

	helper.TestSort(MergeSort2.MergeSort, arr1)
	helper.TestSort(MergeSortBU.Sort, arr2)

	// 测试2 测试近乎有序的数组
	swapTimes := 10

	fmt.Println("Test for nearly ordered array, size= ", N, " swap time = ", swapTimes)

	arr1 = helper.GenerateNearlyOrderedArray(N, swapTimes)
	copy(arr2, arr1)

	helper.TestSort(MergeSort2.MergeSort, arr1)
	helper.TestSort(MergeSortBU.Sort, arr2)
}
