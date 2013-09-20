package main

import (
	"euler"
	"fmt"
	"time"
)

const mod = 1307674368000

func f(n int64, x int) int64 {
	ans := int64(0)
	for i := int64(0); i <= n; i++ {
		term := fib(i)
		for j := 0; j < int(i); j++ {
			term *= int64(x)
			term %= mod
		}
		ans += term
		ans %= mod
	}
	return ans
}

func fib(n int64) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if ans, ok := fibTable[n]; ok {
		return ans
	}
	ans := (fib(n-1) + fib(n-2)) % mod
	fibTable[n] = ans
	return ans
}

var fibTable map[int64]int64

//(x - Fibonacci[n + 1] (x^(n + 1) - x^(n + 2)) - Fibonacci[n + 2] x^(n + 2))/(1 - x - x^2)
func main() {
	starttime := time.Now()

	fmt.Println(60 + 60 + euler.ChineseRemainder([]int64{2, 3, 1}, []int64{3, 4, 5}))

	fmt.Println("Elapsed time:", time.Since(starttime))
}