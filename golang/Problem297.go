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

func main() {
	starttime := time.Now()

	ZRLInit()

	total := 0
	for i := 1; i < 1000000; i++ {
		//fmt.Printf("%d %d\n", i, zRepLength(i))
		total += zRepLength(i)
	}

	fmt.Printf("%d\n", total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
