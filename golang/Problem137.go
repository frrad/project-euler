package main

import (
	"euler"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	i := 15
	fmt.Println(euler.Fibonacci(2*i) * euler.Fibonacci(2*i+1))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
