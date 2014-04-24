package main

import (
	"fmt"
	"time"
)

func checker(n uint64) string {
	if n == 1 {
		return ""
	}

	if n%3 == 0 {
		return "D" + checker(n/3)
	}

	if n%3 == 1 {
		return "U" + checker((4*n+2)/3)
	}

	return "d" + checker((2*n-1)/3)
}

func main() {
	starttime := time.Now()

	fmt.Println(checker(2))
	fmt.Println(checker(4))
	fmt.Println(checker(8))
	fmt.Println(checker(16))
	fmt.Println(checker(32))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
