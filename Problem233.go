package main

import (
	"./euler"
	"fmt"
	"math"
	"strconv"
	"time"
)

const threshold = .0000000001

func solveBottom(x float64, n float64) float64 {
	return (1.0 / 2.0) * (n - math.Sqrt((n*n)+(4*n*x)-(4*x*x)))
}

func naive(n int64) int64 {
	total := int64(0)
	N := float64(n)
	for x := math.Ceil(N / 2); x < N/2+(N*math.Sqrt(2)/2); x++ {
		y := solveBottom(x, N)
		yround := float64(int(y))
		if math.Abs(y-yround) < threshold {
			total++
		}
	}
	return total * 4
}

//number of pythagorean triples with "C"=n
func better(n int64) int64 {

	count := int64(0)
	for i := int64(0); i < n; i++ {
		a2 := n*n - i*i
		if euler.IsSquare(a2) {
			count++
		}
	}
	return 4 * count

}

func isSmall(str string) bool {
	return len(str) < 11
}

func isMultiple(n int64) bool { //returns true if there are no pythagorean prime factors
	for _, factor := range euler.Factor(n) {
		if factor%4 == 1 {
			return false
		}
	}
	return true
}

func lastSmall(table []string, multiple string) int {
	min := 0
	max := len(table)
	for max-min > 1 {

		if isSmall(euler.StringProd(multiple, table[(min+max)/2])) {
			min = (min + max) / 2
		} else {
			max = (min + max) / 2
		}

	}
	return min
}

func main() {

	starttime := time.Now()

	//We're looking for numbers of the form 2^k * \prod pi * q1 * q2^2 * q3^3
	//for pi,qi prime and pi % 4 = 3 and qi % 4 = 1. Start by populating table
	primes1 := make([]string, 0) //Table of pythagorean primes

	for i := int64(1); i < 350000; i++ {
		num := euler.Prime(i)
		prime := strconv.FormatInt(num, 10)

		if num%4 == 1 {
			primes1 = append(primes1, prime)

		}
	}

	multitable := make([]string, 1) //numbers of the form 2^k \prod pi
	multitable[0] = "1"

	for i := int64(1); i < 30000; i++ {

		if isMultiple(i) {
			multitable = append(multitable, strconv.FormatInt(i, 10))
		}
	}

	sumtable := make([]string, len(multitable)) //partial sums of multitable
	sumtable[0] = multitable[0]
	for i := 1; i < len(multitable); i++ {
		sumtable[i] = euler.StringSum(sumtable[i-1], multitable[i])
	}

	fmt.Println("Tables written")

	total := "0"
	pentagon := "0"
	counttttt, dice := 0, 0
	for two := 0; isSmall(pentagon); two++ { //Index of q2

		pentagon = euler.StringExp(primes1[two], 10)

		hex := pentagon
		for one := 0; isSmall(hex); one++ { //Index of q3
			if one == two {
				one++
			}

			hex = euler.StringProd(pentagon, euler.StringExp(primes1[one], 2))

			if isSmall(hex) {
				dice := lastSmall(multitable, hex)
				contribution := euler.StringProd(hex, sumtable[dice])
				counttttt += (dice + 1)
				total = euler.StringSum(total, contribution)

			}

		}
	}
	fmt.Println(counttttt)

	fmt.Println(total)

	pentagon = "0"
	for two := 0; isSmall(pentagon); two++ { //Index of q2

		pentagon = euler.StringExp(primes1[two], 7)

		hex := pentagon
		for one := 0; isSmall(hex); one++ { //Index of q3
			if one == two {
				one++
			}

			hex = euler.StringProd(pentagon, euler.StringExp(primes1[one], 3))

			if isSmall(hex) {
				dice = lastSmall(multitable, hex)
				contribution := euler.StringProd(hex, sumtable[dice])
				total = euler.StringSum(total, contribution)
				counttttt += (dice + 1)

			}

		}
	}

	fmt.Println(total)

	cube := "0"

	for three := 0; isSmall(cube); three++ { //Index of q3
		cube = euler.StringExp(primes1[three], 3)

		pentagon := cube
		for two := 0; isSmall(pentagon); two++ { //Index of q2
			if two == three {
				two++
			}
			pentagon = euler.StringProd(cube, euler.StringExp(primes1[two], 2))

			hex := pentagon
			for one := 0; isSmall(hex); one++ { //Index of q3
				for three == one || one == two {
					one++
				}
				hex = euler.StringProd(pentagon, primes1[one])
				if isSmall(hex) {
					dice = lastSmall(multitable, hex)
					contribution := euler.StringProd(hex, sumtable[dice])
					counttttt += dice + 1
					total = euler.StringSum(total, contribution)

				}

			}
		}

	}

	fmt.Println(total, counttttt)

	fmt.Println("Elapsed time:", time.Since(starttime))

}
