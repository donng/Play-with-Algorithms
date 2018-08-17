package QuickSort3Ways

import (
	"practice/Play-with-Algorithms/03-Sorting-Advance/08-Quick-Sort-Three-Ways/src/InsertionSort"
	"math/rand"
)

// 递归使用快速排序,对arr[l...r]的范围进行排序
func quickSort(arr []int, l int, r int) {
	// 对于小规模数组, 使用插入排序
	if r-l <= 15 {
		InsertionSort.SortForMerge(arr, l, r)
		return
	}

	// 随机在arr[l...r]的范围中, 选择一个数值作为标定点pivot
	swap(arr, l, rand.Int()%(r-l+1)+l);
	v := arr[l]

	lt := l      // arr[l+1...lt] < v
	gt := r + 1  // arr[gt...r] > v
	i := l + 1   // arr[lt+1...i) == v
	for i < gt {
		if arr[i] < v {
			swap(arr, i, lt+1)
			lt++
			i++
		} else if arr[i] > v {
			swap(arr, i, gt-1)
			gt--
		} else {
			i++
		}
	}

	swap(arr, l, lt)

	quickSort(arr, l, lt-1)
	quickSort(arr, gt, r)
}

func Sort(arr []int, n int) {
	quickSort(arr, 0, n-1)
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 三路快速排序算法也是一个O(nlogn)复杂度的算法
// 可以在1秒之内轻松处理100万数量级的数据
