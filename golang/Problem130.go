package main

import (
	"euler"
	"fmt"
	"time"
)

const max = 25

func a(n int) int {
	remainder := 1
	i := 1

	for ; remainder != 0; i++ {
		remainder *= 10
		remainder++
		remainder %= n
	}

	return i
}

func main() {
	starttime := time.Now()

	count := 1
	sum := 0

	for n := 3; count <= max; n += 2 {
		if n%5 == 0 {
			//ensure coprime (could do without explicit test)
			n += 2
		}

		if !euler.IsPrime(int64(n)) && (n-1)%a(n) == 0 {
			fmt.Println(n)
			sum += n
			count++
		}

	}

	fmt.Println(sum)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
