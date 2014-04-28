package main

import (
	"euler"
	"fmt"
	"time"
)

const top = 40

func moves(i int) int64 {
	return int64(i*i + 2*i)
}

func main() {
	starttime := time.Now()

	Ti, Mi := 1, 1
	list := make([]int, 0)

	for len(list) < top {
		T, M := euler.TriangleNumber(Ti), moves(Mi)
		if T == M {
			list = append(list, Mi)
			fmt.Println(list, len(list))
		}

		if T > M {
			Mi++
		} else {
			Ti++
		}
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
