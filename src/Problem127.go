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
	if answer, ok := PPCache[[2]int{max, p}]; ok {
		return answer
	}

	if p == 0 {
		return 1
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

	i, j := 23, 5
	fmt.Println(possiPi(i, j))
	fmt.Println(possible(i, j))

	/*


			cmax := 120000
		pmax := int(euler.PrimePi(int64(cmax))) + 1

			ccount := 0
			csum := 0

				fmt.Println(pmax)
				ABtries := possible(cmax*cmax/4, pmax)

				fmt.Println(len(ABtries), "\n")

				for _, AB := range ABtries {
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
							csum += evalC
						}

					}

				}

				fmt.Println("\n", ccount, csum)

	*/

	fmt.Println("Elapsed time:", time.Since(starttime))
}
