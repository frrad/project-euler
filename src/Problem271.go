package main

import (
	"euler"
	"fmt"
	"time"
)

//http://en.wikipedia.org/wiki/Chinese_remainder_theorem#A_constructive_algorithm_to_find_the_solution
func main() {
	starttime := time.Now()

	fmt.Println("Hello, World", euler.Prime(10000))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
