package main

import (
	"./euler"
	"fmt"
	"math"
	"strconv"
	"time"
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

//number of pythagorean triples with "C"=n
func better(n int64) int64 {

	count := int64(0)
	for i := int64(0); i < n; i++ {
		a2 := n*n - i*i
		if euler.IsSquare(a2) {
			count++
		}
	}
	return 4 * count

}

func main() {

	starttime := time.Now()

	//We're looking for numbers of the form 2^k * \prod pi * q1 * q2^2 * q3^3
	//for pi,qi prime and p_1 % 4 = 3 and qi%4 = 1. Start by populating tables
	primes1 := make([]string, 0) //Table of pythagorean primes
	primes3 := make([]string, 0) //Primes which are 3 mod 4

	for i := int64(0); i < 350000; i++ {
		num := euler.Prime(i)
		prime := strconv.FormatInt(num, 10)

		if num%4 == 1 {
			primes1 = append(primes1, prime)

		} else {
			primes3 = append(primes3, prime)
		}
	}

	for i := 0; i < count; i++ {

	}

	fmt.Println(primes1, primes3)

	fmt.Println("Elapsed time:", time.Since(starttime))

}
