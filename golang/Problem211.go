package main

import (
	"euler"
	"fmt"
	"math"
	"time"
)

var sigtab map[uint64]uint64

func sigma(n uint64) uint64 {
	if ans, ok := sigtab[n]; ok {
		return ans
	}

	top := uint64(math.Sqrt(float64(n))) + 1
	prime := uint64(2)

	for i := int64(1); prime <= top; i++ {
		prime = uint64(euler.Prime(i))
		if n%prime == 0 {
			j := 0

			oldn := n

			for n%prime == 0 {
				n /= prime
				j++
			}

			adjust := uint64(1)
			factor := prime * prime

			for k := 1; k <= j; k++ {
				adjust += factor
				factor *= prime * prime
			}

			ans := adjust * sigma(n)

			sigtab[oldn] = ans
			return sigma(oldn)
		}

	}

	// n prime:
	sigtab[n] = n*n + 1

	return sigma(n)
}

func isSquare(n uint64) bool {
	sqrt := uint64(math.Sqrt(float64(n)))
	if sqrt*sqrt == n || (sqrt+1)*(sqrt+1) == n {
		return true
	}
	return false
}

func main() {
	starttime := time.Now()
	sigtab = make(map[uint64]uint64)
	sigtab[1] = 1

	total := uint64(0)
	count := 1

	for i := uint64(1); i < 64000000; i++ {
		sig := sigma(i)

		if isSquare(sig) {
			// fmt.Println(count, ":", i, "\t\t\t", sig, "\t\t\t", total)
			count++
			total += i
		}
	}
	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
