package main

import (
	"euler"
	"fmt"
	"time"
)

func zRep(n int) []int {
	var i int
	for i = 1; euler.Fibonacci(i) <= int64(n); i++ {
	}

	end := int(euler.Fibonacci(i - 1))
	if end == n {
		return []int{end}
	}

	return append(zRep(n-end), end)
}

var zrlmemo = make(map[int]int)

func zRepLength(n int) int {
	if ans, ok := zrlmemo[n]; ok {
		return ans
	}

	var i int
	for i = 1; euler.Fibonacci(i) <= int64(n); i++ {
	}

	end := int(euler.Fibonacci(i - 1))
	ans := 1 + zRepLength(n-end)
	zrlmemo[n] = ans

	return zRepLength(n)
}

func ZRLInit() {
	for i := 0; i < 31; i++ {
		zrlmemo[int(euler.Fibonacci(i))] = 1
	}
}

func rng(a, b int) int {
	total := 0
	for i := 1; i <= b; i++ {

		total += zRepLength(i)
	}
	return total
}

func main() {
	starttime := time.Now()

	ZRLInit()

	for i := 2; i < 25; i++ {
		fmt.Printf("%d-%-5d\t%d\n", 1, euler.Fibonacci(i), rng(1, int(euler.Fibonacci(i))))
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
