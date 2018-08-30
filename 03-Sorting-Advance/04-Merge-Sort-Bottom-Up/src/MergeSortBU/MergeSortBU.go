package MergeSortBU

import "math"

// 将 arr[l...mid] 和 arr[mid+1...r] 两部分归并
func merge(arr []int, l int, mid int, r int) {
	aux := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		aux[i-l] = arr[i]
	}

	// 初始化，i指向左半部分的起始索引位置l；j指向右半部分起始索引位置mid+1
	i := l
	j := mid + 1
	for k := l; k <= r; k++ {
		if i > mid { // 如果左半部分元素已经全部处理完毕
			arr[k] = aux[j-l]
			j++
		} else if j > r { // 如果右半部分元素已经全部处理完毕
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] { // 左半部分所指元素 < 右半部分所指元素
			arr[k] = aux[i-l]
			i++
		} else { // 左半部分所指元素 >= 右半部分所指元素
			arr[k] = aux[j-l]
			j++
		}
	}
}

func Sort(arr []int, n int) {
	// Merge Sort Bottom Up 无优化版本
	for size := 1; size < n; size *= 2 {
		for i := 0; i < n-size; i += size * 2 {
			// 对 arr[i...i+sz-1] 和 arr[i+sz...i+2*sz-1] 进行归并
			merge(arr, i, i+size-1, int(math.Min(float64(i+size*2-1), float64(n-1))))
		}
	}

	// Merge Sort Bottom Up 优化
	// 对于小数组, 使用插入排序优化
}
