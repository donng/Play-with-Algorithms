package main

import (
	"practice/Play-with-Algorithms/04-Heap/03-Shift-Up/src/MaxHeap"
	"math/rand"
	"fmt"
)

func main() {
	maxHeap := MaxHeap.GetMaxHeap(100)

	for i := 0; i < 100; i++ {
		maxHeap.Insert(rand.Int()%100)
	}

	fmt.Println(maxHeap)
}
