package main

import (
	"fmt"

	"github.com/kristain09/earn-data-structures/structures"
)

func main() {
	m := &structures.MaxHeap{}

	buildHeap := []int{10, 20, 30, 2, 4, 56, 84, 97, 64, 15, 21}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i <= 5; i++ {
		m.Extract()
		fmt.Println(m)
	}

	
}
