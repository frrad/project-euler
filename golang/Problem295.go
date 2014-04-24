package main

import (
	"euler"
	"fmt"
	"time"
)

const (
	top     = 450
	N       = 100000
	epsilon = .2
)

var dupes map[int64][][2]int

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
func enumerate(a, b int, rMin float64) (radii []float64) {
	radii = make([]float64, 0)

	A, B := &euler.Point{float64(a), 0}, &euler.Point{0, float64(b)}

	base, _ := euler.LineFromPoints(A, B)

	x, y := int64(0), int64(0)

	if a != b {

		lhs := b*b - a*a

		lhs /= 2 //divisible by construction

		gcd := int(euler.GCD(int64(a), int64(b)))

		if lhs%gcd != 0 {
			return []float64{}
		}

		mult := lhs / gcd

		x, y = euler.ExtendedEuclidean(int64(a), int64(b))
		x *= -1

		x *= int64(mult)
		y *= int64(mult)
	}

	testpt := &euler.Point{float64(x), float64(y)}

	//VITAL: Must start at the beginning
	for testpt.RightOf(base) {
		testpt.X -= float64(b)
		testpt.Y -= float64(a)
	}

	for !testpt.RightOf(base) {
		testpt.X += float64(b)
		testpt.Y += float64(a)
	}

	trip2 := int64(-1)

	for trip2 <= 0 {

		rad := testpt.Distance(A)

		trip1 := rMin*rMin - float64(rdSqr(rad))
		trip2 = rdSqr(rad) - N*N

		if trip1 < 0 && trip2 <= 0 {
			radii = append(radii, rad)
		}

		testpt.X += float64(b)
		testpt.Y += float64(a)

	}

	return

}

func rdSqr(x float64) int64 {
	sq := x * x
	ans := int64(sq)

	if sq-float64(ans) > epsilon {
		ans++
	}

	return ans

}

//we assume list sorted
func count(list []float64, a, b int) (ct int64) {
	for i := 0; i < len(list); i++ {
		for j := i; j < len(list); j++ {
			if check(list[i], list[j]) {
				ct++

			}

		}
	}

	for i := 0; i < len(list); i++ {
		dupes[rdSqr(list[i])] = append(dupes[rdSqr(list[i])], [2]int{a, b})

	}
	return
}

func check(a, b float64) bool {

	for _, tuple := range dupes[rdSqr(a)] {
		for _, duple := range dupes[rdSqr(b)] {
			if tuple == duple {
				return false
			}
		}
	}

	return true
}

func main() {
	starttime := time.Now()

	dupes = make(map[int64][][2]int)

	total := int64(0)

	for a := 1; a < top; a++ {

		for b := a; b < top; b++ {
			gcd := int(euler.GCD(int64(a), int64(b)))
			lhs := b*b - a*a

			if gcd != 1 || lhs%2 != 0 || (lhs/2)%gcd != 0 {
				continue
			}

			rMin := rLowerBound(float64(a), float64(b))

			if rMin <= float64(N) {
				fmt.Println("A=", a, "B=", b)
				fmt.Printf("Radius between %.2f and %d\n", rMin, N)
				en := enumerate(a, b, rMin)
				fmt.Println("Found:", len(en))
				total += count(en, a, b)
				fmt.Println("Total:", total)
				fmt.Println("++++++++")

			}

		}
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
