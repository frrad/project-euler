package main

import (
	"fmt"
	"time"
)

const mod = 20092010

var memo map[uint64]int

func fib(a uint64) int {
	if a <= 1999 {
		return 1
	}

	if answer, ok := memo[a]; ok {
		return answer
	}

	//fmt.Println(a)

	answer := (fib(a-2000) + fib(a-1999)) % mod

	memo[a] = answer
	return answer

}

//this is quite slow
func main() {
	starttime := time.Now()

	memo = make(map[uint64]int)

	for i := uint64(10000000); i < 12000000; i++ {
		fmt.Print(fib(i), ",")
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
