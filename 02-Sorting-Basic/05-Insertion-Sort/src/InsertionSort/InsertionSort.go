package InsertionSort

import (
	"practice/Play-with-Algorithms/02-Sorting-Basic/05-Insertion-Sort/src/helper"
)

func Sort(arr []int, length int) {
	for i := 0; i < length-1; i++ {
		// 写法 1
		/*for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				swap(arr, j-1, j)
			} else {
				break
			}
		}*/

		// 写法 2
		for j := i + 1; j > 0 && arr[j] < arr[j-1]; j-- {
			swap(arr, j-1, j)
		}
	}
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 测试InsertionSort
func main() {
	N := 20000
	arr := helper.GenerateRandomArray(N, 0, 100000)
	helper.TestSort(Sort, arr)
}
