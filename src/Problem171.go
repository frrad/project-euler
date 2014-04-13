package main

import (
	"euler"
	"fmt"
	"math"
	"time"
)

const top int = 5

//return all strictly decreasing sets of numbers at most max, and at least 0 of
//length lth whose squares sum to target
func enumerate(max, lth, target int) (ans [][]int, ok bool) {
	if target < 0 {
		return nil, false
	}

	if max*max*lth < target {
		//can't get big enough
		return nil, false
	}

	if lth == 1 {
		if euler.IsSquare(int64(target)) {
			sqrt := int(math.Sqrt(float64(target)))
			return [][]int{[]int{sqrt}}, true
		} else {
			return nil, false
		}
	}

	ans = make([][]int, 0)

	for i := 0; i <= max && i*i <= target; i++ {
		tail, works := enumerate(i, lth-1, target-(i*i))
		if works {
			ok = true
			ans = append(ans, paint(tail, i)...)
		}
	}

	return
}

//gives the sum of all digits appearing in a `following' place, with appropriate
//multiplicities.
func following(in []int) int64 {
	delta := int64(0)

	state := compress(in)
	places := len(in)

	zeroes := 0
	if state[0][0] == 0 {
		zeroes = state[0][1]
	}

	if zeroes+1 >= places {
		return 0
	}

	//index of first nonzero elt
	nonzero := 0
	if state[0][0] == 0 {
		nonzero++
	}

	places-- //our distinguished element takes a place

	for i := nonzero; i < len(state); i++ {

		multiplicity := int64(1)
		available := places
		multiplicity *= euler.Choose(int64(places-1), int64(zeroes))
		available -= zeroes
		for j := nonzero; j < len(state); j++ {
			if j != i {
				multiplicity *= euler.Choose(int64(available), int64(state[j][1]))
				available -= state[j][1]
			} else {
				multiplicity *= euler.Choose(int64(available), int64(state[j][1]-1))
				available -= (state[j][1] - 1)
			}
		}
		delta += multiplicity * int64(state[i][0])

	}
	return delta
}

func paint(topaint [][]int, color int) [][]int {
	ans := make([][]int, 0)
	for _, strip := range topaint {
		ans = append(ans, append(strip, color))
	}
	return ans
}

//returns compressed representation of slice
func compress(input []int) [][2]int {
	ans := [][2]int{[2]int{input[0], 1}}
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

func main() {
	starttime := time.Now()

	hittable := make(map[int]bool)

	potato := top * 9 * 9

	count := 0

	for i := 1; i*i <= potato; i++ {
		hittable[i*i] = true
	}

	fmt.Println(hittable)

	for i := range hittable {

		ways, _ := enumerate(9, top, i)

		fmt.Println(i, len(ways))
		count += len(ways)

	}

	fmt.Printf("Count: %d\n", count)

	//	example, _ := enumerate(9, top, 25)
	exemplary := []int{0, 0, 0, 0, 3, 4}
	//	exemplary := example[0]

	fmt.Printf("%d\n%d\n", exemplary, following(exemplary))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
