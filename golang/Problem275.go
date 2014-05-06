package main

import (
	"fmt"
	"time"
)

const (
	order    = 6
	top      = 1 << order
	infinity = 99999999
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
	for i := 0; i <= top; i++ {
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
	trackMemo := make(map[[3]int]int)

	//how tall above plinth

	for balance := uint(1); balance <= uint(order); balance++ {

		for blocks := 0; blocks <= order-int(balance); blocks++ {

			minCOM := 0
			maxCOM := infinity
			dual := order - blocks - int(balance)
			if blocks > dual {
				for key := range trackMemo {
					if key[0] == int(balance) && key[1] == dual {
						COM := key[2]
						if COM < minCOM {
							minCOM = COM
						}
						if COM > maxCOM {
							maxCOM = COM
						}
					}
				}
			}

			seq := [][][]int{{{1<<balance - 1}}}

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
						if bAdd := count(possible); bAdd+blocksUsed-int(balance) <= blocks && !(bAdd+blocksUsed-int(balance) < blocks && possible == 0) {
							longer := make([]int, depth)
							copy(longer[:depth-1], list)
							longer[depth-1] = possible
							if com(longer) <= maxCOM {
								seq[depth-1] = append(seq[depth-1], longer)

							}

						}
					}

				}
			}

			list := seq[len(seq)-1]

			for _, config := range list {

				trackMemo[[3]int{int(balance), blocks, com(config)}]++
			}

		}
	}

	//fmt.Println(trackMemo)

	total := 0
	for key, multiplicity := range trackMemo {
		bal, block, com := key[0], key[1], key[2]
		dual := order - bal - block
		if dual < block {
			tweak := multiplicity * trackMemo[[3]int{bal, dual, com}]
			total += tweak
			if tweak != 0 {
				fmt.Printf("Bal: %d Blocks: %d COM: %d => %d\n", bal, block, com, tweak)
				fmt.Println(multiplicity)
			}
		}
		if dual == block {
			tweak := multiplicity * (multiplicity + 1) / 2
			total += tweak
			if tweak != 0 {
				fmt.Printf("Bal: %d Blocks: %d COM: %d => %d\n", bal, block, com, tweak)
			}
		}
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))

}
