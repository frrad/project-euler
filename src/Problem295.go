package main

import (
	"euler"
	"fmt"
	"time"
)

const (
	top = 450
	N   = 10
)

//we expect a > b
func rLowerBound(b, a float64) float64 {
	if a == 1 && b == 1 {
		return .5
	}

	A, B := &euler.Point{a, 0}, &euler.Point{0, b}
	ln, _ := euler.LineFromPoints(A, B)

	best := 0.

	for test := float64(1); test < a; test++ {
		testy := float64(int(ln.Evaluate(test)))
		testP := &euler.Point{test, testy}

		circ, _ := euler.CircleFromPoints(A, B, testP)
		if circ.Radius > best {
			best = circ.Radius
		}
	}

	return best
}

//Solve b y - a x ==( b^2 -a^2) / 2 to get candidates, then check radii
//r = sqrt((x-a)^2 + y^2)
func enumerate(a, b int, rMin, rMax float64) int {
	if a == b {
		return 123 //fix
	}

	lhs := b*b - a*a
	if lhs%2 != 0 {
		return 0
	}
	lhs /= 2

	gcd := int(euler.GCD(int64(a), int64(b)))

	if lhs%gcd != 0 {
		return 0
	}

	mult := lhs / gcd

	x, y := euler.ExtendedEuclidean(int64(a), int64(b))
	x *= -1

	x *= int64(mult)
	y *= int64(mult)

	fmt.Println("DEBUG", x, y)
	fmt.Println(a, "*", x, "+", b, "*", y, "=", mult)
	return 0
}

func main() {
	starttime := time.Now()

	ct := 0
	total := 0

	for a := 1; a < top; a++ {
		fmt.Print("\r==================")
		for b := a; b < top; b++ {
			if euler.GCD(int64(a), int64(b)) != 1 {
				continue
			}

			rMin := rLowerBound(float64(a), float64(b))
			rMax := float64(N)

			if rMin <= rMax {
				fmt.Println(a, b)
				fmt.Println(rMin, rMax, rMax-rMin)
				total += enumerate(a, b, rMin, rMax)
				fmt.Println("total ", total)
				ct++
			}

		}
	}

	fmt.Println(ct)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
