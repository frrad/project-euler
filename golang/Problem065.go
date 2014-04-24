package main

import (
	"euler"
	"fmt"
	"math/big"
	"time"
)

func ctdFrac(list []int) (num string, den string) {
	current := big.NewRat(int64(list[len(list)-1]), 1)

	for i := len(list) - 2; i >= 0; i-- {
		contrib := big.NewRat(int64(list[i]), 1)

		current.Inv(current)
		current.Add(current, contrib)

		fmt.Printf("%s\n", current)

	}

	return current.Num().String(), current.Denom().String()
}

func eList(n int) []int {
	answer := make([]int, n)

	answer[0] = 2

	for i := 1; i < n; i++ {
		answer[i] = 1
	}

	for i := 0; 3*i+2 < n; i++ {
		answer[3*i+2] = 2 * (i + 1)

	}

	return answer
}

func main() {
	starttime := time.Now()

	numerator, _ := ctdFrac(eList(100))
	fmt.Println(euler.StringDigitSum(numerator))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
