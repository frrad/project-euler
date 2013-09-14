package main

import (
	//"euler"
	"fmt"
	"time"
)

const flips = 1000
const billion = 1000000000
const delta = .001

func canMiss(f float64) int {
	misses := 0
	bucks := float64(1)
	for i := 0; i < flips; i++ {
		if bucks < billion {
			bucks *= (1 + 2*f)
		} else {
			bucks *= (1 - f)
			misses++
		}
	}
	if bucks >= billion {
		return misses
	}
	if misses > 0 {
		return misses - 1
	}

	return 0

}

func main() {
	starttime := time.Now()

	maxmiss := 0
	for i := 0.; i <= 1; i += .001 {
		//fmt.Print("{", i, ",", canMiss(i), "},")
		if canMiss(i) > maxmiss {
			maxmiss = canMiss(i)
		}
	}

	fmt.Println("Can miss at most", maxmiss)

	//N[Sum[Binomial[1000, k], {k, 1, 568}]/2^1000, 12]

	fmt.Println("Elapsed time:", time.Since(starttime))
}
