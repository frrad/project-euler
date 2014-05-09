package main

import (
	"euler"
	"fmt"
	"math/big"
	"time"
)

const (
	top    = 5000
	bottom = 1000
)

var (
	NCONST = euler.IntExp(10, 18)
	KCONST = euler.IntExp(10, 9)
)

func main() {
	starttime := time.Now()

	euler.PrimeCache(top)

	startP := euler.PrimePi(bottom) + 1
	stopP := euler.PrimePi(top)

	fmt.Println(euler.Prime(startP), euler.Prime(stopP))

	lookup := make([][2]int64, 0)
	for pindex := startP; pindex <= stopP; pindex++ {
		p := euler.Prime(pindex)
		val := euler.ChooseModP(NCONST, KCONST, p)

		lookup = append(lookup, [2]int64{p, val})
	}

	size := len(lookup)
	fmt.Println("built table with", size, "entries")

	total := big.NewInt(0)

	for i := 0; i < size; i++ {
		a1, p1 := lookup[i][1], lookup[i][0]

		for j := i + 1; j < size; j++ {
			a2, p2 := lookup[j][1], lookup[j][0]

			for k := j + 1; k < size; k++ {
				a3, p3 := lookup[k][1], lookup[k][0]

				a := []int64{a1, a2, a3}
				p := []int64{p1, p2, p3}

				crt := euler.BigChineseRemainder(a, p)

				total.Add(total, crt)
			}
		}
		fmt.Println(1+i, "of", size, "(", total, ")")
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
