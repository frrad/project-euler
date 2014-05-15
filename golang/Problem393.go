package main

import (
	"fmt"
	"time"
)

const top = 4

func compatible(a, b int) bool {
	gate := 1 + 1<<(top-1)
	if (a|b)&gate != gate {
		return false
	}

	for a != 0 || b != 0 {
		if (a|b)&1 == 0 {
			return false
		}

		//fmt.Println("<>")

		for a&b&1 == 1 {
			a >>= 1
			b >>= 1
			//fmt.Println("take one")
		}

		if (a^b)&1 == 1 {
			//fmt.Println("start")
			a >>= 1
			b >>= 1
			for (a|b)&1 == 0 && (a > 0 || b > 0) {
				a >>= 1
				b >>= 1
				//fmt.Println("gloop")
			}
			if (a^b)&1 == 1 {
				a >>= 1
				b >>= 1
				//fmt.Println("stop")
			} else {
				return false
			}

		}

		//fmt.Println(a%2, b%2)
	}

	return true
}

func shew(x int) {
	for i := 0; i < top; i++ {
		if x%2 == 1 {
			fmt.Print("|")
		} else {
			fmt.Print("-")
		}
		x /= 2
	}
	fmt.Print("\n")
}

func workable(a int) []int {
	ans := make([]int, 0)

	for b := 0; b < 1<<top; b++ {
		if compatible(a, b) {
			ans = append(ans, b)
		}
	}
	return ans
}

func main() {
	starttime := time.Now()

	settable := [][]int{{0}}

	for i := 0; i < top-1; i++ {

		gloop := make([][]int, 0)

		for _, set := range settable {

			gasket := set[len(set)-1]
			for _, x := range workable(gasket) {
				newb := make([]int, len(set)+1)
				copy(newb[:len(set)], set)
				newb[len(newb)-1] = x
				gloop = append(gloop, newb)
			}

		}
		settable = gloop
		fmt.Println(i, len(settable))
	}

	for _, set := range settable {
		if compatible(0, set[len(set)-1]) {
			fmt.Println(set)
			for _, i := range set {
				shew(i)
			}

		}
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
