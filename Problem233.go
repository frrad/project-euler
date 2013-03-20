package main

import (
	//	"./euler"
	"fmt"
	"math"
)

const threshold = .000000000001

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

func main() {

	max := int64(0)
	fmt.Print("\n{")

	for i := int64(1); ; i++ {
		if naive(i) > max {
			fmt.Print("{", i, ",", i, "},")
			max = naive(i)
		}
	}
	fmt.Print("\b}\n")

}
