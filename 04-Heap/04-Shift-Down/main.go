package main

import (
	"math/rand"
	"fmt"
	"practice/Play-with-Algorithms/04-Heap/04-Shift-Down/src/MaxHeap"
)

func main() {
	maxHeap := MaxHeap.GetMaxHeap(100)

	for i := 0; i < 100; i++ {
		maxHeap.Insert(rand.Int()%100)
	}

	for !maxHeap.IsEmpty() {
		fmt.Print(maxHeap.ExtractMax(), " ")
	}
}
