package main

import (
	"euler"
	"fmt"
	"math"
	"time"
)

var phi float64 = (1 + math.Sqrt(5)) / 2
var phi2 = phi * phi

const (
	A = "1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679"
	B = "8214808651328230664709384460955058223172535940812848111745028410270193852110555964462294895493038196"
)

var size = int64(len(A))

func sequence(n int) int64 {
	return (127 + 19*int64(n)) * euler.IntExp(7, int64(n))
}

//is smart(k) == n for some k?
func nth(n int64) bool {
	if n == 1 {
		return true
	}
	if n == 2 {
		return false
	}

	//first double until we get an upper bound
	k := int64(1)
	for smart(k) < n {
		k *= 2
	}

	//now binary search
	a, b := int64(1), k
	for b-a > 1 {

		c := (b + a) / 2
		if smart(c) > n {
			b = c
		} else if smart(c) < n {
			a = c
		} else if smart(c) == n {
			return true
		}

	}

	return smart(a) == n || smart(b) == n
}

func smart(i int64) int64 {
	return 2 + int64(float64(i-1)*phi2)
}

func main() {
	starttime := time.Now()

	for i := 17; i >= 0; i-- {
		seq := sequence(i)

		n, k := seq/size+1, seq%size-1

		if k == -1 {
			k = size - 1
			n--
		}

		if nth(n) {
			fmt.Print(string(A[k]))
		} else {
			fmt.Print(string(B[k]))
		}
	}

	fmt.Println("\nElapsed time:", time.Since(starttime))
}
