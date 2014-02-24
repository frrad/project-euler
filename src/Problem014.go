package main

import (
	"fmt"
	"time"
)

const top = 1000000

var memo map[uint]uint

func collatz(n uint) uint {
	if ans, ok := memo[n]; ok {
		return ans
	}

	if n%2 == 0 {
		memo[n] = collatz(n/2) + 1
		return collatz(n)
	}

	memo[n] = collatz(3*n+1) + 1
	return collatz(n)

}

func main() {
	starttime := time.Now()

	memo = make(map[uint]uint)
	memo[1] = 1
	var best, where uint

	for i := uint(1); i <= top; i++ {
		if collatz(i) > best {
			best = collatz(i)
			where = i
			// fmt.Println(i, best)
		}
	}

	fmt.Println(where)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
