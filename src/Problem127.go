package main

import (
	"euler"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	searchTop := int64(120000)

	csum := int64(0)
	ccount := 0

	for c := int64(1); c < searchTop; c++ {

		if c%1500 == 0 {
			fmt.Println("C=", c)
		}

		for a := int64(1); a < c/2; a++ {

			b := c - a

			if euler.GCD(a, b) != 1 {
				continue
			}

			rad := int64(1)
			ABC := [3]int64{a, b, c}
			broke := false

			for _, abc := range ABC {

				p := int64(2)

				for abc > 1 && broke == false {
					if abc%p == 0 {
						rad *= p
						for abc%p == 0 {
							abc /= p
						}
					}
					if rad > c {
						broke = true
						break
					}
					p++
				}

				if broke {
					continue
				}
			}

			if broke {
				continue
			}

			ccount++
			csum += c
			fmt.Println(a, b, c)
			fmt.Println(ccount, csum, "\n")

		}

	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
