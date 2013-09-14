package main

import (
	//	"euler"
	"fmt"
	"time"
)

var memo map[int]int64

func fib(index int) int64 {
	if index == 1 {
		return 1
	}

	if index == 2 {
		return 4
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

	for i := 1; i < 50; i++ {

		fmt.Print(fib(i), ",")

		//fmt.Println(fib(i) * fib(i+1))

	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
