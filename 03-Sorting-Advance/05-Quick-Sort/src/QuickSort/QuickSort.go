package QuickSort

// 递归使用快速排序,对arr[l...r]的范围进行排序
func quickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}

	p := partition(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

// 对arr[l...r]部分进行partition操作
// 返回p, 使得arr[l...p-1] < arr[p] ; arr[p+1...r] > arr[p]
func partition(arr []int, l int, r int) int {
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < arr[l] {
			j++
			swap(arr, i, j)
		}
	}
	swap(arr, l, j)

	return j
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func Sort(arr []int, n int) {
	quickSort(arr, 0, n-1)
}
