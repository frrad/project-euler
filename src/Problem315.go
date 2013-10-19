package main

import (
	"euler"
	"fmt"
	"time"
)

var tab map[int][7]bool

func shew(state [7]bool) {
	if state[0] {
		fmt.Println(" -")
	}
	if state[1] && state[2] {
		fmt.Println("| |")
	} else if state[1] {
		fmt.Println("|")
	} else if state[2] {
		fmt.Println("  |")
	}

	if state[3] {
		fmt.Println(" -")
	}

	if state[4] && state[5] {
		fmt.Println("| |")
	} else if state[4] {
		fmt.Println("|")
	} else if state[5] {
		fmt.Println("  |")
	}

	if state[6] {
		fmt.Println(" -")
	}
}

func cost(an int) int {
	if an < 10 {
		// fmt.Println("Cost of", an, "is", sum(tab[an]))
		return sum(tab[an])
	}

	base, rest := an%10, an/10

	// fmt.Println("Cost of", base, rest, "is", sum(tab[base]), "+", cost(rest))
	return sum(tab[base]) + cost(rest)

}

func trans(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		if a == 0 {
			return 0
		}
		return cost(a)
	}

	a, aend := a%10, a/10
	b, bend := b%10, b/10

	return sum(XOR(a, b)) + trans(aend, bend)

}

func sum(a [7]bool) (ans int) {
	for i := 0; i < 7; i++ {
		if a[i] {
			ans++
		}
	}
	return
}

func XOR(a, b int) [7]bool {
	A, B := tab[a], tab[b]
	c := [7]bool{}
	for i := 0; i < 7; i++ {
		c[i] = zor(A[i], B[i])
	}
	return c
}

func zor(a, b bool) bool {
	if a == b {
		return false
	}
	return true
}

func next(an int) int {
	if an < 10 {
		return an
	}
	return next(an/10) + an%10
}

func rtChain(an int) []int {
	slice := make([]int, 1)
	slice[0] = an

	for an != next(an) {
		slice = append(slice, next(an))
		an = next(an)
	}

	return slice
}

func bad(chain []int) (ans int) {
	for i := 0; i < len(chain); i++ {
		ans += cost(chain[i])
		// fmt.Println(cost(chain[i]))
	}
	return 2 * ans
}

func good(chain []int) (ans int) {
	ans += cost(chain[0]) + cost(chain[len(chain)-1])
	for i := 0; i < len(chain)-1; i++ {
		ans += trans(chain[i], chain[i+1])
	}
	return
}

func diff(an int) int {
	chain := rtChain(an)
	return bad(chain) - good(chain)
}
func main() {
	starttime := time.Now()

	tab = make(map[int][7]bool)
	tab[0] = [7]bool{true, true, true, false, true, true, true}
	tab[1] = [7]bool{false, false, true, false, false, true, false}
	tab[2] = [7]bool{true, false, true, true, true, false, true}
	tab[3] = [7]bool{true, false, true, true, false, true, true}
	tab[4] = [7]bool{false, true, true, true, false, true, false}
	tab[5] = [7]bool{true, true, false, true, false, true, true}
	tab[6] = [7]bool{true, true, false, true, true, true, true}
	tab[7] = [7]bool{true, true, true, false, false, true, false}
	tab[8] = [7]bool{true, true, true, true, true, true, true}
	tab[9] = [7]bool{true, true, true, true, false, true, true}

	sum := 0

	euler.PrimeCache(20000000)
	fmt.Println("Built Cache")

	start, end := euler.PrimePi(10000000)+1, euler.PrimePi(20000000)

	fmt.Println(start, end)

	fmt.Println(euler.Prime(start), euler.Prime(end))

	for i := start; i <= end; i++ {
		// fmt.Println(i)
		sum += diff(int(euler.Prime(i)))
	}

	fmt.Println(sum)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
