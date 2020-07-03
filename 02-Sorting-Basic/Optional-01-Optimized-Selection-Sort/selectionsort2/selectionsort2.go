package selectionsort2

// 感谢github @zhengquan45 提出, 可以对选择排序进行优化
// 在每一轮中, 可以同时找到当前未处理元素的最大值和最小值
func Sort(arr []int, length int) {
	left := 0
	right := length - 1

	for left < right {
		minIndex := left
		maxIndex := right

		// 在每一轮查找时, 要保证arr[minIndex] <= arr[maxIndex]
		// left 和 right 在往中间移动时，会出现初始的 left值 > right值 的情况
		if arr[left] > arr[right] {
			swap(arr, left, right)
		}

		for i := left + 1; i < right; i++ {
			if arr[i] < arr[minIndex] {
				minIndex = i
			}
			if arr[i] > arr[maxIndex] {
				maxIndex = i
			}
		}
		swap(arr, left, minIndex)
		swap(arr, right, maxIndex)

		left++
		right--
	}
}

func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
