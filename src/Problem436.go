package main

import (
	"fmt"
	"math"
	"time"
)

func round(num, digits float64) float64 {

	num *= math.Pow(10, digits)

	temp := float64(int64(num))

	if num-temp > .5 {
		temp++
	}

	num = temp * math.Pow(10, -1*digits)

	return num
}

func main() {
	starttime := time.Now()

	ans := (1. / 4.) + (7. * math.E / 2.) - (5. * math.E * math.E / 4.)

	fmt.Println(round(ans, 10))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
