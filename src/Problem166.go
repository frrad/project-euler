package main

import (
	"fmt"
	"time"
)

const dim = 4

func count(state [dim][dim]int) (ways int64) {

	isComplete := true
	min, max := 0, 9*dim
	options := make(map[[2]int]bool)

	for i := 0; i < dim; i++ {
		min1, max1, min2, max2 := 0, 0, 0, 0
		for j := 0; j < dim; j++ {
			if state[i][j] != -1 {
				min1 += state[i][j]
				max1 += state[i][j]
			} else {
				max1 += 9
				isComplete = false
				options[[2]int{i, j}] = true
			}

			if state[j][i] != -1 {
				min2 += state[j][i]
				max2 += state[j][i]
			} else {
				max2 += 9
			}

		}

		if min1 > min {
			min = min1
		}
		if min2 > min {
			min = min2
		}

		if max1 < max {
			max = max1
		}
		if max2 < max {
			max = max2
		}

		if max < min {

			return 0
		}
	}

	min1, max1, min2, max2 := 0, 0, 0, 0
	for j := 0; j < dim; j++ {
		if state[j][j] != -1 {
			min1 += state[j][j]
			max1 += state[j][j]
		} else {
			max1 += 9
		}

		if state[j][dim-j-1] != -1 {
			min2 += state[j][dim-j-1]
			max2 += state[j][dim-j-1]
		} else {
			max2 += 9
		}

	}

	if min1 > min {
		min = min1
	}
	if min2 > min {
		min = min2
	}

	if max1 < max {
		max = max1
	}
	if max2 < max {
		max = max2
	}

	if max < min {
		return 0
	}

	if isComplete {
		return 1
	}

	bestWidth := 99
	bestBelow, bestAbove := 0, 9
	bestX, bestY := 0, 0

	//fmt.Println(options)

	//given an unknown coordinate, what are its min / max?

	for opt, _ := range options {
		x, y := opt[0], opt[1]
		below, above := 0, 9

		sum1a, sum1b, sum2a, sum2b := 0, 0, 0, 0
		for i := 0; i < dim; i++ {
			if state[x][i] != -1 {
				sum1a += state[x][i]
				sum1b += state[x][i]
			} else {
				sum1b += 9
			}
			if state[i][y] != -1 {
				sum2a += state[i][y]
				sum2b += state[i][y]
			} else {
				sum2b += 9
			}
		}

		sum1b -= 9
		sum2b -= 9

		below = sup(min-sum1b, below)
		below = sup(min-sum2b, below)

		above = inf(max-sum1a, above)
		above = inf(max-sum2a, above)

		if x-y == 0 {
			suma, sumb := 0, 0
			for i := 0; i < dim; i++ {
				if state[i][i] != -1 {
					suma += state[i][i]
					sumb += state[i][i]
				} else {
					sumb += 9
				}

			}

			sumb -= 9

			below = sup(min-sumb, below)
			above = inf(max-suma, above)
		} else if x+y == dim-1 { //If we're on the diagonal we're not on antidiagonal
			suma, sumb := 0, 0
			for i := 0; i < dim; i++ {
				if state[i][dim-1-i] != -1 {
					suma += state[i][dim-1-i]
					sumb += state[i][dim-1-i]
				} else {
					sumb += 9
				}
			}
			sumb -= 9

			below = sup(min-sumb, below)
			above = inf(max-suma, above)
		}

		if above < below {
			return 0
		}
		if above-below < bestWidth {
			bestAbove, bestBelow, bestWidth = above, below, above-below
			bestX, bestY = x, y
		}
	}

	ans := int64(0)
	restrict := state
	for set := bestBelow; set <= bestAbove; set++ {
		restrict[bestX][bestY] = set
		//fmt.Println(restrict)
		ans += count(restrict)
	}

	return ans
}

func sup(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func inf(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	starttime := time.Now()

	test := [dim][dim]int{
		[dim]int{-1, -1, -1, -1},
		[dim]int{-1, -1, -1, -1},
		[dim]int{-1, -1, -1, -1},
		[dim]int{-1, -1, -1, -1},
	}

	fmt.Println(count(test))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
