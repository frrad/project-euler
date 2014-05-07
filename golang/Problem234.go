package main

import (
	"euler"
	"fmt"
	"math"
	"time"
)

const top = 999966663333

func main() {
	starttime := time.Now()

	//value of largest prime we should be concerned with (
	topPrime := int64(math.Ceil(math.Sqrt(top)))
	euler.PrimeCache(topPrime)

	//index of our last prime
	lastPrime := euler.PrimePi(topPrime) + 1

	total := int64(0)

	for i := int64(1); i < lastPrime; i++ {
		a, b := euler.Prime(i), euler.Prime(i+1)
		//fmt.Println(a, b)

		for square := a*a + a; square < b*b; square += a {
			if square%b != 0 && square <= top {
				total += square
				//fmt.Println(square)
			}
		}

		for square := b*b - b; square > a*a; square -= b {
			if square%a != 0 && square <= top {
				total += square
				//fmt.Println(square)
			}
		}

	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
