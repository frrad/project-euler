package main

import (
	"fmt"
	"math/big"
	"time"
)

const tablesize = 100000000

var table [tablesize]*big.Int

//Recurrence equation for partition function, due to Euler
func P(n int) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	if n < 0 {
		return big.NewInt(0)
	}

	if n < tablesize && table[n] != nil {
		return table[int(n)]
	}

	sum := big.NewInt(0)

	for k := 1; k <= n; k++ {
		summand := new(big.Int)
		if k%2 == 0 {
			summand = big.NewInt(-1)
		} else {
			summand = big.NewInt(1)
		}

		next := new(big.Int)
		next.Add(P(n-(k*(3*k-1)/2)), P(n-(k*(3*k+1)/2)))

		summand.Mul(summand, next)

		sum.Add(sum, summand)
	}

	if n < tablesize {
		table[int(n)] = sum
	}

	return sum
}

func main() {
	starttime := time.Now()

	i := 1

	mod := new(big.Int)
	million := big.NewInt(1000000)
	zero := big.NewInt(0)

	for mod.Mod(P(i), million).Cmp(zero) != 0 {

		fmt.Println(i, P(i))

		i += 1

	}

	fmt.Println(i, P(i))

	fmt.Println("Elapsed time:", time.Since(starttime))

}
