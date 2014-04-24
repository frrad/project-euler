package main

import (
	"fmt"
	"time"
)

const (
	mod = 1000000033
	N   = 100
)

var chooseMem, memo map[[2]int]int64
var twoMem map[int]int64

//This is the T of http://oeis.org/A227322
func R(n, m int) (ans int64) {
	if n < m {
		return R(m, n)
	}

	if n == 0 {
		return 1
	}

	if ans, ok := memo[[2]int{n, m}]; ok {
		return ans
	}

	if !(n == 1 && m == 0) {
		ans += inter(1, 0, n, m)
		ans %= mod
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {

			if !(i == n && j == m) {
				ans += inter(i, j, n, m)
				ans %= mod

			}

		}
	}

	big := twoPow(n * m)
	ans = big - ans

	for ans < 0 {
		ans += mod
	}

	memo[[2]int{n, m}] = ans
	return R(n, m)

}

func twoPow(an int) int64 {
	if an == 0 {
		return 1
	}

	if an == 1 {
		return 2
	}

	if ans, ok := twoMem[an]; ok {
		return ans
	}

	half := twoPow(an / 2)
	whole := (half * half) % mod

	if an%2 == 0 {
		twoMem[an] = whole
		return twoPow(an)
	}

	twoMem[an] = (2 * whole) % mod
	return twoPow(an)

}

func choose(n, k int) int64 {
	if k == 0 || n == k {
		return 1
	}

	if ans, ok := chooseMem[[2]int{n, k}]; ok {
		return ans
	}

	ans := choose(n-1, k) + choose(n-1, k-1)
	ans %= mod
	chooseMem[[2]int{n, k}] = ans

	return choose(n, k)
}

func inter(i, j, n, m int) int64 {
	ans := R(i, j)
	ans *= choose(n-1, i-1)
	ans %= mod
	ans *= choose(m, j)
	ans %= mod
	ans *= twoPow((n - i) * (m - j))
	ans %= mod

	return ans
}

func main() {
	starttime := time.Now()

	memo, chooseMem = make(map[[2]int]int64), make(map[[2]int]int64)
	twoMem = make(map[int]int64)

	S := int64(0)

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			S += R(i, j)
			S %= mod
		}
	}

	fmt.Println(S)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
