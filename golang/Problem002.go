package main

import (
	"euler"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	sum := 0
	for i := 1; euler.Fibonacci(i) < 4000000; i++ {
		if fib := euler.Fibonacci(i); fib%2 == 0 {
			sum += int(fib)
		}
	}

	fmt.Println(sum)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
