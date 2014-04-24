package main

import (
	"euler"
	"fmt"
	"time"
)

// f[k_] := (-1 + 2^k - k) Binomial[26, k];
// Max[Table[f[n], {n, 1, 26}]]

func ways(k int64) int64 {
	return (euler.Exp2(int(k)) - 1 - k) * euler.Choose(26, k)
}

func main() {
	starttime := time.Now()

	record := int64(-1)

	for i := int64(1); i < 26; i++ {
		if ways(i) > record {
			record = ways(i)
		} else {
			break
		}
	}

	fmt.Println(record)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
