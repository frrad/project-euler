package main

import (
	"euler"
	"fmt"
	"math"
	"time"
)

const top = 10000

func M(n int64) (term bool) {
	try := float64(n) / math.E
	try1 := float64(int64(try))
	try2 := try1 + 1

	test := math.Pow(try2/try1, try1)
	test *= try2 / float64(n)

	var den int64
	if test > 1 {
		den = int64(try1)
	} else {
		den = int64(try2)
	}

	for den%2 == 0 {
		den /= 2
	}

	for den%5 == 0 {
		den /= 5
	}

	gcd := euler.GCD(n, den)
	den /= gcd

	if den == 1 {
		term = true
	}

	return

}

func main() {
	starttime := time.Now()

	sum := 0

	for N := 5; N <= top; N++ {
		term := M(int64(N))

		if term {
			sum -= N
		} else {
			sum += N
		}
	}

	fmt.Println(sum)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
