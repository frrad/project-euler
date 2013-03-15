package main

import "fmt"

var current3 int = 0
var current5 int = 0
var current6 int = 0

func triangle() int {

	current3++
	return current3 * (current3 + 1) / 2
}

func pentagon() int {

	current5++
	return current5 * (3*current5 - 1) / 2
}

func hexagon() int {

	current6++
	return current6 * (2*current6 - 1)
}

func main() {

	pent :=0
	hex := 0
	tri :=0

	for {
		for pent < tri || pent < hex{
			pent = pentagon()
		}
		for tri < pent || tri < hex{
			tri = triangle()
		}
		for hex < pent || hex < tri{
			hex = hexagon()
		}

		if tri == pent && tri == hex {
			fmt.Println(tri)
			tri = triangle()
		}
	}
}

