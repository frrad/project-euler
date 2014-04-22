package main

import (
	"fmt"
	"time"
)

const top int64 = 1000000000

func accum(to int64, isOdd bool, eval func(int64, int64) bool) (store int64) {
	point := int64(0)

	for i := int64(0); i <= to; i++ {
		for eval(i, point+1) {
			point++
		}
		store += point
	}
	store *= 4

	if isOdd {
		store -= 2 * point
	}

	return store
}

func main() {
	starttime := time.Now()

	omega := top / 4
	ans := omega * omega * 24

	aLength := omega + 1
	aOdd := aLength%2 == 1
	a := func(i, x int64) bool { return x*x < i*omega-i*i }
	ans += accum((aLength+1)/2-1, aOdd, a)

	bLength := omega
	bOdd := bLength%2 == 1
	b := func(i, x int64) bool { return 2*x*x-2*x+1 < -2*i-2*i*i+omega+2*i*omega }
	ans += accum((bLength+1)/2-1, bOdd, b)

	fmt.Printf("%d\n", ans)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
