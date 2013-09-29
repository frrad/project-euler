package main

import (
	"fmt"
	"math/rand"
	"time"
)

const dim = 4

func coefficients(epsilon [dim]float64) [dim]float64 {
	return [dim]float64{24 + 24*epsilon[0] + 12*epsilon[1] + 12*epsilon[0]*epsilon[1] + 8*epsilon[2] + 8*epsilon[0]*epsilon[2] + 4*epsilon[1]*epsilon[2] + 4*epsilon[0]*epsilon[1]*epsilon[2] + 6*epsilon[3] + 6*epsilon[0]*epsilon[3] + 3*epsilon[1]*epsilon[3] + 3*epsilon[0]*epsilon[1]*epsilon[3] + 2*epsilon[2]*epsilon[3] + 2*epsilon[0]*epsilon[2]*epsilon[3] + epsilon[1]*epsilon[2]*epsilon[3] + epsilon[0]*epsilon[1]*epsilon[2]*epsilon[3], -50 - 26*epsilon[0] - 19*epsilon[1] - 7*epsilon[0]*epsilon[1] - 14*epsilon[2] - 6*epsilon[0]*epsilon[2] - 5*epsilon[1]*epsilon[2] - epsilon[0]*epsilon[1]*epsilon[2] - 11*epsilon[3] - 5*epsilon[0]*epsilon[3] - 4*epsilon[1]*epsilon[3] - epsilon[0]*epsilon[1]*epsilon[3] - 3*epsilon[2]*epsilon[3] - epsilon[0]*epsilon[2]*epsilon[3] - epsilon[1]*epsilon[2]*epsilon[3], 35 + 9*epsilon[0] + 8*epsilon[1] + epsilon[0]*epsilon[1] + 7*epsilon[2] + epsilon[0]*epsilon[2] + epsilon[1]*epsilon[2] + 6*epsilon[3] + epsilon[0]*epsilon[3] + epsilon[1]*epsilon[3] + epsilon[2]*epsilon[3], -10 - epsilon[0] - epsilon[1] - epsilon[2] - epsilon[3]}
}

func dist(a, b [dim]float64) (dist float64) {

	for i := 0; i < dim; i++ {
		dist += (a[i] - b[i]) * (a[i] - b[i])
	}
	return
}

func add(a, b [dim]float64) (ans [dim]float64) {
	for i := 0; i < dim; i++ {
		ans[i] = a[i] + b[i]
	}
	return
}

func randVec(delta float64) (vec [dim]float64) {
	for i := 0; i < dim; i++ {
		vec[i] = (rand.Float64() * delta) - .5*delta
	}
	return
}

func crazy(vec [dim]float64, thresh float64) bool {

	for i := 0; i < dim; i++ {
		if vec[i] < -1*thresh || vec[i] > 1+thresh {
			return true
		}
	}

	return false

}

func sanitize(vec [dim]float64) [dim]float64 {

	for i := 0; i < dim; i++ {
		if vec[i] < 0 {
			vec[i] = 0
		}
		if vec[i] > 1 {
			vec[i] = 1 - .00001
		}
	}

	return vec

}

func main() {
	starttime := time.Now()

	delta := .001      //Constrols jump speed
	threshold := .0001 //convergence accuracy
	tired := 100000

	targets := make(map[[dim]float64]bool)

	grid := .1
	for a := 0.; a < 1; a += grid {
		for b := 0.; b < 1; b += grid {
			for c := 0.; c < 1; c += grid {
				for d := 0.; d < 1; d += grid {

					co := coefficients([dim]float64{a, b, c, d})

					for i := 0; i < dim; i++ {
						co[i] = float64(int(co[i]))
					}

					targets[co] = true

				}
			}
		}
	}

	potato := 0

	for target, _ := range targets {
		fmt.Print("\r", potato, len(targets))
		potato++

		current := [dim]float64{.5, .5, .5, .5}

		count := 0

		for dist(target, coefficients(current)) > threshold && count < tired {
			count++
			push := randVec(delta)
			nu := add(current, push)
			nu = sanitize(nu)

			myScore := dist(target, coefficients(current))
			newScore := dist(target, coefficients(nu))

			if newScore < myScore {
				current = nu

				// fmt.Println(current)
				// fmt.Println(coefficients(current))
				// fmt.Println("=============")
			}

		}

		if tired == count {
			// fmt.Println("tired")
		} else {
			fmt.Println("===========")
			fmt.Println(current)
			fmt.Println(coefficients(current))
			fmt.Println("TARGET=", target)
		}

	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
