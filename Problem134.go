package main

import (
	"./euler"
	"fmt"
	"time"
)

func pair(p1, p2 int64, c chan int64) {
	mod := int64(10)

	for ; mod < p1; mod *= 10 {
	}

	for try := int64(0); ; try += p2 {
		if try%mod == p1 {
			c <- try
			return
		}

	}

}

func main() {
	starttime := time.Now()

	euler.PrimeCache(1000000)

	c := make(chan int64)

	count := 0
	for i := int64(3); euler.Prime(i) < 1000000; i++ {
		go pair(euler.Prime(i), euler.Prime(i+1), c)
		count++

	}

	total := int64(0)
	for i := 0; i < count; i++ {
		total += <-c
	}
	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))

}
