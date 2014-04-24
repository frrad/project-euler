package main

import (
	"euler"
	"fmt"
	"time"
)

const points = 20

func prob(x int, q float64) float64 {
	return 1. - float64(x)/q
}

func search(f func(float64) float64, a, b, target, epsilon float64) float64 {

	for b-a > epsilon {
		fmt.Println(a, "---", b)

		mid := (a + b) / 2

		if f(mid) < target {
			b = mid
		} else {
			a = mid
		}

	}

	return a

}

func transition(x int, q float64) [][]float64 {
	matrix := make([][]float64, points+1)

	for i := 0; i < points+1; i++ {
		matrix[i] = make([]float64, points+1)

		for j := 0; j < points+1; j++ {
			if i-j == 0 {
				matrix[i][j] = 1 - prob(x, q)
			}
			if i-j == 1 {
				matrix[i][j] = prob(x, q)
			}

		}
	}

	return matrix
}

func twenty(q float64) float64 {

	state := transition(1, q)

	for x := 2; x <= 50; x++ {
		state = euler.MatrixProd(transition(x, q), state)
	}

	return state[points][0]

}

func main() {
	starttime := time.Now()

	answer := search(twenty, 50, 60, .02, .00000000001)
	fmt.Println(answer, twenty(answer))

	fmt.Printf("%.10f\n", answer)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
