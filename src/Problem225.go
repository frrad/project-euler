package main

import (
	"fmt"
	"time"
)

var memo map[[2]int][3]int

func trib(index, modulus int) [3]int {
	if index == 1 {
		return [3]int{1, 1, 1}
	}

	if answer, ok := memo[[2]int{index, modulus}]; ok {
		return answer
	}

	answer := trib(index-1, modulus)

	answer[0], answer[1], answer[2] = answer[1], answer[2], (answer[0]+answer[1]+answer[2])%modulus

	memo[[2]int{index, modulus}] = answer

	return answer
}

func divides(p int) bool {

	for i := 2; trib(i, p)[0] != 0; i++ {
		if trib(i, p) == [3]int{1, 1, 1} {
			return false
		}

	}

	return true

}

func main() {
	starttime := time.Now()
	memo = make(map[[2]int][3]int)

	count := 0

	for i := 1; i < 10000; i += 2 {

		if !divides(i) {
			count++
			//fmt.Println(i, count)

		}

		if count == 124 {
			fmt.Println(i)
			break
		}
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
