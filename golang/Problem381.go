package main

import (
	"euler"
	"fmt"
	"time"
)

const top = 100000000

func S(p int) int {
	ans := p

	first := int(euler.InverseMod(int64(p-2), int64(p)))
	ans += first

	delta := int(euler.InverseMod(int64(p-3), int64(p)))
	first = (first * delta) % p
	ans += first

	delta = int(euler.InverseMod(int64(p-4), int64(p)))
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
