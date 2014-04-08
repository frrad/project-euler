package main

import (
	"euler"
	"fmt"
	"time"
)

var (
	holes int
	bins  [][2]int
)

//number of ways to assign balls to bins in index
//less than or equal to index
func ways(a, b, limit int) (total int64) {
	//olda, oldb := a, b
	if a == 0 || b == 0 || limit < 0 {
		return int64(euler.IntPartitions(a) * euler.IntPartitions(b))
		//no more bins!
	}

	current := bins[limit]
	thisA, thisB := current[0], current[1]

	most := euler.MinInt(a/thisA, b/thisB)

	for i := 0; i <= most; i++ {
		total += ways(a, b, limit-1) * int64(euler.IntPartitions(i))
		a -= thisA
		b -= thisB
	}

	//fmt.Printf("ways(%d,%d,%d)=%d\n", olda, oldb, limit, total)

	return
}

func main() {
	starttime := time.Now()
	for B := 0; B < 8; B++ {
		for A := 1; A < 8; A++ {
			holes = 0
			bins = make([][2]int, 0)
			for a := 1; a <= A; a++ {
				for b := 1; b <= B; b++ {
					if euler.GCD(int64(a), int64(b)) == 1 {
						bins = append(bins, [2]int{a, b})
						holes++
					}
				}
			}

			fmt.Printf("%d:%d = %d\t", A, B, ways(A, B, holes-1))

		}
		fmt.Printf("\n")

	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
