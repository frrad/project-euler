package main

import (
	"euler"
	"fmt"
	"time"
)

const size = 7
const start = 19
const maxgap = 11

func binary(bin, index, siz int) []int {
	if siz == 0 {
		return []int{}
	}

	return append(binary(bin, index/bin, siz-1), index%bin+1)
}

func test(i int) []int {
	ret := binary(maxgap, i, size)
	ret[0] += start
	for i := 1; i < size; i++ {
		ret[i] += ret[i-1]
	}
	return ret
}

func check1(in []int) bool {
	set := make(map[int]bool)
	for i := 1; i < int(euler.Exp2(size)); i++ {
		memb := binary(2, i, size)
		sum := 0
		for j := 0; j < size; j++ {
			if memb[j] == 1 {
				sum += in[j]
			}
		}
		if set[sum] {
			return false
		}
		set[sum] = true
	}

	return true

}

func check2(in []int) bool {
	sum1, sum2 := in[0], 0
	for i := 0; i < size/2; i++ {
		sum1 += in[i+1]
		sum2 += in[size-i-1]
		if sum1 < sum2 {
			return false
		}
	}

	return true
}

func total(in []int) (sum int) {
	for _, x := range in {
		sum += x
	}
	return

}

func main() {
	starttime := time.Now()

	best := 99999

	for i := 0; i < int(euler.IntExp(maxgap, size)); i++ {
		try := test(i)
		if total(try) < best && check2(try) && check1(try) {
			best = total(try)
			fmt.Println(test(i))
		}
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
