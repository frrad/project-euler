package main

import (
	"fmt"
	"math/big"
	"time"
)

const (
	top = 50000000
	bar = 50
)

func progress(percent float32) string {
	ans := "["
	length := float32(bar)
	for i := float32(0); i < length; i++ {
		if i/length < percent {
			ans += "X"
		} else {
			ans += " "
		}
	}
	ans += "]"
	return ans
}

func clear(n int) (ans string) {
	for i := 0; i < n; i++ {
		ans += "\b"
	}
	return
}

func main() {
	starttime := time.Now()

	two := big.NewInt(2)
	one := big.NewInt(1)

	validate := make(map[int64]bool)

	for i := int64(2); i < top; i++ {

		if i%(top/bar) == 0 {
			fmt.Printf("%s%s", clear(bar+2), progress(float32(i)/float32(top)))
		}

		val := big.NewInt(i)

		val.Mul(val, val)
		val.Mul(val, two)
		val.Sub(val, one)

		if val.ProbablyPrime(1) {
			validate[i] = true
		}

	}

	fmt.Printf("\nFirst pass found %d primes. (Elapsed Time: %s)\n", len(validate), time.Since(starttime))

	//max value of j found experimentally
	for j := 2; j < 3; j++ {

		for i := range validate {
			val := big.NewInt(i)

			val.Mul(val, val)
			val.Mul(val, two)
			val.Sub(val, one)

			if !val.ProbablyPrime(j) {
				delete(validate, i)
			}

		}

		fmt.Printf("Pass %d found %d primes. (Elapsed Time: %s)\n", j, len(validate), time.Since(starttime))
	}

	fmt.Println(len(validate))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
