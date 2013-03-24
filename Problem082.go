package main

import (
	"./euler"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	starttime := time.Now()

	data := euler.Import("problemdata/matrix82.txt")

	matrix := make([][]int, len(data))

	for j, line := range data {
		words := strings.Split(line, ",")
		row := make([]int, len(words))
		for i, word := range words {
			row[i], _ = strconv.Atoi(word)
		}
		matrix[j] = row
	}

	matrix = make([][]int, 5)
	matrix[0] = []int{131, 673, 234, 103, 18}
	matrix[1] = []int{201, 96, 342, 965, 150}
	matrix[2] = []int{630, 803, 746, 422, 111}
	matrix[3] = []int{537, 699, 497, 121, 956}
	matrix[4] = []int{805, 732, 524, 37, 331}

	fmt.Println("Elapsed time:", time.Since(starttime))

}
