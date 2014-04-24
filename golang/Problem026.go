package main

import (
	"fmt"
	"time"
)

const level = 10

func chomp(numerator int, denominator int, height int) int {

	for i := 1; i <= height; i++ {

		numerator *= 10
		numerator = numerator % denominator

	}

	return numerator
}

func main() {
	starttime := time.Now()

	record, submit := 0, 0

	for den := 2; den < 1000; den++ {

		newmerator := chomp(1, den, level)

		answer := 0
		for j := 1; newmerator != chomp(newmerator, den, j); j++ { //this is pretty wasteful
			answer = j + 1
		}
		if answer > record {
			record, submit = answer, den
			fmt.Println(den, ",", answer)
		}

	}

	fmt.Println(submit)
	fmt.Println("Elapsed time:", time.Since(starttime))
}
