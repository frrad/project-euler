package main

import (
	"fmt"
	"math"
	"time"
)

func f(x float64) float64 {
	return math.Floor(math.Pow(2, 30.403243784-(x*x))) / 1000000000
}

func main() {
	starttime := time.Now()

	start := -1.
	seen := make(map[float64]bool)

	for !seen[start] {

		seen[start] = true

		start = f(start)

	}

	fmt.Println(start + f(start))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
