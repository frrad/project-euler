package main

import (
	"fmt"
	"time"
)

var memo map[[2]int]int

func ways(coins []int, amount int) (answer int) {
	if len(coins) == 1 {
		return 1
	}

	if answer, ok := memo[[2]int{len(coins), amount}]; ok {
		return answer
	}

	for i := 0; i <= amount/coins[0]; i++ {
		answer += ways(coins[1:], amount-(coins[0]*i))
	}

	memo[[2]int{len(coins), amount}] = answer

	return
}

func main() {
	starttime := time.Now()

	//Memoization is overkill
	memo = make(map[[2]int]int)

	fmt.Println(ways([]int{200, 100, 50, 20, 10, 5, 2, 1}, 200))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
