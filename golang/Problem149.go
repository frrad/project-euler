package main

import (
	"fmt"
	"time"
)

const dim = 2000

var randMemo map[int]int64

func S(k int) int64 {

	if ans, ok := randMemo[k]; ok {
		return ans
	}

	if k <= 55 {
		kay := int64(k)
		temp := ((100003 - (200003 * kay) + (300007 * kay * kay * kay)) % 1000000) - 500000

		randMemo[k] = (temp)
		return S(k)
	}

	randMemo[k] = ((S(k-24) + S(k-55) + 1000000) % 1000000) - 500000
	return S(k)
}

func tab(i, j int) int64 {
	return S(1 + j + (dim * i))
}

func main() {
	starttime := time.Now()

	randMemo = make(map[int]int64)

	max := int64(0)

	//Diagonal not implemented because lazy
	for i := 0; i < dim; i++ {
		current, corrent := int64(0), int64(0)

		for j := 0; j < dim; j++ {

			current += tab(i, j)

			if current < 0 {
				current = 0
			}

			if current > max {
				max = current
				fmt.Println(max)
			}

			corrent += tab(j, i)

			if corrent < 0 {
				corrent = 0
			}

			if corrent > max {
				max = corrent
				fmt.Println(max)
			}

		}
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
