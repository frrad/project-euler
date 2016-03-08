package main

import (
	"euler/plane"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	BC, AC, AB := 6., 4., 5.

	x := plane.Point{0, 0}
	y := plane.Point{BC, 0}

	c1, c2 := plane.Circle{x, AC}, plane.Circle{y, AB}


	fmt.Println(c1.IntersectionCircle(c2))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
