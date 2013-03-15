package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const zero = .001

func cancel(num int, denom int) (newnumerator int, newdenominator int) {
	newnumerator = num
	newdenominator = denom
	s := strconv.Itoa(num)
	slist := strings.Split(s, "")
	t := strconv.Itoa(denom)
	tlist := strings.Split(t, "")

	if slist[1] == "0" {
		return
	}

	if slist[0] == tlist[0] {
		newnumerator, _ = strconv.Atoi(slist[1])
		newdenominator, _ = strconv.Atoi(tlist[1])

	} else if slist[0] == tlist[1] {
		newnumerator, _ = strconv.Atoi(slist[1])
		newdenominator, _ = strconv.Atoi(tlist[0])
	} else if slist[1] == tlist[0] {
		newnumerator, _ = strconv.Atoi(slist[0])
		newdenominator, _ = strconv.Atoi(tlist[1])
	} else if slist[1] == tlist[1] {
		newnumerator, _ = strconv.Atoi(slist[0])
		newdenominator, _ = strconv.Atoi(tlist[0])

	}

	return
}

func reverse(n int) int {
	s := strconv.Itoa(n)

	var reversed string

	for i := len(s) - 1; i >= 0; i-- {
		reversed += s[i : i+1]
	}

	m, _ := strconv.Atoi(reversed)
	return m
}

func isSame(a int, b int, c int, d int) bool {
	if math.Abs(float64(a)/float64(b)-float64(c)/float64(d)) < zero {
		return true
	}
	return false

}

func main() {

	for i := 10; i < 100; i++ {

		for j := i + 1; j < 100; j++ {

			a, b := cancel(i, j)
			if a != i && isSame(a, b, i, j) {
				fmt.Println(a, b, i, j)

			}

		}
	}

}
