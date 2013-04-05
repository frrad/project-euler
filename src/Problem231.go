package main

import (
	"euler"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	N := int64(20000000)
	K := int64(15000000)

	factors := make(map[int64]int64)

	for n := N; n > N-K; n-- {
		nfactors := euler.Factors(n)
		for i := 0; i < len(nfactors); i++ {
			factors[nfactors[i][0]] += nfactors[i][1]
		}
		fmt.Println(n)
	}

	for k := K; k >= 2; k-- {

		kfactors := euler.Factors(k)
		for i := 0; i < len(kfactors); i++ {
			factors[kfactors[i][0]] -= kfactors[i][1]
		}
		fmt.Println(k)

	}

	fmt.Println(factors)

	answer := int64(0)

	for prime, multiplicity := range factors {

		for i := int64(0); i < multiplicity; i++ {
			answer += prime
		}

	}

	fmt.Println(answer)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
