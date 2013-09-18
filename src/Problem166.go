package main

import (
	"fmt"
	"time"
)

const dim = 4

var cheat map[[dim][dim]int]int64
var mask [dim][dim][11]map[[dim][dim]int]bool

func count(state [dim][dim]int) (ways int64) {
	//fmt.Println(state)

	//First check the memo
	if ans, ok := cheat[state]; ok {
		return ans
	}

	//Now we try very hard to show this is impossible
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
			store(state, 0)
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
		store(state, 0)
		return 0
	}

	if isComplete {
		store(state, 1)

		return 1
	}

	bestWidth := 99
	bestBelow, bestAbove := 0, 9
	bestX, bestY := 0, 0

	//fmt.Println(options)

	//given an unknown coordinate, what are its min / max?

	for x := dim - 1; x >= 0; x-- {

		for y := dim - 1; y >= 0; y-- {

			below, above := 0, 9

			if !options[[2]int{x, y}] {
				continue
			}

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
				store(state, 0)
				return 0
			}
			if above-below < bestWidth {
				bestAbove, bestBelow, bestWidth = above, below, above-below
				bestX, bestY = x, y
			}
		}
	}

	ans := int64(0)
	restrict := state
	for set := bestBelow; set <= bestAbove; set++ {
		restrict[bestX][bestY] = set
		//fmt.Println(restrict)
		ans += count(restrict)
	}

	store(state, ans)
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

func store(state [dim][dim]int, ans int64) {
	/*if ans == 1 {
		fmt.Println("-------")
		for _, line := range state {
			fmt.Println(line)
		}
		fmt.Println("-------")
	}*/

	state = rotate(state)

	for i := 0; i < 3; i++ {
		cheat[state] = ans
		state = rotate(state)
	}

	state = flip(state)
	for i := 0; i < 4; i++ {
		cheat[state] = ans
		state = rotate(state)
	}

}

func addMask(state [dim][dim]int) {

	if deathMask(state) {
		return
	}

	//should dedupe!

	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {

			mask[i][j][state[i][j]+1][state] = true

		}
	}

}

func deathMask(state [dim][dim]int) bool {
	if len(match(state)) > 0 {
		return true
	}
	return false
}

func match(state [dim][dim]int) map[[dim][dim]int]bool {
	sets := make([][2]map[[dim][dim]int]bool, 0)
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			sets = append(sets, [2]map[[dim][dim]int]bool{mask[i][j][state[i][j]+1], mask[i][j][0]})

		}
	}

	for i := 0; i < len(sets); i++ {
		for j := 0; j < len(sets)-i-1; j++ {
			if len(sets[j][0])+len(sets[j][1]) > len(sets[j+1][0])+len(sets[j+1][1]) {
				sets[j], sets[j+1] = sets[j+1], sets[j]
			}
		}
	}

	temp := make(map[[dim][dim]int]bool)

	for key, _ := range sets[0][0] {
		temp[key] = true
	}

	for key, _ := range sets[0][1] {
		temp[key] = true
	}

	i := 1

	for len(temp) > 0 && i < len(sets) {
		for key, _ := range temp {
			if !sets[i][0][key] && !sets[i][1][key] {
				delete(temp, key)
			}
		}
		i++
	}

	return temp

}

func rotate(in [dim][dim]int) (out [dim][dim]int) {

	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			out[j][dim-i-1] = in[i][j]
		}
	}
	return
}

func flip(in [dim][dim]int) (out [dim][dim]int) {
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			out[i][dim-j-1] = in[i][j]
		}
	}
	return
}

func main() {
	starttime := time.Now()

	cheat = make(map[[dim][dim]int]int64)
	//mask = make([dim][dim][11]map[[dim][dim]int]bool)
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			for k := 0; k < 11; k++ {
				mask[i][j][k] = make(map[[dim][dim]int]bool)
			}
		}
	}

	total := int64(0)

	for i := 9; i >= 0; i-- {
		for j := 9; j >= 0; j-- {
			test := [dim][dim]int{
				[dim]int{i, j, -1, -1},
				[dim]int{-1, -1, -1, -1},
				[dim]int{-1, -1, -1, -1},
				[dim]int{-1, -1, -1, -1},
			}
			temp := count(test)
			total += temp
			fmt.Println(total, "\t", temp, "\t", i, j, "\t", time.Since(starttime))
		}
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
