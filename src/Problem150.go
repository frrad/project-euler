package main

import (
	"fmt"
	"time"
)

var S map[int]int
var triMem, rowMem map[[3]int]int64
var srcMem map[[2]int]int

func triangle(i, j, depth int) int64 {

	if ans, ok := triMem[[3]int{i, j, depth}]; ok {
		return ans
	}

	if depth == 1 {
		return int64(tri(i, j))
	}

	ans := triangle(i, j, depth-1) + row(i+depth-1, j, depth)

	triMem[[3]int{i, j, depth}] = ans

	return ans
}

func row(i, j, depth int) int64 {
	if ans, ok := rowMem[[3]int{i, j, depth}]; ok {
		return ans
	}

	if depth == 1 {
		return int64(tri(i, j))
	}

	ans := row(i, j, depth-1) + int64(tri(i, j+depth-1))

	rowMem[[3]int{i, j, depth}] = ans

	return ans
}

func tri(i, j int) int {
	if j > i {
		panic("wuttt")
	}

	if ans, ok := srcMem[[2]int{i, j}]; ok {
		return ans
	}

	j++
	ans := S[j+(i*i+i)/2]
	j--

	srcMem[[2]int{i, j}] = ans
	return ans
}

func main() {
	starttime := time.Now()

	S = make(map[int]int)
	triMem, rowMem = make(map[[3]int]int64), make(map[[3]int]int64)
	srcMem = make(map[[2]int]int)

	t := int64(0)
	for k := 1; k <= 500500; k++ {
		t = (615949*t + 797807) % (1 << 20)
		S[k] = int(t) - (1 << 19)
		//S[k] = k
	}

	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(tri(i, j), "\t")
		}
		fmt.Print("\n")
	}

	min := int64(0)

	for depth := 1; depth < 1000; depth++ {

		for i := 0; i+depth-1 < 1000; i++ {
			for j := 0; j <= i; j++ {

				if triangle(i, j, depth) < min {

					min = triangle(i, j, depth)
					fmt.Println(min, "\t", i, "\t", j, "\t", depth)
				}

			}
		}

		for key, _ := range triMem {
			d := key[2]
			if d < depth-1 {
				delete(triMem, key)
			}
		}

		for key, _ := range rowMem {
			d := key[2]
			if d < depth-1 {
				delete(rowMem, key)
			}
		}

	}

	fmt.Println(min)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
