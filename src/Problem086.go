package main

import (
	"euler"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	var total, M int

	for M = 1; total < 1000000; M++ {

		total = 0
		for a := 1; a <= M; a++ {
			for b := a; b <= M; b++ {
				for c := b; c <= M; c++ {

					best := (a+b)*(a+b) + (c * c)

					if euler.IsSquare(int64(best)) {
						total++
					}

				}
			}
		}

		fmt.Println(M, total)
	}
	fmt.Println(M - 1)
	fmt.Println("Elapsed time:", time.Since(starttime))
}
