package PrintableMaxHeap

import (
	"practice/Play-with-Algorithms/04-Heap/03-Shift-Up/src/MaxHeap"
	"fmt"
)

// todo 打印未完成
func testPrint(maxHeap *MaxHeap.MaxHeap)  {
	if maxHeap.Size() >= 100 {
		fmt.Println("This print function can only work for less than 100 integer")
		return
	}

	fmt.Println("The max heap size is: ", maxHeap.Size())
	fmt.Println("Data in the max heap: ")
	for i := 1; i <= maxHeap.Size(); i++ {
		// 我们的print函数要求堆中的所有整数在[0, 100)的范围内
	}
}