package main

import (
	"fmt"
	"math/big"
	"time"
)

const top = 1000

func main() {
	starttime := time.Now()

	frac := big.NewRat(2, 1)
	n, d := frac.Num(), frac.Denom()
	table := big.NewInt(0)

	two := big.NewRat(2, 1)

	var total int

	for i := 0; i < top; i++ {
		frac.Add(two, frac.Inv(frac))

		if len(table.Add(n, d).String()) > len(n.String()) {
			total++
		}

	}

	fmt.Printf("%d\n", total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
