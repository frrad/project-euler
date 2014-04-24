package main

import (
	"euler"
	"fmt"
	"strconv"
	"time"
)

const (
	digits  = 12
	modulus = 100000
	max     = 1000000000000
)

//compute a^b mod mod
func powMod(a, b, mod int64) int64 {
	if b == 1 {
		//fmt.Println(a, "^", b, "mod", mod, "=", a%mod)
		return a % mod
	}
	if b%2 == 0 {
		half := powMod(a, b/2, mod)
		//fmt.Println(a, "^", b, "mod", mod, "=", (half*half)%mod)
		return (half * half) % mod
	}

	half := powMod(a, (b-1)/2, mod)
	//fmt.Println(a, "^", b, "mod", mod, "=", (a*half*half)%mod)

	return (a * half * half) % mod

}

func tens(n int) int {

	ans := 1
	for i := 1; i <= n; i++ {
		ans *= 10

	}
	return ans
}

//how many numbers <= max have signature sig
//signatures have 5 digits
func multiplicity(sig, max int64) int64 {
	lgth := euler.NumberDigits(sig)
	mlgth := euler.NumberDigits(max)

	if lgth > mlgth {
		return 0
	}

	if lgth == mlgth && sig <= max {
		return 1
	}
	if lgth == mlgth && sig > max {
		return 0
	}

	if mlgth <= 5 {
		return 1 + multiplicity(sig, max/10)

	}

	answer := multiplicity(sig, max/10)

	max -= sig

	for i := 0; i < 5; i++ {
		max /= 10
	}

	answer += max
	answer++

	return answer
}

func mix(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func shorten(n int64) int64 {
	s := strconv.FormatInt(n, 10)
	s = s[1:]
	ans, _ := strconv.ParseInt(s, 10, 64)
	return ans
}

func main() {
	starttime := time.Now()

	eff := int64(1)
	twoCount := int64(0)

	for preBase := int64(1); preBase < modulus; preBase++ {
		if preBase%2 == 0 || preBase%5 == 0 {
			continue
		}
		mult := multiplicity(preBase, max)

		eff *= powMod(preBase, mult, modulus)
		eff %= modulus

		twoMult := 1
		for base := int64(preBase) * 2; base < max; base *= 2 {

			mult := multiplicity(preBase, max/powMod(2, int64(twoMult), max))

			eff *= powMod(preBase, mult, modulus)
			//fmt.Println(twoMult)
			eff %= modulus

			twoCount += mult * int64(twoMult)

			twoMult++
		}

		//fmt.Println(twoCount, eff)

		fiveMult := 1
		for base := int64(preBase) * 5; base < max; base *= 5 {

			mult := multiplicity(preBase, max/powMod(5, int64(fiveMult), max))

			eff *= powMod(preBase, mult, modulus)
			eff %= modulus

			twoCount -= mult * int64(fiveMult)
			fiveMult++
		}

		//fmt.Println(twoCount, eff)

	}

	fmt.Println((powMod(2, twoCount, modulus) * eff) % modulus)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
