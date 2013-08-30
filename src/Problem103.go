package main

import (
	"fmt"
	"time"
)

func pop(this [][]int, new int) (bool, [][]int) {
	all := make([]int, len(this[0]))
	copy(all, this[0])
	all = append(all, new)

	fmt.Println(all)

	for i, item := range this {
		for j, current := range item {
			this[i][j] = current + new
		}
	}

	return true, append(all, this)
}

func main() {
	starttime := time.Now()

	blob := [][]int{[]int{1, 2, 3}}

	pop(blob, 45)
	fmt.Println(blob)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
