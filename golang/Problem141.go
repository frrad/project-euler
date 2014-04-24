package main

import (
	"euler"
	"fmt"
	"time"
)

const top int64 = 1000000000000

func type2(a, b int64) int64 {
	return b*b*b/a + a
}

func c(a, b int64) int64 {
	return b * b / a
}

func main() {
	starttime := time.Now()

	cans := make(map[int64]bool)

	a, b := int64(1), int64(1)

	sum := int64(0)

	for a = 1; type2(a, a) < top; a++ {
		for b = a; type2(a, b) < top; b++ {

			if (b*b)%a == 0 {

				if euler.IsSquare(type2(a, b)) {

					if _, ok := cans[type2(a, b)]; !ok {
						cans[type2(a, b)] = true
						sum += type2(a, b)
					}

					fmt.Println(a, b, c(a, b), type2(a, b), "\t", sum)

				}
			}

		}
	}

	fmt.Println(sum)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
