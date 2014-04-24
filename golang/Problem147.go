package main

import (
	"euler"
	"fmt"
	"time"
)

var borMemo map[[2]int]uint64

//returns the number of subrectangles in a x b grid
//with the boring orientation
func boring(a, b int) uint64 {

	if ans, ok := borMemo[[2]int{a, b}]; ok {
		return ans
	}

	if a > 1 {
		//we decompose into configurations which avoid the top row
		//those which avoid the bottom row, and those of full height
		borMemo[[2]int{a, b}] = 2*boring(a-1, b) - boring(a-2, b) + boring(1, b)
		return boring(a, b)
	}

	if a == 1 && b > 1 {
		borMemo[[2]int{a, b}] = uint64(b) + boring(a, b-1)
		return boring(a, b)
	}

	if a <= 0 || b <= 0 {
		return 0
	}

	//a==b==1
	return 1

}

func sort(a, b int) (int, int) {
	if a < b {
		return b, a
	}
	return a, b
}

func predict(a, b int) uint64 {
	b, a = sort(a, b)

	if a == 1 {
		return uint64(b - 1)
	}

	offset := 0

	factors := euler.Factors(int64(a))

	if factors[0][0] > 3 {
		offset = -1
	}

	answer := int64(offset - (a / 2) + (a*a)/6 - (2*a*a*a*a)/3 - (a*b)/3 + (4*a*a*a*b)/3)

	if answer < 0 {
		panic("negative answer is bad")
	}

	return uint64(answer)
}

func blend(a, b int) uint64 {

	return boring(a, b) + predict(a, b)
}

func main() {
	starttime := time.Now()

	borMemo = make(map[[2]int]uint64)

	sum := uint64(0)
	tip, top := 47, 43

	for i := 1; i <= tip; i++ {
		for j := 1; j <= top; j++ {
			sum += blend(i, j)
		}
	}

	fmt.Println(sum)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
