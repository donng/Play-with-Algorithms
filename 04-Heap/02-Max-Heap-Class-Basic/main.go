package main

import (
	"practice/Play-with-Algorithms/04-Heap/02-Max-Heap-Class-Basic/src/MaxHeap"
	"fmt"
)

func main()  {
	maxHeap := MaxHeap.GetMaxHeap(100)
	fmt.Println(maxHeap.Size())
}
