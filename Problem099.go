package main

import (
	"./euler"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	data := euler.Import("problemdata/base_exp.txt")

	for _, line := range data {
		fmt.Println(line)
	}

	fmt.Println("Elapsed time:", time.Since(starttime))

}
