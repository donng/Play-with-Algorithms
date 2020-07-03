package bubblesort2

func Sort(arr []int, length int) {
	var newn int
	for {
		newn = 0
		for i := 1; i < length; i++ {
			if arr[i] < arr[i-1] {
				swap(arr, i-1, i)
				newn = i
			}
		}
		// 最后一次交换的点，意味着后面的值都正序，不需要再遍历
		length = newn

		// for 循环结束没有数值交换说明已正序
		if newn == 0 {
			break
		}
	}
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
