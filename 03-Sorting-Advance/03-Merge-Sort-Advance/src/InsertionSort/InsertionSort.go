package InsertionSort

func Sort(arr []int, length int) {
	for i := 1; i < length; i++ {
		// 寻找元素 arr[i] 合适的插入位置
		temp := arr[i]
		// j 保存元素 temp 应该插入的位置
		var j int
		for j = i; j > 0 && temp < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}

func SortForMerge(arr []int, l int, r int) {
	for i := l; i <= r; i++ {
		temp := arr[i]
		var j int
		for j = i; j > l && arr[j-1] > temp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}
