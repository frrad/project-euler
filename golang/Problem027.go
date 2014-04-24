package main

import (
	"euler"
	"fmt"
	"time"
)

func f(a, b, n int) int {
	return n*n + a*n + b
}

func length(a, b int) int {
	for n := 0; ; n++ {
		if f(a, b, n) < 0 {
			return n - 1
		}
		if !euler.IsPrime(int64(f(a, b, n))) {
			return n - 1
		}
	}

	return -100
}

func main() {
	starttime := time.Now()

	max := 0
	best := -1

	for a := -1000; a < 1000; a++ {
		for b := -1000; b < 1000; b++ {
			if euler.IsPrime(int64(f(a, b, max))) {
				if depth := length(a, b); depth > max {
					max = depth
					best = a * b
					fmt.Println(a, "*", b, "=", a*b, ":", max)
				}
			}

		}
	}

	fmt.Println(best)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
