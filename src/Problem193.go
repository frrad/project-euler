package main

import (
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	seive := [1 << 50]bool{}
	//That's a bad idea!
	//Insted: inclusion / exclusion on the set of square-divisible numbers

	fmt.Println("Elapsed time:", time.Since(starttime))
}
