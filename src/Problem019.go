package main

import (
	"fmt"
	"time"
)

func month(monthindex, startday, year int) (endday int) {
	lookup := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	endday = startday + lookup[monthindex]
	endday %= 7

	if monthindex == 1 && year%4 == 0 { //leap year
		if year%100 == 0 && !(year%400 == 0) { //or not
			return
		}
		return (endday + 1) % 7
	}

	return
}

func yearS(n int) int {
	if n == 1900 {
		return 1
	}

	ylength := 365
	if month(1, 0, n-1) != 0 {
		ylength++
	}

	return (yearS(n-1) + ylength) % 7
}

func year(n int) int {
	count := 0
	yearstart := yearS(n)
	if yearstart == 0 {
		count++
	}

	for i := 0; i < 11; i++ {
		yearstart = month(i, yearstart, n)
		if yearstart == 0 {
			count++
		}
	}

	return count

}

func main() {
	starttime := time.Now()

	answer := 0
	for i := 1901; i <= 2000; i++ {
		answer += year(i)
	}

	fmt.Println(answer)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
