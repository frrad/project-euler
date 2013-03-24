package main

import (
	"./euler"
	"fmt"
	"time"
)

const height = 1000000

func next(n int) int {
	fmt.Println(n)
	if euler.IsPrime(int64(n)) {
		return 1
	}

	answer := 1
	factors := euler.Factor(int64(n))

	for k := 0; k < len(factors); k++ {

		p := int(factors[k])

		//	fmt.Println("p", p)

		term := 1

		term *= p

		for i := k; i < len(factors) && int(factors[i]) == p; i++ {
			//fmt.Println("p", p)

			term *= p
			k = i
		}

		answer *= (term - 1) / (p - 1)

		fmt.Println(answer)
	}

	return answer - n
}

func main() {
	starttime := time.Now()

	duplication := [height]int{}

	winnar := 0
	longest := 0

	for start := 0; start < height; start++ {

		for duplication[start] != 0 {
			start++
		}
		//	fmt.Println("starting at", start)

		chain := map[int]bool{}
		chain[start] = true
		current := next(start)

		length := 1
		for !chain[current] && current > 0 && current < height && duplication[current] == 0 {
			chain[current] = true
			current = next(current)
			length++
		}

		if current < height && duplication[current] == 0 && current > 0 {

			//	fmt.Println("detected loop at", current)
			split := current

			for ; start != current; length-- {
				//		fmt.Println("crunching", start)
				duplication[start] = -1
				start = next(start)
			}
			//	fmt.Println("loop has length", length)

			//	fmt.Println("contains", current)
			if length > longest {
				longest = length
				winnar = current
				fmt.Println(winnar, longest)

			}

			duplication[current] = length
			current = next(current)

			for current != split {
				//		fmt.Println("contains", current)
				duplication[current] = length
				current = next(current)
			}

		} else {
			current = start
			for current < height && duplication[current] == 0 && current > 0 {
				//		fmt.Println("crunching", current)
				duplication[current] = -1
				current = next(current)
			}
		}

	}

	fmt.Println(winnar, longest)

	fmt.Println("Elapsed time:", time.Since(starttime))

}
