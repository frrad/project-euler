package main

import (
	"euler"
	"fmt"
	"time"
)

//http://en.wikipedia.org/wiki/Chinese_remainder_theorem#A_constructive_algorithm_to_find_the_solution
func main() {
	starttime := time.Now()

	n := int64(13082761331670030)

	factors := euler.Factors(n)

	plist := make([]int64, 0)

	for _, p := range factors {
		if p[1] != 1 {
			panic("ruh roh")
		}

		plist = append(plist, p[0])
	}

	aposs := make([][]int64, len(plist))

	for i, p := range plist {
		for a := int64(1); a < p; a++ {

			if a*a*a%p == 1 {
				aposs[i] = append(aposs[i], a)
			}

		}

	}

	possibilityCount := 1

	for _, list := range aposs {
		possibilityCount *= len(list)
	}

	S := int64(0)

	for i := 0; i < possibilityCount; i++ {

		index := i

		a := make([]int64, 0)

		for j := 0; j < len(plist); j++ {
			a = append(a, aposs[j][index%len(aposs[j])])

			index /= len(aposs[j])
		}

		x := euler.ChineseRemainder(a, plist)

		for x < 0 {
			x += n
		}

		x %= n
		fmt.Println(x)

		S += x

	}

	fmt.Println(S - 1)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
