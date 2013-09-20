package main

import (
	"fmt"
	"time"
)

const top = 1 << 12

//a supports b
func supports(a, b int) bool {

	mask1, mask2 := a&b, a&b
	b = ^a & b

	for (mask1 != 0 || mask2 != 0) && b != 0 {
		mask1 <<= 1
		mask2 >>= 1

		m1, m2 := mask1&b, mask2&b

		b = ^mask1 & b
		b = ^mask2 & b

		mask1, mask2 = m1, m2

	}

	if b == 0 {
		return true
	}
	return false
}

func print(an int) {
	for an != 0 {
		fmt.Print(an & 1)
		an >>= 1
	}
	fmt.Print("\n")
}

func count(an int) (bits int) {

	for an != 0 {
		bits += 1 & an
		an >>= 1
	}
	return
}

func main() {
	starttime := time.Now()
	sups := make([][]int, top)

	for i := 0; i < top; i++ {

		for j := 0; j < top; j++ {

			if supports(i, j) {
				sups[i] = append(sups[i], j)
			}
		}
		fmt.Println(i)
	}

	sequences := [][][]int{[][]int{[]int{1<<6 - 1}}}

	test := 12

	for length := 1; length <= test; length++ {
		sequences = append(sequences, make([][]int, 0))

		for _, subseq := range sequences[length-1] {

			//fmt.Println(sequences)
			trailing := subseq[length-1]

			if trailing == 0 {
				continue
			}

			blocks := 0

			for _, line := range subseq {
				blocks += count(line)
			}

			for _, options := range sups[trailing] {

				if count(options)+blocks <= 18 {
					adder := make([]int, length+1)
					copy(adder[:length], subseq)
					adder[length] = options
					sequences[length] = append(sequences[length], adder)

				}

			}

		}

	}

	for _, size := range sequences {
		fmt.Println(len(size))
	}

	for _, sequence := range sequences[5] {
		for _, line := range sequence {
			print(line)
		}
		fmt.Println("====")
	}

	fmt.Println("Elapsed time:", time.Since(starttime))

}
