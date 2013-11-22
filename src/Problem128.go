package main

import (
	"euler"
	"fmt"
	"time"
)

const target = 2000

func diff(an int, list [6]int) [6]int {
	for i := 0; i < 6; i++ {
		if an-list[i] > 0 {
			list[i] = an - list[i]
		} else {
			list[i] = list[i] - an
		}
	}
	return list
}

func PD(list [6]int) (count int) {
	for i := 0; i < 6; i++ {
		if euler.IsPrime(int64(list[i])) {
			count++
		}
	}
	return
}

func up(n int) int {
	return 3*n*n - 3*n + 2
}

func side(n int) int {
	return up(n) - 1
}

func oops(n int) [6]int {
	return [6]int{
		up(n + 1),
		side(n + 2),
		side(n + 1),
		up(n - 1),
		up(n) + 1,
		up(n+1) + 1,
	}
}

func soups(n int) [6]int {
	return [6]int{
		side(n + 1),
		side(n+1) - 1,
		side(n) - 1,
		side(n - 1),
		up(n - 2),
		up(n - 1),
	}
}

func main() {
	starttime := time.Now()

	var current, pd int

	seen := 2

	for i := 5; seen < target; i++ {
		if n := i / 2; i%2 == 0 {
			current = side(n)
			pd = PD(diff(side(n), soups(n)))
		} else {
			current = up(n)
			pd = PD(diff(up(n), oops(n)))
		}

		if pd == 3 {
			seen++
		}
	}

	fmt.Println(current)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
