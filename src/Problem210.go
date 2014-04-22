package main

import (
	"fmt"
	"math/big"
	"time"
)

const top int64 = 1000000000

func accum(from, to int64, store *big.Int, isOdd bool, eval func(int64, int64) bool) {
	point := int64(0)
	for i := from; i <= to; i++ {
		for eval(i, point+1) {
			point++
		}
		//		fmt.Printf("Found %d->%d\n", i, point)
		value := big.NewInt(point)
		store.Add(store, value)
		store.Add(store, value)
		store.Add(store, value)
		store.Add(store, value)
	}
	if isOdd {
		store.Add(store, big.NewInt(point*-1))
		store.Add(store, big.NewInt(point*-1))
	}
	return
}

func main() {
	starttime := time.Now()

	omega := top / 4

	ans := big.NewInt(omega)
	ans.Mul(ans, ans)
	ans.Mul(ans, big.NewInt(24))

	trashL := omega + 1
	trashOdd := trashL%2 == 1
	trashcan := func(i, x int64) bool {
		return x*x < i*omega-i*i
	}

	hipL := omega - 1 + 1
	hOdd := hipL%2 == 1
	hippo := func(i, x int64) bool {
		return (2*x-1)*(2*x-1) < 2*omega*(2*i+1)-(2*i+1)*(2*i+1)
	}

	accum(0, (trashL+1)/2-1, ans, trashOdd, trashcan)
	accum(0, (hipL+1)/2-1, ans, hOdd, hippo)

	fmt.Printf("%s\n", ans.String())

	fmt.Println("Elapsed time:", time.Since(starttime))
}
