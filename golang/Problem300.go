package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

const top = 15
const max_trinary = 1594323 // 3^(top-2)

var around = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

// Given a folding configuration (trinary number) return a slice of bitmasks
// to compute the score of a binary number in this folding.
func masks(arranged int) []int {
	var positions [top][2]int

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

			if val, ok := reverse_lookup[aim]; ok && val > i {
				ans = append(ans, 1<<uint(i)|1<<uint(val))
			}
		}
	}

	return ans
}

// Given a set of masks, score a protein
func score(pattern_masks []int, protein int) (score int) {
	for _, mask := range pattern_masks {
		if mask&protein == mask {
			score++
		}
	}
	return
}

func makeKey(input []int) (ans string) {
	for _, num := range input {
		ans += strconv.Itoa(num)
		ans += "|"
	}
	return
}

func main() {
	starttime := time.Now()

	opt := make(map[int]int)
	configurations := make([][]int, 0)

	seen := make(map[string]bool)

	for pattern := 0; pattern < max_trinary; pattern++ {
		if config := masks(pattern); len(config) > 0 {
			sort.Ints(config)
			if key := makeKey(config); !seen[key] {
				seen[key] = true
				configurations = append(configurations, config)
			}
		}
	}

	fmt.Println(len(configurations))

	for _, pattern_masks := range configurations {
		if len(pattern_masks) == 0 {
			continue
		}

		for protein := 0; protein < 1<<top; protein++ {
			if my_score := score(pattern_masks, protein); my_score > opt[protein] {
				opt[protein] = my_score
			}
		}
	}

	total := 0
	for _, score := range opt {
		total += score
	}
	fmt.Println(total, float64(total)/(1<<top))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
