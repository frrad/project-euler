package main

import (
	"fmt"
	"time"
)

func memoize(f func(int) int) func(int) int {

	memo := make(map[int]int)

	return func(n int) int {

		if answer, ok := memo[n]; ok {
			return answer
		}

		answer := f(n)
		memo[n] = answer
		return answer
	}

}

func fibonacci(n int) int {

	if n <= 1 {
		return 1
	}

	return (fibonacci(n-1) + fibonacci(n-2))

}

func main() {
	starttime := time.Now()

	fib := memoize(fibonacci)

	for j := 0; j < 10; j++ {

		for i := 0; i < 40; i++ {
			fmt.Println(fib(i))

		}

	}
	fmt.Println("Elapsed time:", time.Since(starttime))

}
