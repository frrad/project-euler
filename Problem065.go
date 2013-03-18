package main

import (
	"./eulerlib"
	"fmt"
)

func ctdFrac(list []int64) (num int64, den int64) {
	num = list[len(list)-1]
	den = 1

	for i := len(list) - 2; i >= 0; i-- {

		num, den = eulerlib.FracReduce(eulerlib.FracAdd(list[i], 1, den, num))

		fmt.Println(num, den)
	}

	return
}

func min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func eList(n int) []int64 {
	answer := make([]int64, n)

	answer[0] = 2

	for i := 1; i < n; i++ {
		answer[i] = 1
	}

	for i := int64(0); 3*i+2 < int64(n); i++ {
		answer[3*i+2] = int64(2) * (i + int64(1))

	}

	return answer
}

func main() {

	fmt.Println(eulerlib.StringSum("0", "2345654322"))

	numerator, _ := ctdFrac(eList(30))
	fmt.Println(eulerlib.DigitSum(numerator))

}
