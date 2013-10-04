package main

import (
	"euler"
	"fmt"

	"time"
)

const (
	top  = 50000000
	pTop = 4157407
	// pTop is PrimePi[Floor[Sqrt[2]*50*10^6]]
)

func main() {
	starttime := time.Now()

	fmt.Println("Building Cache")
	euler.PrimeCache(pTop)
	fmt.Println("Built:", time.Since(starttime))

	seive := [top + 1]bool{}
	for i := 2; i < top+1; i++ {
		seive[i] = true
	}

	for i := int64(3); i < pTop; i++ {
		p := euler.Prime(i)

		residue := (p + 1) / 2

		//first we solve 2 n^2 - 1 = p  in F_p
		//Too slow : use T-S
		if an1, an2, ok := euler.SqrtMod(residue, p); ok {
			fmt.Println(p)

			point := an1
			//This is slow too, don't check directly
			for euler.IsPrime(point*point*2 - 1) {
				point += p
			}
			for ; point < top+1; point += p {
				seive[point] = false
			}

			point = an2
			for euler.IsPrime(point*point*2 - 1) {
				point += p
			}
			for ; point < top+1; point += p {
				seive[point] = false
			}

		}

	}

	count := 0
	for i := range seive {
		if seive[i] {
			count++
		}
	}
	fmt.Println(count)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
