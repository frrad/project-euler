package main

import (
	"euler"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func runs(s string) []int {
	answer := make([]int, 10)

	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for i, digit := range digits {

		answer[i] = strings.Count(s, digit)

	}

	return answer
}

func main() {
	starttime := time.Now()

	euler.PrimeCache(1000000)

	d := int64(10)

	start := euler.IntExp(10, d-1)
	end := euler.IntExp(10, d)

	m, n, s := make([]int, 10), make([]int, 10), make([]int64, 10)

	for prime := start; prime < end; prime++ {

		if prime%1000000 == 0 {
			fmt.Println((100 * float64(prime-start) / float64(end-start)), "\b%")
		}

		if euler.IsPrime(prime) {
			this := runs(strconv.FormatInt(prime, 10))

			for index, value := range this {
				if value == m[index] {
					n[index]++
					s[index] += prime
				}
				if value > m[index] {
					m[index] = value
					n[index] = 1
					s[index] = prime
				}
			}
		}
	}

	total := int64(0)
	for _, summand := range s {
		total += summand
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
