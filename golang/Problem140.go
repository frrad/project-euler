package main

import (
	"fmt"
	"time"
)

const height = 30

//Note: Brute force is too slow, solved Diophantine equation
//5 x^2 + 14 x + 1 == y^2 with Mathematica
//By math x is FGN iff 5 x^2 + 14 x + 1 is a perfect square
func main() {
	starttime := time.Now()

	square := int64(1)
	i := int64(1)
	count := 0
	total := int64(0)

	for x := int64(1); count < height; x++ {
		test := 1 + 14*x + 5*x*x

		for square < test {
			square += (2 * i) + 1
			i++
		}

		if test == square {
			fmt.Println(count+1, "\t", x)
			count++
			total += x
		}

	}

	fmt.Println("==========")
	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
