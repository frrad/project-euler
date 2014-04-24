package main

import (
	"euler"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	factors := euler.Factors(600851475143)

	fmt.Println(factors[len(factors)-1][0])

	fmt.Println("Elapsed time:", time.Since(starttime))
}
