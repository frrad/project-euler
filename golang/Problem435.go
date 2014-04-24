package main

import (
	"fmt"
	"math/big"
	"time"
)

var trim int64

const (
	n   = 1000000000000000
	mod = 1307674368000
)

func divmod(a, b int64) (q, r int64) {
	return a / b, a % b
}

func fibFast(i int64) int64 {
	var a, b int64

	if i < 0 {
		a, b = 1, -1
		i = -i
	} else {
		a, b = 1, 0
	}

	var c, d int64
	i, n := divmod(i, 2)
	if n != 0 {
		c, d = a, b
	} else {
		c, d = 0, 1
	}

	for i > 0 {
		a, b = fibtimes(a, b, a, b)
		i, n = divmod(i, 2)
		if n != 0 {
			c, d = fibtimes(a, b, c, d)
		}
	}
	return c
}

func safeTimes(a, b int64) int64 {
	A, B := big.NewInt(a), big.NewInt(b)
	modulus := big.NewInt(mod * trim)
	A.Mul(A, B)
	A.Mod(A, modulus)
	return A.Int64()
}

func fibtimes(a, b, c, d int64) (int64, int64) {
	return safeTimes(a, c+d) + safeTimes(b, c), safeTimes(a, c) + safeTimes(b, d)
}

func powMod(x, n int64) int64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	half := powMod(x, n/2)
	whole := safeTimes(half, half)

	if n%2 == 0 {
		return whole
	}

	return safeTimes(x, whole)

}

// (Fibonacci[n + 1] x^(n + 1) + Fibonacci[n] x^(n + 2) - x)/(x^2 + x - 1)
func main() {
	starttime := time.Now()

	total := int64(0)

	for x := int64(0); x <= 100; x++ {

		denom := x*x + x - 1
		trim = denom

		numer := safeTimes(fibFast(n+1), powMod(x, n+1))
		numer += safeTimes(fibFast(n), powMod(x, n+2))
		numer -= x

		total += (numer / denom) % mod
		total %= mod

		fmt.Println(x, "\t", (numer/denom)%mod)

	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
