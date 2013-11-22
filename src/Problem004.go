package main

import (
	"euler"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	max := int64(0)

	for i := int64(0); i < 1000; i++ {
		for j := int64(0); j < 1000; j++ {
			if euler.IsPalindrome(i*j) && i*j > max {
				max = i * j
			}
		}
	}

	fmt.Println(max)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
