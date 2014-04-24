package main

import (
	"euler"
	"fmt"
	"time"
)

const (
	primepi = 1 << (power / 2)
	power   = 50
	top     = 1 << power
)

func next(align []int) []int {
	for i := 0; i < len(align); i++ {
		if eval(fiddle(align, i)) < top {
			return fiddle(align, i)
		}
	}

	return nil
}

func eval(state []int) int64 {
	ans := int64(1)
	for _, pindex := range state {
		p := int64(pindex)
		ans *= int64(euler.Prime(p) * euler.Prime(p))
	}
	return ans
}

func fiddle(state []int, index int) []int {
	size := len(state)
	location := size - index - 1

	faddle := make([]int, size)
	copy(faddle[:location], state[:location])

	faddle[location] = state[location] + 1

	for i := location + 1; i < size; i++ {
		faddle[i] = faddle[i-1] + 1
	}

	return faddle

}

func rng(n int) []int {
	ans := make([]int, n)

	for i := 1; i <= n; i++ {
		ans[i-1] = i
	}
	return ans
}

func sign(n int) int64 {
	if n%2 == 0 {
		return -1
	}
	return 1
}

func main() {
	starttime := time.Now()

	euler.PrimeCache(primepi)
	fmt.Printf("Finished Sieving\t(%s)\n", time.Since(starttime))

	i := 1
	for ; eval(rng(i)) < top; i++ {
	}
	height := i - 1

	total := int64(0)

	for long := 1; long <= height; long++ {
		particle := rng(long)

		for particle != nil {
			delta := top / eval(particle)
			total += sign(long) * delta
			particle = next(particle)
		}

		fmt.Printf("%d of %d\t\t\t\t(%s)\n", long, height, time.Since(starttime))
	}

	fmt.Printf("%d\n", top-total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
