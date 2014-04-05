package main

import (
	"euler"
	"fmt"
	"math/big"
	"time"
)

const (
	top   = 100
	moved = 22
)

func main() {
	starttime := time.Now()

	//how many primes less than top
	primes := int(euler.PrimePi(top))
	composites := top - primes
	fixedP := primes - moved

	accumulate := big.NewInt(0)

	for fixedC := 0; fixedC <= top-primes; fixedC++ {
		derange := top - fixedC - fixedP

		fmt.Printf("We fix %d composites, %d primes, derange %d (rest).\n", fixedC, fixedP, derange)

		choice := euler.BigChoose(primes, fixedP)
		choice.Mul(choice, euler.BigChoose(composites, fixedC))
		choice.Mul(choice, euler.BigSubfactorial(derange))

		accumulate.Add(accumulate, choice)
	}

	total := euler.BigFactorial(top)

	fmt.Printf("%d\n--------------------\n%d\n", accumulate, total)

	num, den := new(big.Rat), new(big.Rat)
	num.SetInt(accumulate)
	num.Quo(num, den.SetInt(total))

	fmt.Printf("%s\n", num.FloatString(12))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
