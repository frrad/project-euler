package main

import (
	//	"euler"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	count := 1
	s := int64(1)
	s2 := int64(1)
	it := int64(8)

	for n := int64(1); n < 1000000000000; n += 1 {

		for s2 < it {
			s2 += 1 + 2*s
			s++
		}

		if it == s2 {
			fmt.Println(count, "\t", n)
			count++
		}

		it += 7 + 10*n

	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
