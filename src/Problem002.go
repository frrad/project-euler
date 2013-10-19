package main

import (
	"fmt"
	"time"
)

var fibMemo map[int]int

func fib(n int) int {
	if ans, ok := fibMemo[n]; ok {
		return ans
	}

	fibMemo[n] = fib(n-1) + fib(n-2)
	return fib(n)
}

func main() {
	starttime := time.Now()

	fibMemo = make(map[int]int)
	fibMemo[0] = 0
	fibMemo[1] = 1

	sum := 0

	for i := 1; fib(i) < 4000000; i++ {
		if fib(i)%2 == 0 {
			sum += fib(i)
		}
	}

	fmt.Println(sum)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
