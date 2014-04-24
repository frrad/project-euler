package main

import (
	"fmt"
	"math"
	"time"
)

//gives pipe length for config
func length(config []int) (sum float64) {

	for i := 0; i < len(config)-1; i++ {
		a, b := float64(config[i]), float64(config[i+1])
		sum += 10 * math.Sqrt(2) * math.Sqrt(a+b-50)
	}

	sum += float64(config[0])
	sum += float64(config[len(config)-1])

	return
}

func opt(from, to int) []int {
	ans := make([]int, 0)

	for i := from; i <= to; i++ {

		if i%2 == 0 {
			ans = append(ans, i)
		} else {
			ans = append([]int{i}, ans...)
		}

	}

	return ans
}

func main() {
	starttime := time.Now()

	answer := length(opt(30, 50))

	fmt.Println(int(answer * 1000))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
