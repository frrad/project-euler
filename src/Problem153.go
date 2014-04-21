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

var cheat = make(map[int]int)

func line(z *gauss) (total int) {
	mag := z.squareNorm()

	if ans, ok := cheat[top/mag]; ok {
		total = ans
	} else {
		for a := 1; mag*a <= top; a++ {
			for b := a; b*a*mag <= top; b++ {
				if b == a {
					total += a
				} else {
					total += (a + b)
				}
			}
		}
		cheat[top/mag] = total
	}

	total *= z.r

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
			delta := line(&z)
			ans += delta
		}
		//fmt.Printf("(%d) total:%d\t\t %d+%di\n", len(cheat), ans, z.r, z.i)
		z.i = 0
	}

	fmt.Println(ans)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
