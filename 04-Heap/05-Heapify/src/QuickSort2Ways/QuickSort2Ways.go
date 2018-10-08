package QuickSort2Ways

import (
	"practice/Play-with-Algorithms/03-Sorting-Advance/06-Quick-Sort-Deal-With-Nearly-Ordered-Array/src/InsertionSort"
	"math/rand"
)

// 双路快速排序的partition
// 返回p, 使得arr[l...p-1] < arr[p] ; arr[p+1...r] > arr[p]
func partition(arr []int, l int, r int) int {
	// 随机在arr[l...r]的范围中, 选择一个数值作为标定点pivot
	swap(arr, l, rand.Int()%(r-l+1)+l);
	v := arr[l]

	// arr[l+1...i] <= v; arr(j...r] >= v
	i := l + 1
	j := r

	for {
		// 注意这里的边界, arr[i].compareTo(v) < 0, 不能是arr[i].compareTo(v) <= 0
		// 思考一下为什么?
		for i <= r && arr[i] < v {
			i++
		}
		// 注意这里的边界, arr[j].compareTo(v) > 0, 不能是arr[j].compareTo(v) >= 0
		// 思考一下为什么?
		for j >= l+1 && arr[j] > v {
			j--
		}
		// 对于上面的两个边界的设定, 有的同学在课程的问答区有很好的回答:)
		// 大家可以参考: http://coding.imooc.com/learn/questiondetail/4920.html
		if i > j {
			break
		}

		swap(arr, i, j)
		i++
		j--
	}

	swap(arr, l, j)

	return j
}

// 递归使用快速排序,对arr[l...r]的范围进行排序
func quickSort(arr []int, l int, r int) {
	// 对于小规模数组, 使用插入排序
	if r-l <= 15 {
		InsertionSort.SortForMerge(arr, l, r)
		return
	}
	//if l >= r {
	//	return
	//}

	p := partition(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

func Sort(arr []int, n int) {
	quickSort(arr, 0, n-1)
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
