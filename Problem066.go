package main

import (
	"./euler"
	"fmt"
	"math"
	"time"
)

func y(x, d int) bool {
	X := float64(x)
	D := float64(d)

	y := int(math.Sqrt((1 - (X * X)) / (-D)))

	return (x*x)-(d*y*y) == 1

}

func main() {
	starttime := time.Now()

	for d := 2; d <= 1000; d++ {

		if euler.IsSquare(int64(d)) {
			d++
		}
		X := 1
		for x := 2; !y(x, d); x++ {

			X = x + 1
		}
		fmt.Println(d, X)

	}

	fmt.Println("Elapsed time:", time.Since(starttime))

}
