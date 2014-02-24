package main

import (
	"fmt"
	"math/big"
	"time"
)

const top = 1000

func main() {
	starttime := time.Now()

	zero := big.NewInt(0)
	two := big.NewInt(2)
	ten := big.NewInt(10)
	tothe := big.NewInt(top)

	mod := big.NewInt(0)
	exp := big.NewInt(0)

	exp.Exp(two, tothe, nil)

	// fmt.Println(exp.String())
	sum := int64(0)

	for zero.Cmp(exp) < 0 {

		mod.Mod(exp, ten)
		exp.Quo(exp, ten)

		// fmt.Println(mod.String())
		sum += mod.Int64()

	}

	fmt.Println(sum)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
