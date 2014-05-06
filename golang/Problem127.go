package main

import (
	"fmt"
	"time"
)

const max = 120000

var table map[int64]map[int64]bool
var rable map[int64]int64

func coprime(a, b int64) bool {
	for prime := range sig(a) {
		if sig(b)[prime] {
			return false
		}
	}
	return true
}

func sig(n int64) map[int64]bool {
	return table[n]
}

func populate() {
	for i := int64(2); i < max; i++ {
		if _, ok := table[i]; !ok { //prime

			for j := int64(1); i*j < max; j++ {
				if _, there := table[i*j]; !there {
					table[i*j] = make(map[int64]bool)
				}
				table[i*j][i] = true

			}
		}
	}
}

func rad(a, b, c int64) (r int64) {
	if b == 1 && c == 1 {
		if answer, ok := rable[a]; ok {
			return answer
		}
	}

	if b != 1 || c != 1 {
		return rad(a, 1, 1) * rad(b, 1, 1) * rad(c, 1, 1)
	}

	r = 1
	for prime := range sig(a) {
		r *= prime
	}
	for prime := range sig(b) {
		r *= prime
	}
	for prime := range sig(c) {
		r *= prime
	}

	rable[a] = r
	return
}

func main() {
	starttime := time.Now()

	table = make(map[int64]map[int64]bool)
	populate()

	rable = make(map[int64]int64)

	count, sum := 0, int64(0)

	for c := int64(2); c < max; c++ {

		if rad(c, 1, 1) >= c/2 { //rad(abc) > c for all a,b
			continue
		}

		for a := int64(1); a < c/2; a++ {
			b := c - a
			if coprime(a, b) && rad(a, b, c) < c {
				fmt.Println(a, "\t", b, "\t", c, "\t\t", count, "\t", sum)
				count++
				sum += int64(c)

			}
		}

	}

	fmt.Println(count, sum)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
