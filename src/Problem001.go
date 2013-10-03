package main

import (
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
