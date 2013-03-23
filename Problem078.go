package main

import (
	"fmt"
	"math/big"
	"time"
)

const (
	tablesize = 10000000
	mod       = 1000000
)

var table [tablesize]int

//Recurrence equation for partition function, due to Euler
func P(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}

	if n < tablesize && table[n] != 0 {
		return table[n]
	}

	sum := 0

	for k := 1; k <= n; k++ {
		var summand int
		if k%2 == 0 {
			summand = -1
		} else {
			summand = 1
		}

		summand *= P(f(n, k)) + P(g(n, k))

		sum += summand
	}

	if n < tablesize {
		table[n] = (sum + 10*mod) % mod
	}

	return (sum + mod) % mod
}

func f(n, k int) int {
	N := big.NewInt(int64(n))
	K := big.NewInt(int64(k))

	a, b, c, d, e := new(big.Int), new(big.Int), new(big.Int), new(big.Int), new(big.Int)

	a.Mul(K, big.NewInt(3))
	b.Sub(a, big.NewInt(1))
	c.Mul(b, K)
	d.Div(c, big.NewInt(2))
	e.Sub(N, d)

	return int(e.Int64())

}

func g(n, k int) int {
	N := big.NewInt(int64(n))
	K := big.NewInt(int64(k))

	a, b, c, d, e := new(big.Int), new(big.Int), new(big.Int), new(big.Int), new(big.Int)

	a.Mul(K, big.NewInt(3))
	b.Add(a, big.NewInt(1))
	c.Mul(b, K)
	d.Div(c, big.NewInt(2))
	e.Sub(N, d)

	return int(e.Int64())

}

func main() {
	starttime := time.Now()
	i := 2

	for P(i) != 0 {

		fmt.Println(i, P(i))
		i++

	}

	fmt.Println(P(i))

	fmt.Println("Elapsed time:", time.Since(starttime))

}
