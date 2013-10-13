package main

import (
	"fmt"
	"time"
)

func row(n int) int64 {
	if n < 7 {
		return int64(n + 1)
	}

	ans := (1 + int64(n%7)) * row(n/7)

	return ans
}

//Pattern: Look at this picture
//Table[Table[If[Mod[Binomial[n, k], 7] == 0, " ", "X"], {k, 0, n}], {n, 0, 97}] // MatrixForm
func main() {
	starttime := time.Now()

	sum := int64(0)

	for i := 0; i < 1000000000; i++ {
		sum += row(i)
	}

	fmt.Println(sum)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
