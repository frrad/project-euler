package main

import (
	"euler"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	fmt.Println(euler.IntPartitions(100) - 1)

	fmt.Println("Elapsed time:", time.Since(starttime))

}
