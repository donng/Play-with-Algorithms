package BubbleSort

func Sort(arr []int, length int) {
	var swapped bool
	for {
		swapped = false
		for i := 1; i < length; i++ {
			if arr[i] < arr[i-1] {
				swap(arr, i-1, i)
				swapped = true
			}
		}
		length--
		// 循环中没有数据交换，说明已正序，停止循环
		if !swapped {
			break
		}
	}
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
