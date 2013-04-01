package main

import (
	"euler"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	for i := int64(1); i < 100000000000; i++ {

		answer, ok := euler.IntSqrt((2 * i * i) - (2 * i) + 1)
		if ok && answer%2 == 1 {
			fmt.Println(i)
		}
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
