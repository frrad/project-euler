package main

import (
	"euler"
	"fmt"
	"time"
)

const (
	notExceeding = 100000000
	half         = notExceeding / 2
)

func validate(n int64) bool {

	for d := int64(2); d*d <= n; d++ {
		if n%d == 0 && !euler.IsPrime(d+(n/d)) {
			return false
		}
	}

	return true
}

func main() {
	starttime := time.Now()

	euler.PrimeCache(half)
	fmt.Println("Built Prime Cache...")

	check := []int64{1}

	for i := int64(1); i <= euler.PrimePi(half); i++ {
		prime := euler.Prime(i)
		candidate := (prime - 2) * 2
		if euler.IsPrime(candidate + 1) {
			check = append(check, candidate)
		}
	}

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(check[i])
	// }

	fmt.Printf("There are %d initial candidates.\n", len(check))

	total := int64(0)

	for i, can := range check {
		if validate(can) {
			// fmt.Println(can)
			total += can
		}
		if i%10000 == 0 {
			fmt.Printf("Checked %d/%d\n", i, len(check))
		}
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))

}
