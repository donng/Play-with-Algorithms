package BubbleSort3

func Sort(arr []int, length int) {
	for i := 0; i < length-1; i++ {
		for j := 1; j < length-i; j++ {
			if arr[j] < arr[j-1] {
				swap(arr, j-1, j)
			}
		}
	}
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
