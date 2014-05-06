//Reusing ideas  / functions from 171
package main

import (
	"euler"
	"fmt"
	"time"
)

const top = 18

//the ways in which we can add up to goal, with descending sequences of n
//integers, whose values are at most, most
func ways(n, goal, most int) [][]int {
	if goal < 0 || most*n < goal {
		return nil
	}

	if n == 1 {
		return [][]int{{goal}}
	}

	accumulate := make([][]int, 0)

	for i := 0; i <= most; i++ {
		if below := ways(n-1, goal-i, i); below != nil {
			accumulate = append(accumulate, paint(i, below)...)
		}
	}

	return accumulate
}

func paint(color int, topaint [][]int) [][]int {
	ans := make([][]int, 0)
	for _, strip := range topaint {
		ans = append(ans, append(strip, color))
	}
	return ans
}

//returns compressed representation of slice
func compress(input []int) [][2]int {
	ans := [][2]int{{input[0], 1}}
	pointer := 0
	for i := 1; i < len(input); i++ {
		if ans[pointer][0] == input[i] {
			ans[pointer][1]++
		} else {
			ans = append(ans, [2]int{input[i], 1})
			pointer++
		}
	}
	return ans
}

func distribute(state []int) (ans int64) {
	types := len(compress(state))

	for zeroPos := 0; zeroPos < types; zeroPos++ {
		thisAns := int64(1)
		openPlaces := top
		flat := compress(state)

		thisAns *= euler.Choose(int64(openPlaces-1), int64(flat[zeroPos][0]))
		openPlaces -= flat[zeroPos][0]
		flat[zeroPos][1]--

		openDigits := 9

		for iPos := 0; iPos < types; iPos++ {

			thisAns *= euler.Choose(int64(openDigits), int64(flat[iPos][1]))
			openDigits -= flat[iPos][1]

			for i := 0; i < flat[iPos][1]; i++ {
				thisAns *= euler.Choose(int64(openPlaces), int64(flat[iPos][0]))
				openPlaces -= flat[iPos][0]
			}
		}
		ans += thisAns
	}
	return
}

func main() {
	starttime := time.Now()

	total := int64(0)

	vectors := ways(10, top, 3)
	for _, vector := range vectors {
		total += distribute(vector)
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
