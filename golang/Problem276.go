package main

import (
	"euler"
	"fmt"
	"time"
)

// Number of triangles
func Triangle(n int64) int64 {
	return ((n * n / 6) + 1) / 2
}

// Number of primitive triangles
func Primitive(n int64) int64 {
	if euler.IsPrime(n) {
		return Triangle(n)
	}

	ans := int64(0)
	for i := int64(2); i < n; i++ {
		if n%i == 0 {
			ans += Primitive(i)
		}
	}
	return Triangle(n) - ans

}

func main() {
	starttime := time.Now()

	for i := int64(3); i < 10; i++ {
		fmt.Println(i, Triangle(i), Primitive(i))
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
