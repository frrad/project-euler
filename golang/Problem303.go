package main

import (
	"fmt"
	"time"
)

const top uint64 = 10000

func check(n uint64) bool {
	if n < 10 {
		return n == 0 || n == 2 || n == 1
	}

	least := n % 10
	rest := n / 10

	return (least == 0 || least == 2 || least == 1) && check(rest)
}

func first(a uint64) uint64 {
	i := uint64(1)
	for ; !check(i * a); i++ {
	}

	return i
}

func main() {
	starttime := time.Now()

	sum := uint64(0)

	for i := uint64(1); i <= top; i++ {
		multiple := first(i)
		sum += multiple
		fmt.Printf("i:%-5d multiple:%-15d total:%d\n", i, multiple, sum)
	}

	fmt.Println(sum)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
