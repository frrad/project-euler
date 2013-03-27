package euler

var primepilist [primeTableLength]int64

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
