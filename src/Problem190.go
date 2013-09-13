package main

import (
	"fmt"
	"time"
)

//Use Lagrange multipliers: j * x_i = i * x_j

func total(vec []float64) (sum float64) {
	for _, x := range vec {
		sum += x
	}
	return
}

func P(m int) float64 {
	vec := make([]float64, m)
	vec[0] = 1
	for i := 0; i < m; i++ {
		vec[i] = vec[0] * float64(i+1)
	}

	vec[0] = float64(m) / total(vec)

	for i := 0; i < m; i++ {
		vec[i] = vec[0] * float64(i+1)
	}

	answer := 1.
	for i := 0; i < m; i++ {
		for j := 0; j < i+1; j++ {
			answer *= vec[i]
		}
	}

	return answer
}

func main() {
	starttime := time.Now()

	sum := 0

	for i := 2; i <= 15; i++ {
		sum += int(P(i))
	}

	fmt.Println(sum)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
