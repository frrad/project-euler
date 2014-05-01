package main

import (
	"euler"
	"fmt"
	"time"
)

const (
	top   = 6
	snack = 9
)

func reverse(list []int) {
	lth := len(list)
	for i := 0; i < lth/2; i++ {
		list[i], list[lth-i-1] = list[lth-i-1], list[i]
	}
}

//move n to the front. how many moves does this take?
func front(n int, list []int) (moves int) {
	if list[0] == n {
		return 0
	}

	if list[len(list)-1] == n {
		reverse(list)
		return 1
	}

	i := 0
	for ; list[i] != n; i++ {
	}

	reverse(list[i:])

	return 1 + front(n, list)
}

func sort(list []int) (total int) {
	for i := 0; i < len(list); i++ {
		total += front(i, list[i:])
	}
	return
}

func main() {
	starttime := time.Now()

	template := make([]int, top)
	for i := 0; i < top; i++ {
		template[i] = i
	}

	for i := 0; i < int(euler.Factorial(top)); i++ {
		test := make([]int, top)
		copy(test, template)

		test = euler.Permutation(i, test)

		if sort(test) == snack {
			copy(test, template)
			fmt.Println(euler.Permutation(i, test), snack)
		}

	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
