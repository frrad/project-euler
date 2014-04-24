package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

const top int = 100

func main() {
	starttime := time.Now()

	fac := new(big.Int)
	fac.MulRange(int64(1), int64(top))

	str := fac.String()
	total := 0

	for _, letter := range str {
		digit, _ := strconv.Atoi(string(letter))
		total += digit
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
