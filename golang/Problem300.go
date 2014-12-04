package main

import (
	"fmt"
	"time"
)

const top = 15
const max_trinary = 1594323

var around = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

// Given a folding configuration (trinary number) return a slice of bitmasks
// to compute the score of a binary number in this folding.
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

	// Build reverse lookup table
	reverse_lookup := make(map[[2]int]int)
	for i, tuple := range positions {
		if _, ok := reverse_lookup[tuple]; !ok {
			reverse_lookup[tuple] = i
		} else {
			return nil
		}
	}

	ans := make([]int, 0)
	for i, current := range positions {
		for _, faddle := range around {
			aim[0], aim[1] = faddle[0]+current[0], faddle[1]+current[1]

			if val, ok := reverse_lookup[aim]; ok {
				if val > i {
					ans = append(ans, 1<<uint(i)|1<<uint(val))
				}
			}
		}
	}

	return ans
}

func main() {
	starttime := time.Now()

	opt := make(map[int]int)

	for pattern := 0; pattern < max_trinary; pattern++ {
		pattern_masks := masks(pattern)
		for protein := 0; protein < 1<<top; protein++ {
			score := 0
			for mask := range pattern_masks {
				if mask&protein == mask {
					score++
				}
			}
			if score > opt[protein] {
				opt[protein] = score
			}
		}

		fmt.Println(100 * pattern / max_trinary)
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
