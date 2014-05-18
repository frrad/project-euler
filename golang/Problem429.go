package main

import (
	"euler"
	"fmt"
	"time"
)

const (
	mod     = 1000000009
	top     = 100000000
	nPrimes = 5761455 //PrimePi[100000000]
)

func exp(a, b int) int {
	ans := int64(1)
	for i := 0; i < b; i++ {
		ans *= int64(a)
		ans %= mod
	}
	return int(ans)
}

func main() {
	starttime := time.Now()

	euler.PrimeCache(top)

	fmt.Println("There are", nPrimes, "primes to consider")

	facVec := make([][2]int, nPrimes)

	for i := 1; i <= nPrimes; i++ {
		p := int(euler.Prime(int64(i)))
		facVec[i-1][0] = p

		pk := p

		for pk <= top {
			facVec[i-1][1] += top / pk
			pk *= p
		}

		//since we want squares
		facVec[i-1][1] *= 2
	}

	answer := int64(1)
	for i := 0; i < len(facVec); i++ {
		p, k := facVec[i][0], facVec[i][1]

		addend := 1
		addend += exp(p, k)
		addend %= mod
		answer *= int64(addend)
		answer %= mod
		if (i+1)%1483 == 0 {
			fmt.Printf("\r%7d/%7d", i+1, len(facVec))
		}
	}

	fmt.Printf("\n%d\n", answer)

	fmt.Println("Elapsed time:", time.Since(starttime))

}
