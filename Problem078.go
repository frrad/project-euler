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
	if n < tablesize && table[n] != nil {
		return table[int(n)]
	}
	if n <= 0 {
		fmt.Println("ERRORR")
		return nil
	}
	sum := big.NewInt(0)
	for k := 1; k <= n; k++ {
		if k%2 == 0 {
			if n-(k*(3*k-1)/2) < 0 && n-(k*(3*k+1)/2) < 0 {
				//do nothing
			} else if n-(k*(3*k-1)/2) < 0 && n-(k*(3*k+1)/2) >= 0 {
				sum.Sub(sum, P(n-(k*(3*k+1)/2)))
			} else if n-(k*(3*k-1)/2) >= 0 && n-(k*(3*k+1)/2) < 0 {
				sum.Sub(sum, P(n-(k*(3*k-1)/2)))
			} else {
				sum.Sub(sum, P(n-(k*(3*k-1)/2)))
				sum.Sub(sum, P(n-(k*(3*k+1)/2)))
			}
		} else {
			if n-(k*(3*k-1)/2) < 0 && n-(k*(3*k+1)/2) < 0 {
				//do nothing
			} else if n-(k*(3*k-1)/2) < 0 && n-(k*(3*k+1)/2) >= 0 {
				sum.Add(sum, P(n-(k*(3*k+1)/2)))
			} else if n-(k*(3*k-1)/2) >= 0 && n-(k*(3*k+1)/2) < 0 {
				sum.Add(sum, P(n-(k*(3*k-1)/2)))
			} else {
				sum.Add(sum, P(n-(k*(3*k-1)/2)))
				sum.Add(sum, P(n-(k*(3*k+1)/2)))
			}
		}
	}
	if n < tablesize {
		table[int(n)] = sum
	}
	return sum
}

func main() {
	starttime := time.Now()

	table[0] = big.NewInt(1)

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
