package helper

import (
	"math/rand"
	"time"
	"fmt"
	"runtime"
	"reflect"
)

// 生成有n个元素的随机数组,每个元素的随机范围为[rangeL, rangeR]
func GenerateRandomArray(n int, rangeL int, rangeR int) []int {
	if rangeL > rangeR {
		panic("rangeL cannot be bigger than rangeR ")
	}

	arr := []int{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Int() % (rangeR - rangeL + 1) + rangeL)
	}

	return arr
}

// 测试sortClassName所对应的排序算法排序arr数组所得到结果的正确性和算法运行时间
func TestSort(f func([]int, int), arr []int)  {
	startTime := time.Now()
	f(arr, len(arr))

	if !isSorted(arr) {
		panic("sort failed.")
	}

	funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	fmt.Println(funcName, time.Now().Sub(startTime))
}

// 判断arr数组是否有序
func isSorted(arr []int) bool {
	for i := 0; i < len(arr) - 1; i++ {
		if arr[i + 1] < arr[i] {
			return false
		}
	}
	return true
}