package main

import (
	"euler"
	"fmt"
	"time"
)

const top = 100000000

type gauss struct {
	r int
	i int
}

func (z *gauss) squareNorm() int {
	return z.r*z.r + z.i*z.i
}

func (z *gauss) conjugate() *gauss {
	return &gauss{z.r, -z.i}
}

//product, assuming totally real
func (z *gauss) times(w *gauss) int {
	return z.r*w.r - z.i*w.i
}

func (z *gauss) scale(n int) *gauss {
	return &gauss{n * z.r, n * z.i}
}

func line(z *gauss) (total int) {
	w := z.conjugate()

	for a := 1; z.times(w.scale(a)) <= top; a++ {
		for b := a; z.scale(b).times(w.scale(a)) <= top; b++ {
			//fmt.Printf("%d-%d \n", a, b)

			if b == a {
				total += z.r * a
			} else {
				total += z.r * (a + b)
			}
		}
	}

	if z.i > 0 {
		total *= 2
	}
	return
}

func main() {
	starttime := time.Now()

	ans := 0

	z := gauss{1, 0}
	for realPart := 1; z.squareNorm() <= top; realPart++ {
		z.r = realPart
		for imPart := 0; z.squareNorm() <= top; imPart++ {
			z.i = imPart
			if euler.GCD(int64(z.i), int64(z.r)) != 1 {
				continue
			}
			//fmt.Println(z)
			//fmt.Printf("%d\n", z.squareNorm())
			delta := line(&z)
			ans += delta
			fmt.Printf("delta: %d\t total:%d\t\t %d+%di\n", delta, ans, z.r, z.i)

		}
		z.i = 0
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
