package main

import (
	"./euler"
	"fmt"
	"math"
)

const threshold = .0000000000001

func solveBottom(x float64, n float64) float64 {
	return (1.0 / 2.0) * (n - math.Sqrt((n*n)+(4*n*x)-(4*x*x)))
}

func naive(n int64) int64 {
	total := int64(0)

	N := float64(n)
	for x := math.Ceil(N / 2); x < N/2+(N*math.Sqrt(2)/2); x++ {

		y := solveBottom(x, N)
		yround := float64(int(y))

		if math.Abs(y-yround) < threshold {
			total++
		}
	}

	return total * 4
}

func isPythagorean(p int64) bool {
	return (p-1)%4 == 0
}

//4 more than 8 times the number of pythagorean primes
//contained in the prime factorization
func smart(n int64) int64 {
	factors := euler.Factor(n)
	total := int64(0)

	last := int64(0)

	distinct := int64(0)
	dupes := int64(0)

	for _, factor := range factors {
		if isPythagorean(factor) {
			total++

			if last == factor {
				dupes++
			} else {
				distinct++
			}

			last = factor

		}

	}

	answer := 8 * total * distinct
	if distinct > 1 {
		answer += 8 * dupes
	}
	return answer

}

func main() {

	badnumber := 0

	for i := int64(1); i < 10000; i++ {

		if naive(i)-smart(i)-4 != 0 {
			fmt.Println(i, naive(i)-4, smart(i), euler.Factor(i))

			badnumber++
		}
	}

	fmt.Println(badnumber)

}
