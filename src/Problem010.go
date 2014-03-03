package main

import (
	"euler"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	euler.PrimeCache(2000000)

	total := int64(0)
	for i := int64(1); euler.Prime(i) < 2000001; i++ {
		total += euler.Prime(i)
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))

}
