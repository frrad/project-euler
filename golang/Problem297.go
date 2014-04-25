package main

import (
	"euler"
	"fmt"
	"time"
)

const top = 100000000000000000

//returns the largest Fibonacci number <= n
func floor(n int) (i int) {
	for i = 1; euler.Fibonacci(i) <= int64(n); i++ {
	}
	return int(euler.Fibonacci(i - 1))
}

var memo = map[int]int{0: 0}

func interval(x int) int {
	if ans, ok := memo[x]; ok {
		return ans
	}

	recurse := floor(x)
	remainder := x - recurse

	if remainder == 0 {
		memo[x] = 1 + interval(x-1)
		return interval(x)
	}

	return interval(recurse) + interval(remainder) + remainder
}

func main() {
	starttime := time.Now()

	fmt.Println(interval(top - 1))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
