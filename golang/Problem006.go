package main

import (
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	top := 100
	sum := 0
	squaresum := 0

	for i := 1; i <= top; i++ {
		sum += i
		squaresum += i * i
	}

	fmt.Println(sum*sum - squaresum)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
