package main

import (
	"fmt"
	"time"
)

func ways(squares int) int {
	if squares == 0 {
		return 1
	}
	if squares < 3 {
		return 1
	}

	total := 1 //The empty configuration

	for size := 3; size < squares; size++ {
		for start := 0; start <= squares-size; start++ {
			answer := 1
			if start > 0 {
				answer *= ways(start - 1)
			}
			if squares-start-size > 0 {
				answer *= ways(squares - start - size - 1)
			}
			total += answer
			fmt.Println(answer, "ways to fill", squares, "after start:", start, "size:", size)
		}

	}

	return total
}

func main() {
	starttime := time.Now()

	fmt.Println(ways(4) + 2)

	fmt.Println("Elapsed time:", time.Since(starttime))

}
