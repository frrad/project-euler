package main

import (
	"euler"
	"fmt"
	"math/big"
	"time"
)

func isLychrel(number int) bool {
	current := big.NewInt(int64(number))
	reverse := big.NewInt(0)

	for i := 0; i < 50; i++ {
		reverse.SetString(euler.StringReverse(current.String()), 10)
		current.Add(current, reverse)
		if euler.IsStringPalindrome(current.String()) {
			return false
		}

	}

	return true
}

func main() {
	starttime := time.Now()

	total := 0
	for i := 0; i < 10000; i++ {
		if isLychrel(i) {
			//fmt.Printf("%d\n", i)
			total++
		}
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
