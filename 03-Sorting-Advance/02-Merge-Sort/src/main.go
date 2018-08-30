package main

import (
	"practice/Play-with-Algorithms/03-Sorting-Advance/02-Merge-Sort/src/helper"
	"practice/Play-with-Algorithms/03-Sorting-Advance/02-Merge-Sort/src/InsertionSort"
	"practice/Play-with-Algorithms/03-Sorting-Advance/02-Merge-Sort/src/MergeSort"
)

func main() {
	// Merge Sort是我们学习的第一个O(nlogn)复杂度的算法
	// 可以在1秒之内轻松处理100万数量级的数据
	// 注意：不要轻易尝试使用SelectionSort, InsertionSort或者BubbleSort处理100万级的数据
	// 否则，你就见识了O(n^2)的算法和O(nlogn)算法的本质差异：）
	N := 50000
	arr := helper.GenerateRandomArray(N, 0, N)

	helper.TestSort(InsertionSort.Sort, arr)
	helper.TestSort(MergeSort.MergeSort, arr)
}
