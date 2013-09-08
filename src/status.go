package main

import (
	"euler"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	fmt.Println("Hello, World", euler.Prime(10000))

	page := euler.Import("../problemdata/Project Euler.html")

	fmt.Println(page)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
