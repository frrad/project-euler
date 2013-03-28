package main

import (
	"./euler"
	"fmt"
	"time"
)

func seek(start int64, c chan int64) {
	for i := start; ; i++ {
		sum := int64(euler.DigitSum(i))
		if sum > 1 {

			power := sum * sum
			for exp := int64(2); power <= i; exp++ {
				power = euler.IntExp(sum, exp)
				if power == i {
					c <- i
				}
			}
		}
	}

}

func main() {
	starttime := time.Now()

	c, d := make(chan int64), make(chan int64)

	go seek(12, c)

	go seek(614656, d)

	for {
		select {
		case x := <-c:
			fmt.Println(x)
		case x := <-d:
			fmt.Println(x)
		}
	}

	fmt.Println("Elapsed time:", time.Since(starttime))

}
