package main

import (
	"fmt"
	"time"
)

const (
	top         = 50000000
	maxA        = top/4 + 1
	desiredMult = 1
)

//If the sequence is x, x-a, x-2a then we have
//x =  3 a - Sqrt[4 a^2 - n] and x = 3 a + Sqrt[4 a^2 - n]
//Let c = Sqrt[4 a^2 - n]

func main() {
	starttime := time.Now()

	count := make(map[int64]int)

	for a := int64(1); a < maxA; a++ {

		for c := 2*a - 1; c >= 0; c-- {

			n := 4*a*a - c*c

			if n > top {
				break
			}

			if a-c > 0 {
				count[n]++
			}
			if a+c > 0 && c != 0 {
				count[n]++

			}
		}

	}

	total := 0
	for _, mult := range count {

		if mult == desiredMult {

			total++
		}
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
