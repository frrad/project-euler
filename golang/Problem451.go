package main

import (
	"euler"
	"fmt"
	"time"
)

const top = 20000000

func l(n int) int64 {
	primes := euler.Factors(int64(n))

	nVec := make([]int64, len(primes))

	for i, factor := range primes {
		nVec[i] = euler.IntExp(factor[0], factor[1])
	}

	aCans := make([][]int64, len(primes))

	max := 1

	for i, n := range nVec {
		if n%2 == 0 {
			if n == 2 {
				aCans[i] = []int64{1}
			}
			if n == 4 {
				aCans[i] = []int64{1, 3}
				max *= 2
			}
			if n > 4 {
				aCans[i] = []int64{1, n/2 - 1, n/2 + 1, n - 1}
				max *= 4
			}

		} else {
			aCans[i] = []int64{1, n - 1}
			max *= 2
		}
	}

	ans := int64(1)

	aVec := make([]int64, len(primes))

	// speed up by reusing CRT data
	for i := 0; i < max; i++ {
		spec := i
		for j := 0; j < len(primes); j++ {
			possibilities := len(aCans[j])
			aVec[j] = aCans[j][spec%possibilities]
			spec /= possibilities
		}

		sqrtOfUnity := euler.ChineseRemainder(aVec, nVec)
		if sqrtOfUnity > ans && sqrtOfUnity < int64(n-1) {
			ans = sqrtOfUnity
		}
	}

	return ans
}

func main() {
	starttime := time.Now()

	total := int64(0)

	for i := 3; i <= top; i++ {
		total += l(i)
		if i%100000 == 0 {
			fmt.Printf("%8d/%8d : %d\n", i, top, total)
		}
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
