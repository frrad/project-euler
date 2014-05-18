package main

import (
	"fmt"
	"math/big"
	"time"
)

const top1 = 5678027
const top2 = 7208785

func rowHead(n int) int {
	n--
	return n*(n+1)/2 + 1
}

var remember = make(map[int]bool)
var bigPrime = big.NewInt(0)

func prime(n int) bool {
	if n%2 == 0 || n%3 == 0 || n%5 == 0 {
		return false
	}

	if ans, ok := remember[n]; ok {
		return ans
	}

	bigPrime.SetInt64(int64(n))
	remember[n] = bigPrime.ProbablyPrime(2)

	return prime(n)
}

func rowPrimeTotal(row int) (map[int]bool, map[int]bool, map[int]bool) {
	top, mid, bot := make(map[int]bool), make(map[int]bool), make(map[int]bool)

	A, B, C := rowHead(row-1), rowHead(row), rowHead(row+1)

	consMat := [3][3]int{{0, A, A + 1}, {0, B, B + 1}, {0, C, C + 1}}

	for i := 0; i < row; i++ {
		if i%100000 == 0 {
			fmt.Printf("\r%7d / %7d", i, row)
		}

		//do stuff

		if prime(consMat[1][1]) {

			primeCount := 0
			for i := 0; i < 3; i++ {
				if prime(consMat[0][i]) {
					primeCount++
				}
				if prime(consMat[2][i]) {
					primeCount++
				}
			}

			if primeCount >= 2 {
				for i := 0; i < 3; i++ {
					if topVar := consMat[0][i]; prime(topVar) {
						top[topVar] = true
					}
					if midVar := consMat[1][i]; prime(midVar) {
						mid[midVar] = true
					}
					if botVar := consMat[2][i]; prime(botVar) {
						bot[botVar] = true
					}

				}

				//fmt.Println(consMat)
			}

		}

		for line := 0; line < 3; line++ {
			copy(consMat[line][:2], consMat[line][1:])
			if tad := 1 + consMat[line][1]; tad > 1 && i-line+3 < row {
				consMat[line][2] = tad
			} else {
				consMat[line][2] = 0
			}
		}

	}

	return top, mid, bot
}

func S(row int) *big.Int {
	total := big.NewInt(0)

	fmt.Println("Pass 1/3")
	_, _, a := rowPrimeTotal(row - 1)

	fmt.Println("\nPass 2/3")
	_, b, _ := rowPrimeTotal(row)
	fmt.Println("\nPass 3/3")
	c, _, _ := rowPrimeTotal(row + 1)
	fmt.Print("\n")

	//fmt.Println(a, b, c)

	temp := big.NewInt(0)

	for i := rowHead(row); i < rowHead(row+1); i++ {
		if a[i] || b[i] || c[i] {
			temp.SetInt64(int64(i))
			total.Add(total, temp)
		}
	}

	return total
}

func main() {
	starttime := time.Now()

	ans1 := S(top1)
	fmt.Println(ans1)
	remember = make(map[int]bool)
	ans2 := S(top2)
	fmt.Println(ans2)
	fmt.Println(ans1.Add(ans1, ans2))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
