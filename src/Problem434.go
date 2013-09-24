package main

import (
	"fmt"
	"time"
)

func numberize(row []bool) (ans uint64) {
	for i := 0; i < len(row); i++ {
		if row[i] {
			ans += 1 << uint(i)
		}
	}
	return
}

func deNumberize(x uint64, lgth int) (slice []bool) {
	slice = make([]bool, lgth)
	for i := 0; i < lgth; i++ {
		if x&1 == 1 {
			slice[i] = true
		}
		x >>= 1
	}
	return
}

func show(state [][]bool) {
	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state[0]); j++ {
			if state[i][j] {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}

}

func count(state [][]bool) (total int) {
	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state[0]); j++ {
			if state[i][j] {
				total++
			}
		}
	}
	return
}

func evolve(currentState [][]bool) [][]bool {

	x := len(currentState)
	y := len(currentState[0])

	state := make([][]bool, x)
	for i := 0; i < x; i++ {
		line := make([]bool, y)
		copy(line, currentState[i])
		state[i] = line
	}

	oldCount, newCount := -1, count(state)
	for oldCount < newCount {
		state = evolveH(state)
		state = transpose(state)
		state = evolveH(state)
		state = transpose(state)
		oldCount, newCount = newCount, count(state)
	}

	return state
}

func evolveH(state [][]bool) [][]bool {
	oldCount, newCount := -1, count(state)

	for oldCount < newCount {

		list := numberizeState(state)

		for j := 0; j < len(list)-1; j++ {

			if list[j]&list[j+1] != 0 {
				list[j] = list[j] | list[j+1]
				list[j+1] = list[j]
			}

		}

		state = deNumberizeMatrix(list, len(state[0]))

		oldCount, newCount = newCount, count(state)

	}

	return state
}

func transpose(state [][]bool) (transp [][]bool) {
	x, y := len(state), len(state[0])

	transp = make([][]bool, y)

	for i := 0; i < y; i++ {
		transp[i] = make([]bool, x)
		for j := 0; j < x; j++ {
			transp[i][j] = state[j][i]
		}

	}

	return

}

func numberizeState(state [][]bool) []uint64 {
	x := len(state)

	ret := make([]uint64, x)

	for i := 0; i < x; i++ {
		ret[i] = numberize(state[i])
	}

	return ret
}

func deNumberizeMatrix(list []uint64, lgth int) [][]bool {
	ret := make([][]bool, len(list))

	for i := 0; i < len(list); i++ {
		ret[i] = deNumberize(list[i], lgth)
	}
	return ret

}

func rigid(state [][]bool) bool {
	if count(evolve(state)) == len(state)*len(state[0]) {
		return true
	}
	return false
}

func grid(n uint64, x, y int) [][]bool {
	state := make([][]bool, x)
	for i := 0; i < x; i++ {
		state[i] = make([]bool, y)
		for j := 0; j < y; j++ {
			if n&1 == 1 {
				state[i][j] = true
			}

			n >>= 1
		}
	}
	return state
}

func main() {
	starttime := time.Now()

	a, b := 6, 3
	size := uint64(1 << uint(a*b))

	count := 0

	for i := uint64(0); i < size; i++ {
		test := grid(i, a, b)
		if rigid(test) {
			count++
			if count%1000 == 0 {
				fmt.Println(count)
			}
		}

	}

	fmt.Println(a, b, count)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
