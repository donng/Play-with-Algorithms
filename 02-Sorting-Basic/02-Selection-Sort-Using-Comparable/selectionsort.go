package main

import (
	"fmt"
	"github.com/donng/Play-with-Data-Algorithms/02-Sorting-Basic/02-Selection-Sort-Using-Comparable/student"
	"sort"
)

func SelectionSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		// 寻找 [i, n) 区间里的最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swap(arr, i, minIndex)
	}
}

// Students 实现了原生 sort 接口
func Sort(s sort.Interface) {
	length := s.Len()
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if s.Less(j, i) {
				minIndex = j
			}
		}
		s.Swap(minIndex, i)
	}
}

// slice 是引用类型
func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	SelectionSort(arr, len(arr))

	fmt.Println(arr)

	students := student.Students{}

	students = append(students, student.New("D", 90))
	students = append(students, student.New("C", 100))
	students = append(students, student.New("B", 95))
	students = append(students, student.New("A", 95))
	Sort(students)

	for _, s := range students {
		fmt.Println(s)
	}
}
