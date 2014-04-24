package main

import (
	"fmt"
	"time"
)

var memo map[[2]int]int64

func routes(a, b int) int64 {
	if a == 0 || b == 0 {
		return 1
	}

	if answer, ok := memo[[2]int{a, b}]; ok {
		return answer
	}

	answer := int64(0)

	for A := 0; A <= a-1; A++ {
		answer += routes(A, b-1)
	}

	for B := 0; B <= b-1; B++ {
		answer += routes(a-1, B)
	}
	memo[[2]int{a, b}] = answer
	return answer

}

func main() {
	starttime := time.Now()

	memo = make(map[[2]int]int64)

	fmt.Println(routes(20, 20))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
