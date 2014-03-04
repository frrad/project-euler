package main

import (
	"fmt"
	"time"
)

const top int = 10000

func divisorSum(n int) (sum int) {
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func amicableQ(a int) bool {
	b := divisorSum(a)
	if a == b {
		return false
	}
	if divisorSum(b) == a {
		return true
	}
	return false
}

func main() {
	starttime := time.Now()

	total := 0

	for i := 1; i < top; i++ {
		if amicableQ(i) {
			total += i
		}
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
