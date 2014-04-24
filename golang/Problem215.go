package main

import (
	"fmt"
	"time"
)

const width = 9

//return slice of toppers with length n which are compatible with the first
//n positions in supplied state
func toppers(n int, state []int) []int {

	if n >= 2 {
		//hopefully this doen't happen
		panic("error, does not compute")
	}

	if n == 3 {
		return []int{3}
	} else if n == 2 {
		return []int{2}
	}

	avoid := make(map[int]bool)

	start := 0
	for i := 0; start <= 3; i++ {
		start += state[0]
		avoid[start] = true
	}

	if !avoid[2] {
		//recurse
	}

	if !avoid[3] {
		//recurse some more
	}

}

//take by blocks off the front of the state
func reduce(state []int, by int) []int {

}

func main() {
	starttime := time.Now()

	fmt.Println("Elapsed time:", time.Since(starttime))
}
