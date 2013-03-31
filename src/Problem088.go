package main

import (
	//"euler"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	max := 10

	for i := 1; i < max; i++ {
		for j := 1; j < max; j++ {
			for k := 1; k < max; k++ {
				for l := 1; l < max; l++ {
					if i+j+k+l == i*j*k*l {
						fmt.Println(i, j, k, l)
					}

				}
			}
		}
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
