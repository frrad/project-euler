package main

import (
	"fmt"
	"time"
)

const (
	WIDTH = 32
	DEPTH = 10
)

func toppers(bottom []bool) [][]bool {
	layers := make([][]bool, 1)
	ans := make([][]bool, 0)

	if bottom[1] {
		layers[0] = []bool{false, false, true}
	} else {
		layers[0] = []bool{false, true}
	}

	for len(layers) > 0 {
		inception := make([][]bool, 0)
		for _, layer := range layers {

			size := len(layer)

			if size == len(bottom) {
				ans = append(ans, layer)
				continue
			}

			if size+1 == len(bottom) {
				inception = append(inception, append(layer, false))
			}
			if size+2 == len(bottom) {
				inception = append(inception, append(layer, false, false))
			}

			big := size+2 < len(bottom) && size+3 != len(bottom) && !bottom[size+2]
			small := (size+1 < len(bottom) && !bottom[size+1]) && size+2 != len(bottom)

			if big && !small {
				inception = append(inception, append(layer, false, false, true))
			}
			if small && !big {
				inception = append(inception, append(layer, false, true))
			}

			if big && small {
				fake := make([]bool, len(layer))
				copy(fake, layer)
				inception = append(inception, append(layer, false, true), append(fake, false, false, true))
			}

		}
		layers = inception
	}
	return ans
}

func bottom(i, n uint) []bool {
	spec := i
	layer := []bool{}
	for len(layer)+3 < int(n) {
		here := spec % 2
		spec /= 2
		if here == 1 {
			layer = append(layer, false, false, true)
		} else {
			layer = append(layer, false, true)
		}
	}

	if spec != 0 {
		return nil
	}

	if len(layer)+1 == int(n-1) {
		return append(layer, false)
	}
	if len(layer)+2 == int(n-1) {
		return append(layer, false, false)
	}

	return nil
}

var memo = make(map[[2]int]int)

func stack(depth int, bottom []bool) (ans int) {
	if depth == 1 {
		return 1
	}

	if ans, ok := memo[[2]int{depth, memFn(bottom)}]; ok {
		return ans
	}

	for _, bar := range toppers(bottom) {
		ans += stack(depth-1, bar)
	}

	memo[[2]int{depth, memFn(bottom)}] = ans

	return
}

func memFn(x []bool) (ans int) {
	for _, tf := range x {
		ans *= 2
		if tf {
			ans += 1
		}
	}
	return
}

func main() {
	starttime := time.Now()

	total := 0

	for i := uint(0); i < 1<<(WIDTH/2); i++ {
		if base := bottom(i, WIDTH); base != nil {
			total += stack(DEPTH, base)
		}
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
