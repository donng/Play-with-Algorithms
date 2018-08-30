package MergeSort

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
		if i > mid {  // 如果左半部分元素已经全部处理完毕
			arr[k] = aux[j-l]
			j++
		} else if j > r {  // 如果右半部分元素已经全部处理完毕
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {  // 左半部分所指元素 < 右半部分所指元素
			arr[k] = aux[i-l]
			i++
		} else {  // 左半部分所指元素 >= 右半部分所指元素
			arr[k] = aux[j-l]
			j++
		}
	}
}

// 递归使用归并排序，对 arr[l...r] 的范围进行排序
func mergeSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	// 拆分的过程
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	// 合并的过程
	merge(arr, l, mid, r)
}

func MergeSort(arr []int, n int) {
	mergeSort(arr, 0, n-1)
}
