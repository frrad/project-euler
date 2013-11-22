package main

import (
	"euler"
	"fmt"
	"strconv"
	"time"
)

const target = 100

var grid map[[2]int]int
var chart map[int][2]int

func pad(n, lth int) string {
	ret := strconv.Itoa(n)
	for len(ret) < lth {
		ret = ret + " "
	}
	return ret
}

func blank(lth int) string {
	ret := ""
	for len(ret) < lth {
		ret = ret + " "
	}
	return ret
}

func shew(size, lth int) {
	for i := -1 * size; i < size; i++ {
		for j := -1 * size; j < size; j++ {
			digit := grid[[2]int{i, j}]
			if digit == 0 {
				fmt.Print(blank(lth))
			} else {
				fmt.Print(pad(digit, lth))
			}
		}
		fmt.Print("\n")
	}

}

func set(n, x, y int) {
	grid[[2]int{x, y}] = n
	chart[n] = [2]int{x, y}
}

//Yes, this is a horrible hack. Cleaner: use local information to determine
//direction of travel

func fill(last int) {
	i := 1
	set(i, 0, 0)
	x, y := -2, 0
	i++

	for duration := 1; duration < last; duration++ {

		a, b := 1, -1
		for j := 0; j < duration; j++ {
			set(i, x, y)
			x, y = x+a, y+b
			i++

		}

		a, b = 2, 0
		for j := 0; j < duration; j++ {
			set(i, x, y)

			x, y = x+a, y+b
			i++
		}

		a, b = 1, 1
		for j := 0; j < duration; j++ {
			set(i, x, y)

			x, y = x+a, y+b
			i++
		}

		a, b = -1, 1
		for j := 0; j < duration; j++ {
			set(i, x, y)

			x, y = x+a, y+b
			i++
		}

		a, b = -2, 0
		for j := 0; j < duration; j++ {
			set(i, x, y)

			x, y = x+a, y+b
			i++
		}

		a, b = -1, -1
		for j := 0; j < duration; j++ {
			set(i, x, y)

			x, y = x+a, y+b
			i++
		}

		x -= 3 - 1
		y -= 1 - 1

	}
}

func diffs(an int) [6]int {
	var loc [2]int
	if ans, ok := chart[an]; ok {
		loc = ans
	} else {
		panic("CHEESEBURGERS")
	}

	x, y := loc[0], loc[1]

	return diff(an, [6]int{
		grid[[2]int{x + 2, y}],
		grid[[2]int{x - 2, y}],
		grid[[2]int{x + 1, y + 1}],
		grid[[2]int{x + 1, y - 1}],
		grid[[2]int{x - 1, y - 1}],
		grid[[2]int{x - 1, y + 1}]})
}

func diff(an int, list [6]int) [6]int {
	for i := 0; i < 6; i++ {
		if an-list[i] > 0 {
			list[i] = an - list[i]
		} else {
			list[i] = list[i] - an
		}
	}
	return list
}

func PD(n int) (count int) {
	list := diffs(n)
	for i := 0; i < 6; i++ {
		if list[i] == 0 {
			panic("CHEESEBURGERS")
		}
		if euler.IsPrime(int64(list[i])) {
			count++
		}
	}
	return
}

func main() {
	starttime := time.Now()

	grid = make(map[[2]int]int)
	chart = make(map[int][2]int)

	fill(800)

	which := 1

	for i := 1; which <= target; i++ {
		if PD(i) == 3 {
			fmt.Print("{", i, ",", which, "},")
			which++
		}
	}

	fmt.Print("\n")

	fmt.Println("Elapsed time:", time.Since(starttime))
}
