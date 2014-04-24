package main

import (
	"fmt"
	"time"
)

func get(imax, ijmax, digits int) int64 {
	if answer, ok := memo[[3]int{imax, ijmax, digits}]; ok {
		return answer
	}
	if digits <= 0 {
		return int64(1)
	}
	if digits == 1 {
		if imax > ijmax {
			return int64(ijmax)
		} else {
			return int64(imax)
		}
	}
	if digits == 2 {
		answer := int64(0)
		for i := 10; i <= imax*10+9; i++ {
			if i/10+i%10 <= ijmax && i/10+i%10 <= 9 {
				answer++
			}
		}
		memo[[3]int{imax, ijmax, digits}] = answer
		return answer
	}

	answer := int64(0)
	for i := 1; i <= 9 && i <= imax; i++ {
		for j := 0; i+j <= ijmax && j <= 9; j++ {
			answer += get(9-i-j, 9-j, digits-2)
			answer += get(9-j, 9, digits-3)
			for leadingzeroes := 2; leadingzeroes <= digits-2; leadingzeroes++ {
				answer += get(9, 9, digits-2-leadingzeroes)
			}
		}
	}
	memo[[3]int{imax, ijmax, digits}] = answer

	return answer
}

var memo map[[3]int]int64

func main() {
	starttime := time.Now()

	memo = make(map[[3]int]int64)

	fmt.Println(get(9, 9, 20))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
