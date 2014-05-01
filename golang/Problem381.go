package main

import (
	"euler"
	"fmt"
	"time"
)

const top = 100000000

func inverseModP(n, p int) int {
	temp, _ := euler.ExtendedEuclidean(int64(n), int64(p))
	ans := int(temp)
	for ans > 0 {
		ans -= p
	}
	for ans < 0 {
		ans += p
	}
	return ans
}

func S(p int) int {
	ans := p

	first := inverseModP(p-2, p)
	ans += first

	delta := inverseModP(p-3, p)
	first = (first * delta) % p
	ans += first

	delta = inverseModP(p-4, p)
	first = (first * delta) % p

	ans += first

	return ans % p
}

func main() {
	starttime := time.Now()

	euler.PrimeCache(top)
	fmt.Println("Built table")

	total := 0
	for i, p := 4, 5; p < top; i, p = i+1, int(euler.Prime(int64(i))) {
		total += S(p)
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
