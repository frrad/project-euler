package main

import (
	"euler"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	//for unecessary speed, cache primes with seive:
	euler.PrimeCache(120000)

	fmt.Println(euler.Prime(10001))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
