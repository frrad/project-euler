package main

import (
	"fmt"
	"math"
	"time"
)

const threshold = .0000000000001

func solveBottom(x float64, n float64) float64 {
	return (1.0 / 2.0) * (n - math.Sqrt((n*n)+(4*n*x)-(4*x*x)))
}

func naive(n int64) int64 {
	total := int64(0)
	N := float64(n)
	for x := math.Ceil(N / 2); x < N/2+(N*math.Sqrt(2)/2); x++ {
		y := solveBottom(x, N)
		yround := float64(int(y))
		if math.Abs(y-yround) < threshold {
			total++
		}
	}
	return total * 4
}

func better(n int64) int64 {

	count := int64(0)
	for i := int64(0); i < n; i++ {
		a2 := n*n - i*i
		a := int64(math.Sqrt(float64(a2)))
		if a2 == a*a {
			count++
		}
	}
	return 4 * count

}

func main() {

	start := time.Now()

	top := int64(0)
	for i := int64(0); i < 10000; i++ {

		if better(i) > top {
			fmt.Println(better(i))
			top = better(i)

		}

	}

	fmt.Println("Elapsed time:", time.Since(start))

}
