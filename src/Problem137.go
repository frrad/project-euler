package main

import (
	//	"euler"
	"fmt"
	"time"
)

var memo map[int]int64

func fib(index int) int64 {
	if index <= 2 {
		return 1
	}

	if answer, ok := memo[index]; ok {
		return answer
	}

	answer := fib(index-1) + fib(index-2)

	memo[index] = answer
	return answer
}

func main() {
	starttime := time.Now()
	memo = make(map[int]int64)

	i := 15

	fmt.Println(fib(2*i) * fib(2*i+1))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
