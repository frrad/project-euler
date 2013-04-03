package main

import (
	"euler"
	"fmt"
	"time"
)

const (
	maxp   = 10
	target = 100
)

func next(n int64) int64 {
	var ratio float64
	var best int64
	for i := int64(1); i < maxp; i++ {
		test := euler.Prime(i) * n
		if float64(solns(test))/float64(test) > ratio {
			best = euler.Prime(i) * n
			ratio = float64(solns(test)) / float64(test)
		}
	}
	return best
}

func solns(n int64) int64 {
	return euler.Divisors(n*n)/2 + 1
}

func main() {
	starttime := time.Now()

	test := int64(2)

	for solns(test) < target {
		fmt.Println(test, solns(test))
		test = next(test)
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
