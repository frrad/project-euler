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

var PPCache map[[2]int]int

func possiPi(max, p int) int {

	if p <= 0 {
		return 1
	}

	if answer, ok := PPCache[[2]int{max, p}]; ok {
		return answer
	}

	piz := int(euler.Prime(int64(p)))

	answer := 0

	answer += possiPi(max, p-1)

	mix := max
	for mix >= piz {
		mix /= piz

		answer += possiPi(mix, p-1)

	}

	PPCache[[2]int{max, p}] = answer
	return answer
}

var PNCache map[[3]int][]int

func possiN(max, p, index int) []int {
	if p == 0 {
		return []int{}
	}
	/*
		if answer, ok := PNCache[[3]int{max, p, index}]; ok {
			answer2 := make([]int, p)
			copy(answer2, answer)
			return answer2
		}
	*/
	//rmax, rp, rindex := max, p, index

	piz := int(euler.Prime(int64(p)))
	lead := 0

	for possiPi(max, p-1) <= index && max >= piz {
		//	fmt.Println("max=", max, "piz=", piz, "index=", index, "lead=", lead+1)

		index -= (possiPi(max, p-1))

		max /= piz
		lead++
	}

	/*fmt.Println("decide:", lead)

	answer := make([]int, rp)
	answer2 := make([]int, rp)
	copy(answer, append(possiN(max, p-1, index), lead))
	copy(answer2, answer)

	PNCache[[3]int{rmax, rp, rindex}] = answer*/
	return append(possiN(max, p-1, index), lead)

	if possiPi(max, p) < index {
		return []int{}
		fmt.Println("ERROR")
	}

	return []int{}
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

	PPCache = make(map[[2]int]int)
	PNCache = make(map[[3]int][]int)

	cmax := 120000
	pmax := int(euler.PrimePi(int64(cmax))) + 1
	fmt.Println("found pmax:", pmax)
	MAX := cmax * cmax / 4

	ccount := 0
	csum := int64(0)

	for ABcounter := 0; ABcounter < possiPi(MAX, pmax); ABcounter++ {

		if ABcounter%100 == 0 {
			fmt.Println(100 * ABcounter / MAX)
		}

		AB := possiN(MAX, pmax, ABcounter)

		index := make([]int, 0)

		radstart := 1

		for p, pow := range AB {
			if pow > 0 {
				index = append(index, p)
				radstart *= int(euler.Prime(int64(p + 1)))
			}
		}

		nonz := len(index)
		splits := int(euler.Exp2(nonz))

		for i := 0; i < splits/2 && radstart < cmax; i++ {
			splitter := bin(nonz, i)

			evalA := 1
			evalB := 1
			rad := radstart

			for k, p := range index {
				muhprime := int(euler.Prime(int64(p + 1)))
				for j := 0; j < AB[p]; j++ {
					if splitter[k] {
						evalA *= muhprime
					} else {
						evalB *= muhprime

					}
				}
			}

			evalC := evalA + evalB

			if evalA > cmax || evalB > cmax || evalC > cmax || evalC <= rad {

				continue

			}

			see := euler.Factors(int64(evalC))

			flag := false

			for _, pair := range see {

				rad *= int(pair[0])

				if rad >= evalC {
					flag = true
					break
				}
			}

			if !flag {
				fmt.Println(evalA, evalB, evalC)
				ccount++
				csum += int64(evalC)
				fmt.Println("\nccount:", ccount, "csum:", csum)

			}

		}

	}

	//fmt.Println("\n", "ccount:", ccount, "csum:", csum)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
