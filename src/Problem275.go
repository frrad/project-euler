package main

import (
	"fmt"
	"time"
)

const (
	order = 18
	top   = 1 << order
)

var supmemo map[int][]int

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

func com(dist []int) (ans int) {
	for i, line := range dist {
		ans += count(line) * i
	}
	return
}

func count(an int) (bits int) {

	for an != 0 {
		bits += 1 & an
		an >>= 1
	}
	return
}

func sups(an int) []int {

	if ans, ok := supmemo[an]; ok {
		return ans
	}

	ans := make([]int, 0)
	for i := 0; i < top; i++ {
		if supports(an, i) {
			ans = append(ans, i)
		}
	}

	supmemo[an] = ans
	return ans

}

func main() {
	starttime := time.Now()

	supmemo = make(map[int][]int)

	print(top - 1)

	//how tall above plinth
	balance := uint(2)

	seq := [][][]int{[][]int{[]int{1<<balance - 1}}}

	blocks := 2

	for depth := 2; depth < 4 || len(seq[depth-3]) != len(seq[depth-2]); depth++ {
		seq = append(seq, [][]int{})
		short := seq[depth-2]
		for _, list := range short {
			trailing := list[depth-2]

			blocksUsed := 0
			for _, line := range list {
				blocksUsed += count(line)
			}

			for _, possible := range sups(trailing) {
				if count(possible)+blocksUsed-int(balance) <= blocks && !((count(possible)+blocksUsed-int(balance) < blocks) && possible == 0) {
					longer := make([]int, depth)
					copy(longer[:depth-1], list)
					longer[depth-1] = possible
					seq[depth-1] = append(seq[depth-1], longer)
				}
			}

			//fmt.Println(trailing)

		}
	}

	list := seq[len(seq)-1]

	for _, config := range list {
		fmt.Println("COM:", com(config))
		for _, line := range config {
			print(line)
		}
		fmt.Println("======")
	}

	for _, butts := range seq {

		fmt.Println(len(butts))
	}

	//fmt.Println(seq)

	fmt.Println("Elapsed time:", time.Since(starttime))

}
