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

func interesting(l, w int) uint64 {
	tri1, tri2 := l-2, w-2
	square := tri1 + tri2 + 2
	if tri1 < 0 {
		tri1 = 0
	}
	if tri2 < 0 {
		tri2 = 0
	}

	tri1, tri2 = sort(tri1, tri2)

	// fmt.Println("tri1", tri1, "tri2", tri2, "square", square)

	grid := make([][]bool, square)
	for i := range grid {
		grid[i] = make([]bool, square)
	}

	offset1 := square - tri1
	offset2 := square - 1

	for i := 0; i < tri1; i++ {
		for j := 0; j <= i; j++ {
			grid[i+offset1][offset2-j] = true
		}
	}

	offset1 = tri1 - 1

	for i := 0; i < tri1; i++ {
		for j := 0; j <= i; j++ {
			grid[offset1-i][j] = true
		}
	}

	offset1 = tri1 + 2

	for i := 0; i < tri2; i++ {
		for j := 0; j <= i; j++ {
			grid[i+offset1][j] = true
		}
	}

	offset1 = tri2 - 1
	offset2 = tri2 + tri1 + 1

	for i := 0; i < tri2; i++ {
		for j := 0; j <= i; j++ {
			grid[offset1-i][offset2-j] = true
		}
	}

	// show(grid)

	tally := uint64(0)

	for i := 0; i < square; i++ {
		for j := 0; j < square; j++ {
			if !grid[i][j] {
				for k := i; k < square; k++ {
					for l := j; l < square; l++ {

						intermezzo := 0

						for m := i; m <= k; m++ {
							for n := j; n <= l; n++ {
								if !grid[m][n] {
									intermezzo++
								}
							}
						}

						if intermezzo == (k-i+1)*(l-j+1) {
							tally++
						}

					}
				}

			}

		}
	}

	return tally
}

func show(grid [][]bool) {
	for _, line := range grid {
		for _, val := range line {
			if val {
				fmt.Print("X")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println("")
	}

}

func predict(a, b int) int64 {
	b, a = sort(a, b)

	if a == 1 {
		return int64(b - 1)
	}

	if a > b {
		panic("adf")
	}

	offset := 0

	factors := euler.Factors(int64(a))

	if factors[0][0] > 3 {
		offset = -1
	}

	return int64(offset - (a / 2) + (a*a)/6 - (2*a*a*a*a)/3 - (a*b)/3 + (4*a*a*a*b)/3)

}

func blend(a, b int) uint64 {
	bore := uint64(0)
	if predict(a, b) < 0 {
		fmt.Println("anus")
		bore = interesting(a, b)
	} else {
		bore = uint64(predict(a, b))
	}
	return boring(a, b) + bore
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

	// for i := 1; i < 50; i++ {
	// 	a, b := uint64(predict(i, i)), interesting(i, i)
	// 	if a != b {
	// 		fmt.Println(i, ":", a, b)
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }

	fmt.Println("Elapsed time:", time.Since(starttime))
}
