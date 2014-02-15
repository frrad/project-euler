package main

import (
	"fmt"
	"time"
)

var borMemo map[[2]uint64]uint64

//returns the number of subrectangles in a x b grid
//with the boring orientation
func boring(a, b uint64) uint64 {

	if ans, ok := borMemo[[2]uint64{a, b}]; ok {
		return ans
	}

	if a > 1 {
		//we decompose into configurations which avoid the top row
		//those which avoid the bottom row, and those of full height
		borMemo[[2]uint64{a, b}] = 2*boring(a-1, b) - boring(a-2, b) + boring(1, b)
		return boring(a, b)
	}

	if a == 1 && b > 1 {
		borMemo[[2]uint64{a, b}] = b + boring(a, b-1)
		return boring(a, b)
	}

	if a <= 0 || b <= 0 {
		return 0
	}

	//a==b==1
	return 1

}

//think of sideways grid as normal one sans corners
func weird(a, b uint64) uint64 {
	return 0
}

func main() {
	starttime := time.Now()

	borMemo = make(map[[2]uint64]uint64)

	for i := uint64(1); i < 100; i++ {
		for j := uint64(1); j < 100; j++ {
			fmt.Println(j, i, "=", boring(i, j))
		}
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
