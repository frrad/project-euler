package main

import (
	"fmt"
	"time"
)

const target = 1000

func layer(a, b, c, n int) int {
	return 4*(n-1)*(n-1) + 4*(a+b+c-1)*(n-1) + 2*(a*b+a*c+b*c)
}

func main() {
	starttime := time.Now()

	targetLoc := 0
	lid := 100

	for targetLoc == 0 {

		state := make([]int, lid+1)

		for a := 1; layer(a, 1, 1, 1) <= lid; a++ {
			for b := a; layer(a, b, 1, 1) <= lid; b++ {
				for c := b; layer(a, b, c, 1) <= lid; c++ {
					for n := 1; layer(a, b, c, n) <= lid; n++ {
						state[layer(a, b, c, n)]++
					}
				}

			}
		}

		for i := 0; i < lid; i++ {
			if state[i] == target {
				targetLoc = i
				break
			}
		}

		lid *= 2

	}

	fmt.Println(targetLoc)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
