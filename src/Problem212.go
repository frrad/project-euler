package main

import (
	"fmt"
	"time"
)

const cubes = 100

var randMemo map[int]int

func S(k int) int {
	if ans, ok := randMemo[k]; ok {
		return ans
	}

	if k <= 55 {
		randMemo[k] = (100003 - (200003 * k) + (300007 * k * k * k)) % 1000000
		return S(k)
	}

	randMemo[k] = (S(k-24) + S(k-55)) % 1000000
	return S(k)
}

func sort(indices []int, val map[int]int) {

}

//Idea: Find disjoint figures, then size those using
//[][][]bool with specified endpoints
func main() {
	starttime := time.Now()

	randMemo = make(map[int]int)

	xstart, xend := make(map[int]int), make(map[int]int)
	zstart, zend := make(map[int]int), make(map[int]int)

	for i := 1; i <= cubes; i++ {
		x, y, z := S(6*i-5)%10000, S(6*i-4)%10000, S(6*i-3)%10000
		dx, dy, dz := 1+(S(6*i-2)%399), 1+(S(6*i-1)%399), 1+(S(6*i)%399)

		//After init, everything is zero indexed
		xstart[i-1], xend[i-1] = x, x+dx
		ystart[i-1], yend[i-1] = y, y+dy
		zstart[i-1], zend[i-1] = z, z+dz

		//fmt.Println(x, y, z, "\t", dx, dy, dz)
	}

	fmt.Println(xstart)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
