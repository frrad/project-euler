package main

import (
	"fmt"
	"time"
)

const top = 15

// Given a folding configuration (trinary number) return a slice of bitmasks to
// compute the score of a binary number in this folding.
func masks(arranged int) []int {
	var positions [15][2]int

	positions[0] = [2]int{0, 0}
	positions[1] = [2]int{1, 0}

	aim := [2]int{1, 0}

	for i := 2; i < top; i++ {
		indicator := arranged % 3
		arranged /= 3

		if indicator == 1 {
			aim[0], aim[1] = -aim[1], aim[0]
		}
		if indicator == 2 {
			aim[0], aim[1] = aim[1], -aim[0]
		}

		positions[i][0] = positions[i-1][0] + aim[0]
		positions[i][1] = positions[i-1][1] + aim[1]
	}

	//	fmt.Println(positions)

	reverse_lookup := make(map[[2]int]int)
	for i, tuple := range positions {
		if _, ok := reverse_lookup[tuple]; !ok {
			reverse_lookup[tuple] = i
		} else {
			return nil
		}
	}
	return []int{234}
}

func main() {
	starttime := time.Now()

	a, b := 0, 0

	for i := 0; i < 1594323; i++ {
		if len(masks(i)) == 0 {
			a++
		} else {
			b++
		}
	}

	fmt.Println(a, b)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
