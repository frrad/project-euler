package main

import (
	"fmt"
	"time"
)

const top = 15

// Given a folding configuration (trinary number) return a slice of bitmasks to
// compute the score of a binary number in this folding.
func masks(arranged int) {
	var positions [15][2]int

	positions[0] = [2]int{0, 0}
	positions[1] = [2]int{1, 0}

	current := positions[1]
	aim := [2]int{1, 0}

	for i := 0; i < top-2; i++ {
		indicator := arranged % 3
		arranged %= 3

	}

	fmt.Println(positions)

}

func main() {
	starttime := time.Now()

	masks(10)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
