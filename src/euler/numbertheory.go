package euler

import "math"

const (
	primeTableLength = 100000000
	//lastPrime = Prime[primeTableLength - 1]
	lastPrime          = 2038074739
	totientTableLength = 10000000
)

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64
var totientTable [totientTableLength]int64

//number of primes less than or equal to n
func PrimePi(n int64) int64 {
	if n < 2 {
		return 0
	}

	if n == 2 {
		return 1
	}

	if n < primeTableLength && primepilist[n] != 0 {
		return primepilist[n]
	}

	var answer int64
	if IsPrime(n) {
		answer = 1 + PrimePi(n-1)
		if answer < primeTableLength {
			primeTable[answer] = n
		}
	} else {
		answer = PrimePi(n - 1)

	}

	if n < primeTableLength {
		primepilist[n] = answer
	}
	return answer
}

//Finds/caches all primes below n using a seive
func PrimeCache(n int64) {
	composite := make([]bool, n)
	composite[0], composite[1] = true, true

	p := int64(2)

	for p < n {
		marker := p + p
		for ; marker < n; marker += p {
			composite[marker] = true
		}
		p++
		for ; p < n && composite[p]; p++ {
		}
	}

	m := 1
	for i := int64(0); i < n; i++ {
		if !composite[i] {
			primeTable[m] = i
			primepilist[i] = int64(m)
			m++
		}
	}
}

var factorialtable = make(map[int64]int64)

func Factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	if answer, ok := factorialtable[n]; ok {
		return answer
	}

	answer := Factorial(n-1) * n

	factorialtable[n] = answer

	return answer
}

func IntSqrt(n int64) (sqrt int64, square bool) {
	sqrt = int64(math.Sqrt(float64(n)))
	square = sqrt*sqrt == n
	return
}
