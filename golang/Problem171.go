package main

import (
	"euler"
	"fmt"
	"math"
	"time"
)

const (
	top  int = 20
	mod      = 1000000000 //10^9
	ones     = 111111111  // nine 1s
)

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
			return [][]int{{sqrt}}, true
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

func paint(topaint [][]int, color int) [][]int {
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

func process(way []int) int64 {
	comp := compress(way)
	size := int64(len(way))
	choices := int64(1)

	for i := 0; i < len(comp); i++ {
		choices *= euler.Choose(size, int64(comp[i][1]))
		size -= int64(comp[i][1])
	}

	delta := int64(0)
	for i := 0; i < len(comp); i++ {
		delta += int64(comp[i][1]) * int64(comp[i][0])
		delta %= mod
	}

	delta *= choices
	delta /= int64(len(way))
	delta %= mod

	delta = (delta * ones)
	delta %= mod

	return delta
}

func main() {
	starttime := time.Now()

	hittable := make(map[int]bool)
	ans := int64(0)

	for i := 1; i*i <= top*9*9; i++ {
		hittable[i*i] = true
	}

	for i := range hittable {
		ways, _ := enumerate(9, top, i)
		for i := 0; i < len(ways); i++ {
			ans += process(ways[i])
			ans %= mod
		}
	}

	fmt.Println(ans)
	fmt.Println("Elapsed time:", time.Since(starttime))
}
