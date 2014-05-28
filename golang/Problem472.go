package main

import (
	"fmt"
	"time"
)

const (
	top          = 1000000000000
	periodically = 10000000
	mod          = 100000000
)

var lowCache = map[int]int{1: 1, 2: 2, 3: 2, 4: 4, 5: 3, 6: 6, 7: 2, 8: 6, 9: 3, 10: 8, 11: 2, 12: 6, 13: 2, 14: 4, 15: 9, 16: 4}

func wrap(x int) int {
	if x <= 16 {
		return lowCache[x]
	}
	return theory(x)
}

var einsCache = map[int]int{0: 4, 1: 3, 2: 8, 4: 6}

func theory(x int) int {
	eins := clip(x)

	if ans, ok := einsCache[eins]; ok {
		return ans
	}

	clipeins := clip(eins)

	if clipeins == 0 {
		return 4
	}

	clopeins := clop(eins)

	if 2*clipeins <= clopeins {
		return 2 * clipeins
	}

	clopx := clop(x)

	//if top two bits of x are set
	if temp := clopx>>1 + clopx; x&temp == temp {
		if base := (4 + clopeins - clipeins); 2*clipeins <= 2+clopeins {
			return base + clopx/2
		} else {
			return base
		}
	}

	if 2*clipeins > 2+clopeins {
		return 2 * (2 + clopeins - clipeins)
	}

	return 2 + (eins-1)*4/3
}

func clip(x int) int {
	return x &^ clop(x)
}

func clop(x int) int {
	ans := 1
	for x > 0 {
		x >>= 1
		ans <<= 1
	}
	return ans >> 1
}

func main() {
	starttime := time.Now()

	total := 0

	for i := 0; i <= 16 && i <= top; i++ {
		total += wrap(i)
	}

	for i := 17; i <= top; i++ {
		delta := theory(i) % mod
		total += delta
		total %= mod
		if i%periodically == 0 {
			fmt.Printf("%8d %d/%d = %.3f%%  (%s)\n", total, i, top, float32(100*i)/float32(top), time.Since(starttime))
		}
	}
	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
