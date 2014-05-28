package main

import (
	"fmt"
	"time"
)

const (
	tip  = 10000000000 //tip = 10^10
	top  = 49          //top = 1 + n /. Solve[tip == (.5 + .5 Sqrt[5])^n][[1]] // Ceiling
	reps = 22          //reps = n /. Solve[2*n + 5  == top][[1]] // Ceiling
)

var luke = map[int]int64{1: 1, 2: 3}

func lucas(n int) int64 {
	if ans, ok := luke[n]; ok {
		return ans
	}
	ans := lucas(n-1) + lucas(n-2)
	luke[n] = ans
	return ans
}

func check(n int64) bool {
	for n > 0 {
		if n&3 == 3 || n&5 == 5 {
			return false
		}
		n >>= 1
	}
	return true
}

func sum(i int64) (ans int64) {
	n := 0
	for i > 0 {
		if i&1 == 1 {
			ans += 2 * lucas(4+2*n)
		}
		n++
		i /= 2
	}
	return
}

func main() {
	starttime := time.Now()

	total := int64(0)

	for i := int64(0); i < 1<<(reps); i++ {
		if !check(i) {
			continue
		}

		summand := sum(i)

		if summand < tip {
			total += summand

			//the case where we can add phi + phi^-2
			if i&1 == 0 {
				summand += 2

				if summand < tip {
					total += summand
				}
			}
		}
	}

	fmt.Println(total + 1)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
