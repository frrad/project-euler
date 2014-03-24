package main

import (
	"fmt"
	"time"
)

var current3 int64 = 0
var current5 int64 = 0
var current6 int64 = 0

func triangle() int64 {
	current3++
	return current3 * (current3 + 1) / 2
}

func pentagon() int64 {
	current5++
	return current5 * (3*current5 - 1) / 2
}

func hexagon() int64 {
	current6++
	return current6 * (2*current6 - 1)
}

func main() {
	starttime := time.Now()

	var pent int64 = 0
	var hex int64 = 0
	var tri int64 = 0

	for {

		for pent < tri || pent < hex {
			pent = pentagon()
		}
		for tri < pent || tri < hex {
			tri = triangle()
		}
		for hex < pent || hex < tri {
			hex = hexagon()
		}

		if tri == pent && tri == hex {
			fmt.Println(tri, ":", current3, current5, current6)
			if tri > 40755 {
				break
			}
			tri = triangle()

		}

	}

	fmt.Println(tri)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
