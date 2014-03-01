package main

import (
	"euler"
	"fmt"
	"math"
	"time"
)

func sigma(n uint64) uint64 {
	var total uint64
	for i := uint64(1); i <= n; i++ {
		if n%i == 0 {
			total += i * i
		}
	}
	return total
}

func isSquare(n uint64) bool {
	sqrt := uint64(math.Sqrt(float64(n)))
	if sqrt*sqrt == n {
		return true
	}
	return false
}

func main() {
	starttime := time.Now()

	count := 1

	for i := uint64(2); i < 100000000; i++ {
		if isSquare(sigma(i)) {
			sqrt := int64(math.Sqrt(float64(sigma(i))))
			fmt.Println(count, ":",
				i, "=", euler.Factors(int64(i)), "\t",
				sigma(i), "=", sqrt, euler.Factors(sqrt))
			count++
		}
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
