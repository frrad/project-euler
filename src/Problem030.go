package main

import (
	"fmt"
	"strconv"
	"time"
)

func digits(n int) int {
	str := strconv.Itoa(n)
	sum := 0
	for i := 0; i < len(str); i++ {
		x, _ := strconv.Atoi(str[i : i+1])
		sum += x * x * x * x * x
	}
	return sum
}

func main() {
	starttime := time.Now()

	total := 0

	for i := 2; i < 999999; i++ {
		if digits(i) == i {
			fmt.Println(i)
			total += i
		}
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
