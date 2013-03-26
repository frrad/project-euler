package main

import (
	"./euler"
	"fmt"
	"time"
)

const (
	lid    = 1500
	target = 100000000
)

func main() {

	starttime := time.Now()

	counter := 0

	for i := int64(1); i < lid; i++ {

		primei := euler.Prime(i)
		primej := primei

		for j := i + 1; primei*primej < target; j++ {

			counter++
			fmt.Println(i, j, primei*primej)

			primej = euler.Prime(j)

		}

	}

	fmt.Println(counter)

	fmt.Println("Elapsed time:", time.Since(starttime))

}
