package main

import (
	"euler"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func isSpecial(set []int) bool {
	length := len(set)
	a, b := make([]int, length), make([]int, length)
	for i := int64(0); i < euler.Factorial(int64(length)); i++ {
		return false
	}
}

func sum(set []int) (total int) {
	for _, x := range set {
		total += x
	}
	return
}

func main() {
	starttime := time.Now()

	data := euler.Import("../problemdata/sets.txt")
	sets := make([][]int, len(data))

	for i, line := range data {
		sets[i] = make([]int, 0)
		for _, word := range strings.Split(line, ",") {
			number, _ := strconv.Atoi(word)
			sets[i] = append(sets[i], number)
		}
	}

	total := 0

	for _, set := range sets {
		if isSpecial(set) {
			fmt.Println(set)
			total += sum(set)
		}
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
