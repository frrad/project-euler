package main

import (
	"euler"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func isSpecial(set []int) bool {
	set = euler.SortInts(set)
	if check2(set) && check1(set) {
		return true
	}

	return false
}

func binary(index, size int) []int {
	if size == 0 {
		return []int{}
	}
	return append(binary(index/2, size-1), index%2)
}

func check1(in []int) bool {
	size := len(in)

	set := make(map[int]bool)
	for i := 1; i < int(euler.Exp2(size)); i++ {
		memb := binary(i, size)
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
	size := len(in)
	sum1, sum2 := in[size-1], 0
	for i := 0; i < size/2; i++ {
		sum2 += in[i]
		sum1 += in[size-i-2]
		if sum1 < sum2 {
			return false
		}
	}

	return true
}

func sum(set []int) (total int) {
	for _, x := range set {
		total += x
	}
	return
}

func main() {
	starttime := time.Now()

	data := euler.Import("../problemdata/sets.txt")
	sets := make([][]int, len(data))

	for i, line := range data {
		sets[i] = make([]int, 0)
		for _, word := range strings.Split(line, ",") {
			number, _ := strconv.Atoi(word)
			sets[i] = append(sets[i], number)
		}
	}

	total := 0

	for _, set := range sets {
		if isSpecial(set) {
			total += sum(set)
		}
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
