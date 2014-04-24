package main

import (
	"./euler"
	"fmt"
	"math"
	"math/big"
	"time"
)

func ctdFrac(list []int) (num string, den string) {
	current := big.NewRat(int64(list[len(list)-1]), 1)

	for i := len(list) - 2; i >= 0; i-- {
		contrib := big.NewRat(int64(list[i]), 1)

		current.Inv(current)
		current.Add(current, contrib)
	}

	return current.Num().String(), current.Denom().String()
}

//fraction of the form (a +  b \sqrt R) / d
type frac struct {
	a int
	b int
	R int
	d int
}

func flip(F frac) frac {
	return frac{
		a: F.a * F.d,
		b: -F.b * F.d,
		R: F.R,
		d: (F.a * F.a) - (F.b * F.b * F.R)}
}

func reduce(F frac) frac {
	gcd := int(euler.GCD(int64(F.a), int64(F.b)))
	gcd = int(euler.GCD(int64(gcd), int64(F.d)))
	return frac{F.a / gcd, F.b / gcd, F.R, F.d / gcd}
}

func nextFrac(F frac) (n int, next frac) {
	total := (float64(F.a) + (float64(F.b) * math.Sqrt(float64(F.R)))) / float64(F.d)

	n = int(total)

	if n != 0 {
		next = frac{
			a: F.a - (n * F.d),
			b: F.b,
			R: F.R,
			d: F.d}
	} else {
		n = 1
		next = frac{
			a: F.a + F.d,
			b: F.b,
			R: F.R,
			d: F.d}
	}

	next = flip(next)
	next = reduce(next)

	return
}

//"Pell Equation"
func main() {
	starttime := time.Now()

	total := 0
	for rad := 2; rad <= 100; rad++ {

		if !euler.IsSquare(int64(rad)) {

			convergentList := make([]int, 1)

			test := frac{a: 0, b: 1, R: rad, d: 1}
			n := 0

			n, test = nextFrac(test)

			convergentList[0] = n

			for i := 0; i < 200; i++ {
				n, test = nextFrac(test)
				convergentList = append(convergentList, n)

			}

			fLength := len(convergentList) - 1

			p, q := ctdFrac(convergentList[:fLength])

			var r big.Rat

			r.SetString(p + "/" + q)

			deciString := r.FloatString(105) //A bit too long in order to avoid rounding

			total += euler.StringDigitSum(deciString[:101]) //The extra 1 is the decimal point
		}
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))

}
