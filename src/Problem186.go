package main

import (
	"fmt"
	"time"
)

var randMemo map[int]int

func S(k int) int {
	if ans, ok := randMemo[k]; ok {
		return ans
	}

	if k <= 55 {
		kay := int64(k)
		temp := (100003 - (200003 * kay) + (300007 * kay * kay * kay)) % 1000000
		if temp < 0 {
			fmt.Println(k, temp)
		}
		randMemo[k] = int(temp)
		return S(k)
	}

	randMemo[k] = (S(k-24) + S(k-55)) % 1000000
	return S(k)
}

func main() {
	starttime := time.Now()
	randMemo = make(map[int]int)

	lowest := make(map[int]int)
	for i := 0; i < 1000000; i++ {
		lowest[i] = i
	}

	all := make(map[int][]int)
	for i := 0; i < 1000000; i++ {
		all[i] = []int{i}
	}

	i := 1
	calls := 0

	for len(all[lowest[524287]]) < 990000 {
		a, b := S(i), S(i+1)
		i += 2

		if a == b {
			continue
		}

		calls++

		if lowest[a] == lowest[b] {
			continue
		}

		if lowest[a] > lowest[b] {
			a, b = b, a
		}

		old := lowest[b]

		for _, numb := range all[old] {
			lowest[numb] = lowest[a]
			all[lowest[a]] = append(all[lowest[a]], numb)
		}

		delete(all, old)

	}

	fmt.Println(calls)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
