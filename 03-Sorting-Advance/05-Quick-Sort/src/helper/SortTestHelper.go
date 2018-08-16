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
		arr = append(arr, rand.Int()%(rangeR-rangeL+1)+rangeL)
	}

	return arr
}

// 生成一个近乎有序的数组
// 首先生成一个含有[0...n-1]的完全有序数组, 之后随机交换swapTimes对数据
// swapTimes定义了数组的无序程度:
// swapTimes == 0 时, 数组完全有序
// swapTimes 越大, 数组越趋向于无序
func GenerateNearlyOrderedArray(n int, swapTimes int) []int {
	arr := []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}
	rand.Seed(time.Now().Unix())
	for i := 0; i < swapTimes; i++ {
		posx := rand.Int()%n
		posy := rand.Int()%n
		arr[posx], arr[posy] = arr[posy], arr[posx]
	}
	return arr
}

// 测试sortClassName所对应的排序算法排序arr数组所得到结果的正确性和算法运行时间
func TestSort(f func([]int, int), arr []int) {
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
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] < arr[i] {
			return false
		}
	}
	return true
}
