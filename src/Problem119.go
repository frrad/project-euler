package main

import (
	"./euler"
	"fmt"
	"math/big"
	"time"
)

func seek(start, end int64, c chan int64, state chan bool) {
	target := big.NewInt(0)
	bigsum := big.NewInt(0)
	bigexp := big.NewInt(0)
	test := big.NewInt(0)

	//fmt.Println("Seeking", start, "->", end)
	for i := start; i < end; i++ {
		target.SetInt64(i)
		sum := int64(euler.DigitSum(i))
		bigsum.SetInt64(sum)
		if sum > 1 {
			bigexp.SetInt64(2)
			for exp := int64(2); test.Cmp(target) <= 0; exp++ {
				bigexp.SetInt64(exp)
				test.Exp(bigsum, bigexp, nil)
				if test.Cmp(target) == 0 {
					c <- i
				} else {
					//	fmt.Println(test, "is", bigsum, "to the", bigexp, "but it's not", target)
				}
			}
		}
		test.SetInt64(0)
	}
	//	fmt.Println("Sought", start, "->", end)
	state <- true

}

func main() {
	starttime := time.Now()

	c := make(chan int64)
	state := make(chan bool)
	patchsize := int64(10000000)

	routines := int64(4)

	place := int64(0)
	for ; place < routines*patchsize; place += patchsize {
		go seek(place, place+patchsize, c, state)
	}

	all := make(map[int64]bool)

	for {
		select {
		case x := <-c:
			//	fmt.Println(x)
			all[x] = true
		case <-state:
			go seek(place, place+patchsize, c, state)
			place += patchsize
			total := 0
			for i := range all {
				fmt.Println(i)
				total++
			}

			fmt.Println("-----------------", total, place)
		}
	}

	fmt.Println("Elapsed time:", time.Since(starttime))

}
