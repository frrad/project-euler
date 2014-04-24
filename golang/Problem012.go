package main

import (
	"euler"
	"fmt"
	"time"
)

const target = 500

func main() {
	starttime := time.Now()

	var tri, divisors int64

	for i := 1; divisors < target; i++ {

		tri = euler.TriangleNumber(i)
		divisors = euler.Divisors(tri)

	}

	fmt.Println(tri)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
