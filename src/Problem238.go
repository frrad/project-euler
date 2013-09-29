package main

import (
	"fmt"
	"math"

	"time"
)

const (
	dim      = 7
	accuracy = 20
)

func poly(co [dim]float64, x float64) float64 {
	ans := co[0]
	for i := 1; i < dim; i++ {
		term := co[i]
		for j := 0; j < i; j++ {
			term *= x
		}
		ans += term
	}

	monic := x
	for i := 0; i < dim-1; i++ {
		monic *= x
	}

	ans += monic

	return ans
}

func polyD(co [dim]float64, x float64) (ans float64) {
	for i := 1; i < dim; i++ {
		term := co[i] * float64(i)
		for j := 0; j < i-1; j++ {
			term *= x
		}
		ans += term
	}

	monic := float64(dim)

	for i := 0; i < dim-1; i++ {
		monic *= x
	}
	ans += monic

	return ans
}

func NR(quality int, start float64, f, Df func(float64) float64) (x float64) {
	x = start
	for i := 0; i < quality; i++ {
		x = x - (f(x) / Df(x))
	}
	return
}

func quo(f, g func(float64) float64) func(float64) float64 {
	var ans = func(x float64) float64 {

		if g(x) == 0 {
			fmt.Print("asdfasdfasd")
		}

		return f(x) / g(x)

	}
	return ans
}

func quoD(f, fD, g, gD func(float64) float64) func(float64) float64 {
	var ans = func(x float64) float64 {
		if g(x) == 0 {
			fmt.Print("asdfasdfasd")
		}

		return ((g(x) * fD(x)) - (f(x) * gD(x))) / (g(x) * g(x))
	}
	return ans
}

func rt(a float64) func(float64) float64 {
	var ans = func(x float64) float64 {
		return x - a
	}
	return ans
}

func roots(co [dim]float64) [dim]float64 {
	thresher := .001

	var f = func(x float64) float64 {
		return poly(co, x)
	}

	var Df = func(x float64) float64 {
		return polyD(co, x)
	}

	var one = func(x float64) float64 {
		return 1
	}

	plc := 0
	ans := [dim]float64{}

	root := NR(accuracy, .01, f, Df)
	if math.Abs(f(root)) > thresher {
		return ans
	}
	ans[plc] = root
	plc++

	for i := 0; i < dim-1; i++ {

		f, Df = quo(f, rt(root)), quoD(f, Df, rt(root), one)

		root = NR(accuracy, .001, f, Df)
		if math.Abs(f(root)) > thresher {
			return ans
		}
		ans[plc] = root
		plc++

		// fmt.Println(root)
	}

	return ans
}

func shew(a, b float64, f func(float64) float64) {
	fmt.Print("{")
	for i := a; i < b; i += .1 {
		fmt.Print(f(i), ",")
	}
	fmt.Print("}\n\n")
}

func check(co [dim]float64) bool {
	rts := roots(co)

	thresh := .001
	for i := 0; i < dim; i++ {
		if rd := round(rts[i]); math.Abs(rd-rts[i]) < thresh {
			if poly(co, rd) == 0 {
				rts[i] = rd
			}
		}
	}

	// fmt.Println(rts)

	var used [dim]bool

	for i := 0; i < dim; i++ {
		place := int(math.Floor(rts[i]))
		if place < 1 || place > dim {
			return false
		}
		if !used[place-1] {
			used[place-1] = true
		} else {
			return false
		}
	}

	return true
}

func round(x float64) float64 {
	rd := float64(int(x))
	if x-rd > .5 {
		rd++
	}
	return rd
}

func show(a [dim]float64) {
	fmt.Print("{")
	for i := 0; i < dim; i++ {
		fmt.Print(int(a[i]), ",")
	}
	fmt.Print("},")

}

func main() {
	starttime := time.Now()

	sum := 0

	for a := -40320.; a <= -5040; a++ {
		fmt.Println("a", a)
		for b := 13068.; b <= 69264; b++ {
			fmt.Println("b", b)
			for c := -48860.; c <= -13132; c++ {
				for d := 6769.; d <= 18424; d++ {
					for e := -4025.; e <= -1960; e++ {
						for f := 322.; f <= 511; f++ {

							for g := -35.; g < -28; g++ {

								testco := [dim]float64{a, b, c, d, e, f, g}
								if check(testco) {

									show(testco)
									for i := 0; i < dim; i++ {
										sum += int(math.Abs(testco[i]))
									}

									fmt.Println(sum)
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(sum)

	// testco := [dim]float64{26, -52, 48, -12}
	// fmt.Println(check(testco))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
