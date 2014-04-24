package main

import (
	"fmt"
	"time"
)

const n = 5000
const target = -200000000299
const errorTarget = .2

func in(r float64) (out float64) {
	out = (300 - n) * r

	for i := 300 - n + 1; i <= 300-2; i++ {
		out = (out + float64(i)) * r

	}

	return
}

func s(r float64) float64 {
	return 897 + 3*in(r)
}

func main() {
	starttime := time.Now()

	err := errorTarget + 1.
	a, b, mid := .9980, 1.14, 0.

	for err > errorTarget || err < -1*errorTarget {
		mid = (b + a) / 2.

		err = in(mid) - target

		if err > 0 {
			a = mid
		} else {
			b = mid
		}

	}

	fmt.Println(mid)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
