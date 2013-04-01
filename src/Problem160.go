package main

import (
	"fmt"
	"time"
)

var memo = make(map[int64][2]int64)

func fac(n int64) [2]int64 {
	if answer, ok := memo[n]; ok {
		return answer
	}

	base := fac(n - 1)[1]

	answer := fac(n - 1)[0]

	offset := int64(0)
	for n%10 == 0 {
		n = n / 10
		offset++

	}

	answer *= n

	for answer%10 == 0 {
		answer = answer / 10
		offset++
	}

	answer = answer % 100000

	memo[n] = [2]int64{answer, base + offset}

	//hogs memory otherwise
	delete(memo, n-5)

	return [2]int64{answer, base + offset}

}

const target = 1000000000000

func main() {
	starttime := time.Now()

	memo[0] = [2]int64{1, 0}

	for i := int64(0); i <= target; i++ {

		if i%1000000 == 0 {
			fmt.Println(100*float64(i)/float64(target), "\b%", i, fac(i))
		}

	}

	fmt.Println(fac(target))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
