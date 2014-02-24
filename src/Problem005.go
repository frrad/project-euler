package main

import (
	"euler"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	answer := int64(1)

	for i := int64(1); i <= 20; i++ {
		answer = euler.LCM(i, answer)
	}

	fmt.Println(answer)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
