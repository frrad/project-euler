package main

import (
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	//brute force is good enough, no need to resort to generating function
	//maybe implement in euler lib later?

	top := 1000

	for a := 1; a <= top; a++ {
		for b := a; b+a <= top && top-a-b >= b; b++ {

			c := top - b - a

			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
			}
		}
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
