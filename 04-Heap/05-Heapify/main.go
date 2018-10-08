package main

import (
	"practice/Play-with-Algorithms/04-Heap/05-Heapify/src/MaxHeap"
	"fmt"
	"practice/Play-with-Algorithms/03-Sorting-Advance/08-Quick-Sort-Three-Ways/src/QuickSort3Ways"
	"practice/Play-with-Algorithms/04-Heap/05-Heapify/src/helper"
	"practice/Play-with-Algorithms/04-Heap/05-Heapify/src/MergeSort2"
	"practice/Play-with-Algorithms/04-Heap/05-Heapify/src/QuickSort2Ways"
)

func heapSort1(arr []int, n int)  {
	maxHeap := MaxHeap.GetMaxHeap(n)
	for i := 0; i < n; i++ {
		maxHeap.Insert(arr[i])
	}

	for i := n-1; i >= 0; i-- {
		arr[i] = maxHeap.ExtractMax()
	}
}

func heapSort2(arr []int, n int)  {
	maxHeap := MaxHeap.GetSortedMaxHeap(arr, n)
	for i := n-1; i >= 0; i-- {
		arr[i] = maxHeap.ExtractMax()
	}
}

func main() {
	N := 1000000

	// 测试1 一般测试
	fmt.Println("Test for random array, size =", N, ", random range [0, ", N, "]")

	arr1 := helper.GenerateRandomArray(N, 0, N)
	arr2 := make([]int, N)
	arr3 := make([]int, N)
	arr4 := make([]int, N)
	arr5 := make([]int, N)
	copy(arr2, arr1)
	copy(arr3, arr1)
	copy(arr4, arr1)
	copy(arr5, arr1)

	helper.TestSort(MergeSort2.MergeSort, arr1)
	helper.TestSort(QuickSort2Ways.Sort, arr2)
	helper.TestSort(QuickSort3Ways.Sort, arr3)
	helper.TestSort(heapSort1, arr4)
	helper.TestSort(heapSort2, arr5)

	// 测试2 测试近乎有序的数组
	// 但是对于近乎有序的数组, 我们的快速排序算法退化成了O(n^2)级别的算法
	// 思考一下为什么对于近乎有序的数组, 快排退化成了O(n^2)的算法? :)
	swapTimes := 100

	fmt.Println("Test for nearly ordered array, size= ", N, " swap time = ", swapTimes)

	arr1 = helper.GenerateNearlyOrderedArray(N, swapTimes)
	copy(arr2, arr1)
	copy(arr3, arr1)
	copy(arr4, arr1)
	copy(arr5, arr1)

	helper.TestSort(MergeSort2.MergeSort, arr1)
	helper.TestSort(QuickSort2Ways.Sort, arr2)
	helper.TestSort(QuickSort3Ways.Sort, arr3)
	helper.TestSort(heapSort1, arr4)
	helper.TestSort(heapSort2, arr5)

	// 测试3 很多重复值的数据
	fmt.Println("Test for random array, size =", N, ", random range [0, ", N, "]")

	arr1 = helper.GenerateRandomArray(N, 0, 10)
	arr2 = make([]int, N)
	copy(arr2, arr1)
	copy(arr3, arr1)
	copy(arr4, arr1)
	copy(arr5, arr1)

	helper.TestSort(MergeSort2.MergeSort, arr1)
	helper.TestSort(QuickSort2Ways.Sort, arr2)
	helper.TestSort(QuickSort3Ways.Sort, arr3)
	helper.TestSort(heapSort1, arr4)
	helper.TestSort(heapSort2, arr5)
}
