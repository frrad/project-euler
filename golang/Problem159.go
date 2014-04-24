package main

import (
	"fmt"
	"math"
	"time"
)

const top = 1000000

var mdmemo map[int]int

func dr(an int) (sum int) {
	for an > 0 {
		sum += an % 10
		an /= 10
	}

	if sum > 9 {
		return dr(sum)
	}

	return sum
}

func mdrs(an int) (best int) {
	if ans, ok := mdmemo[an]; ok {
		return ans
	}

	for i := 2; i <= int(math.Sqrt(float64(an))); i++ {
		if test := mdrs(i) + mdrs(an/i); an%i == 0 && test > best {
			best = test
		}
	}
	if dr(an) > best {
		best = dr(an)
	}

	mdmemo[an] = best
	return
}

func main() {
	starttime := time.Now()

	mdmemo = make(map[int]int)

	sum := 0

	for i := 2; i < top; i++ {

		// if mdrs(i) == dr(i) {
		// 	fmt.Println(i)
		// }

		sum += mdrs(i)
	}

	fmt.Println(sum)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
