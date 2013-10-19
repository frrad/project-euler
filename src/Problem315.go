package main

import (
	"fmt"
	"time"
)

var tab map[int][7]bool

func shew(an int) {
	state := tab[an]
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
		return sum(tab[an])
	}

	base, rest := an%10, an/10

	return sum(tab[base]) + cost(rest)

}

func sum(a [7]bool) (ans int) {
	for i := 0; i < 7; i++ {
		if a[i] {
			ans++
		}
	}
	return
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
	tab[7] = [7]bool{true, false, true, false, false, true, false}
	tab[8] = [7]bool{true, true, true, true, true, true, true}
	tab[9] = [7]bool{true, true, true, true, false, true, true}

	for i := 0; i < 100; i++ {
		fmt.Println(i, cost(i))
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
