package main

import (
	"euler"
	"fmt"
	"time"
)

func eval(in []int) int {
	prod := 1
	for index, power := range in {
		//fmt.Println(index, power)
		if power > 0 {
			for i := 0; i < power; i++ {
				prod *= int(euler.Prime(int64(1 + index)))
			}
		}
	}
	return prod
}

//returns all possibilities less than max using primes not after prime
func possible(max, prime int) (toret [][]int) {
	toret = [][]int{}

	if prime < 1 {
		return [][]int{[]int{}}
	}

	piz := int(euler.Prime(int64(prime)))

	for _, under := range possible(max, prime-1) {
		toret = append(toret, append(under, 0))
	}

	top := 0
	for max >= piz {

		top++
		max /= piz

		for _, under := range possible(max, prime-1) {
			toret = append(toret, append(under, top))
		}

	}

	return
}

//returns nth binary number \leq 2^a
func bin(a, n int) []bool {
	if a == 0 {
		return []bool{}
	}

	if n%2 == 0 {
		return append(bin(a-1, n/2), false)
	}
	return append(bin(a-1, n/2), true)

}

func main() {
	starttime := time.Now()

	cmax := 1000
	pmax := int(euler.PrimePi(int64(cmax)))

	ABtries := possible(cmax, pmax)

	fmt.Println(len(ABtries))

	for _, AB := range ABtries {
		index := []int{}
		for p, pow := range AB {
			if pow > 0 {
				index = append(index, p)
			}
		}

		if len(index) > 3 {
			fmt.Println(AB, len(index))

		}

	}

	for i := 0; i < 16; i++ {
		fmt.Println(bin(4, i))
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
