package main

import (
	"euler"
	"fmt"
	"time"
)

const target = 12

func main() {
	starttime := time.Now()

	count := 0
	answer := int64(0)

	for n := int64(1); count < target; n++ {

		MCans := make([]int64, 0)

		if sqrt, ok := euler.IntSqrt(5*n*n + 1); ok {
			MCans = append(MCans, 2*n+sqrt)
		}
		if sqrt, ok := euler.IntSqrt(5*n*n - 1); ok {
			MCans = append(MCans, 2*n+sqrt)
		}
		if sqrt, ok := euler.IntSqrt(5*n*n + 2); ok {
			MCans = append(MCans, n+sqrt)
		}
		if sqrt, ok := euler.IntSqrt(5*n*n - 2); ok {
			MCans = append(MCans, n+sqrt)
		}

		for _, m := range MCans {
			a, b, c := m*m-n*n, 2*m*n, m*m+n*n

			count++
			answer += c

			fmt.Println(a, b, c)
		}

	}

	fmt.Println("==========")
	fmt.Println(answer)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
