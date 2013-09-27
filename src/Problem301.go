package main

import (
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	count := 0
	for x := 0; x < 1<<30; x++ {

		if x&(x<<1) == 0 {
			count++
		}

	}

	fmt.Println(count)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
