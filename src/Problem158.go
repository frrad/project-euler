package main

import (
	"fmt"
	"time"
)

const max = 26

var asstable map[[3]int]int64

//number of (weakly) ascending sequences starting with a
//ending on b with length length
func ass(a, b, length int) int64 {
	if length == 1 && a != b {
		return 0
	}

	if length == 2 {
		if a < b {
			return 1
		} else {
			return 0
		}
	}

	if a == b {
		if length == 1 {
			return 1
		} else {
			return 0
		}
	}

	if a > b {
		return 0
	}

	if answer, ok := asstable[[3]int{a, b, length}]; ok {
		return answer
	}

	answer := int64(0)
	for i := a + 1; i < b; i++ {
		answer += ass(i, b, length-1)
	}

	asstable[[3]int{a, b, length}] = answer
	return answer

}

func oneBreak(a, b, length int) int64 {

	answer := int64(0)

	//length of 1st run
	for i := 1; i < length; i++ {

		for A := a; A <= max; A++ {
			for B := 1; B < A && B <= b; B++ {

				change := ass(a, A, i) * ass(B, b, length-i)
				if change != 0 && false {
					fmt.Println("A", A, "B", B, "i", i)
					fmt.Println(change)
				}
				answer += change
			}

		}

	}

	return answer
}

func p(n int) int64 {
	answer := int64(0)
	for a := 1; a <= max; a++ {
		for b := 1; b <= max; b++ {
			if b == a {
				b++

				if b > max {
					continue
				}

			}

			change := oneBreak(a, b, n)

			fmt.Println(a, b, ":\t", change)
			answer += change
		}
	}
	return answer
}

func main() {
	starttime := time.Now()

	asstable = make(map[[3]int]int64)

	fmt.Println(p(27))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
