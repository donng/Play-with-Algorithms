package main

import "fmt"

func selectionSort(arr []int, n int) []int {
	for i:= 0; i < n - 1; i++ {
		// 寻找 [i, n) 区间里的最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex]  {
				minIndex = j
			}
		}
		temp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = temp
	}
	return arr
}

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(selectionSort(arr, len(arr)))
}
