package main

import (
	"fmt"
	"math/big"
	"time"
)

const top int64 = 1000000000

func accum(to int64, isOdd bool, eval func(int64, int64) bool) *big.Int {
	point := int64(0)
	store := big.NewInt(0)
	value := big.NewInt(0)

	for i := int64(0); i <= to; i++ {
		for eval(i, point+1) {
			point++
		}
		store.Add(store, value.SetInt64(point))
	}

	store.Mul(store, big.NewInt(4))
	if isOdd {
		store.Add(store, big.NewInt(-2*point))

	}
	return store
}

func main() {
	starttime := time.Now()

	omega := top / 4

	ans := big.NewInt(omega)
	ans.Mul(ans, ans)
	ans.Mul(ans, big.NewInt(24))

	aLength := omega + 1
	aOdd := aLength%2 == 1
	a := func(i, x int64) bool { return x*x < i*omega-i*i }
	ans.Add(ans, accum((aLength+1)/2-1, aOdd, a))

	bLength := omega
	bOdd := bLength%2 == 1
	b := func(i, x int64) bool { return 2*x*x-2*x+1 < -2*i-2*i*i+omega+2*i*omega }
	ans.Add(ans, accum((bLength+1)/2-1, bOdd, b))

	fmt.Printf("%s\n", ans.String())

	fmt.Println("Elapsed time:", time.Since(starttime))
}
