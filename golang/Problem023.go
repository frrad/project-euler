package main

import (
	"euler"
	"fmt"
	"time"
)

const top int = 28123

func isAbundant(n int) bool {
	return euler.DivisorSigma(int64(n), 1) > int64(2*n)
}

func main() {
	starttime := time.Now()

	abundant := make([]int, 0)

	for x := 10; x < top; x++ {
		if isAbundant(x) {
			abundant = append(abundant, x)
		}
	}

	doubles := make(map[int]bool)

	for j, number := range abundant {

		for i := j; abundant[i]+number <= top; i++ {
			doubles[number+abundant[i]] = true
		}
	}

	var answer int

	for i := 0; i < top; i++ {
		if !doubles[i] {
			answer += i
		}
	}

	fmt.Println(answer)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
