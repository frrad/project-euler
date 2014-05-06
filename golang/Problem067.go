package main

import (
	"euler"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	starttime := time.Now()

	triword := euler.Import("../problemdata/triangle.txt")
	triangle := make([][]int, 0)

	for _, line := range triword {
		bust := strings.Split(line, " ")
		numline := make([]int, 0)
		for _, letters := range bust {
			this, _ := strconv.Atoi(letters)
			numline = append(numline, this)
		}

		triangle = append(triangle, numline)

	}

	for j := len(triangle) - 2; j >= 0; j-- {
		for i := range triangle[j] {
			max := triangle[j+1][i]
			if triangle[j+1][i+1] > max {
				max = triangle[j+1][i+1]
			}
			triangle[j][i] += max
		}
	}

	fmt.Println(triangle[0][0])

	fmt.Println("Elapsed time:", time.Since(starttime))
}
