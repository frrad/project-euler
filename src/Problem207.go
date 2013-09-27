package main

import (
	"euler"
	"fmt"
	"math"
	"time"
)

const (
	epsilon  = .0001
	num, den = 1, 12345
)

func main() {
	starttime := time.Now()

	perfect, total := 0, 0
	k := int64(0)

	for base := int64(1); ; base++ {

		k = 4*base*base - 6*base + 2

		if sqrt, ok := euler.IntSqrt(1 + 4*k); ok && k > 0 {
			total++
			t := math.Log2(.5 + .5*float64(sqrt))
			if math.Abs(float64(int(t))-t) < epsilon {
				perfect++
				fmt.Println(perfect, total, k)
			}

			if den*perfect < num*total {
				fmt.Println(perfect, total, k)
				break
			}
		}

		k = 4*base*base - 2*base

		if sqrt, ok := euler.IntSqrt(1 + 4*k); ok && k > 0 {
			total++
			t := math.Log2(.5 + .5*float64(sqrt))
			if math.Abs(float64(int(t))-t) < epsilon {
				perfect++
				fmt.Println(perfect, total, k)
			}

			if den*perfect < num*total {
				fmt.Println(perfect, total, k)
				break
			}
		}

	}

	fmt.Println(k)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
