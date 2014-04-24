package main

import (
	"euler"
	"fmt"
	"time"
)

const perimeter = 100000000

func main() {
	starttime := time.Now()

	total := int64(0)

	for n := int64(1); 2*n*n < perimeter; n++ {

		for m := n + 1; 2*m*m+n*n < perimeter; m += 2 {

			if euler.GCD(n, m) != 1 {
				continue
			}

			a, b, c := m*m-n*n, 2*m*n, m*m+n*n

			if a > b {
				a, b = b, a
			}

			if c%(b-a) == 0 {
				total += perimeter / (a + b + c)
			}

		}

	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
