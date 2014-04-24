package main

import (
	"fmt"

	"time"
)

func digits(n int) int {

	sum := 0

	for n > 0 {

		x := n % 10
		n /= 10

		sum += x * x * x * x * x
	}

	return sum
}

func main() {
	starttime := time.Now()

	total := 0

	for i := 2; i < 999999; i++ {
		if digits(i) == i {
			fmt.Println(i)
			total += i
		}
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
