package main

import (
	"euler"
	"fmt"
	"math"
	"time"
)

//return all strictly decreasing sets of numbers at most
//max, and at least 0 of length lth whose squares sum to
//target
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

func paint(topaint [][]int, color int) [][]int {
	ans := make([][]int, 0)
	for _, strip := range topaint {
		ans = append(ans, append(strip, color))
	}
	return ans
}

func main() {
	starttime := time.Now()

	hittable := make(map[int]bool)

	potato := 20 * 9 * 9

	for i := 1; i*i <= potato; i++ {
		hittable[i*i] = true
	}

	fmt.Println(hittable)

	for i := range hittable {

		ways, _ := enumerate(9, 20, i)

		fmt.Println(i, len(ways))

	}

	fmt.Println(enumerate(9, 20, 225))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
