package main

import (
	"fmt"
	"strconv"
	"time"
)

type setup struct {
	binary string //this is "backwards"
	top    int
	bottom int
}

var table = make(map[setup]int)

func ways(input setup) int {
	if answer, ok := table[input]; ok {
		return answer
	}

	bin := trim(input.binary)
	answer := 0

	if lookup, ok := table[setup{bin, input.top, input.bottom}]; ok {
		answer = lookup
	} else if sum(bin) == 1 {
		if input.bottom <= len(bin) && input.top >= len(bin) {
			answer = len(bin) - input.bottom + 1
		} else if input.bottom <= len(bin) && input.top < len(bin) {
			fmt.Println("...")
			answer = input.top - input.bottom + 1
		} else {
			answer = 0
		}
	} else {
		front := build(len(bin))
		back := trim(bin[:len(bin)-1])
		fmt.Println(front, back)

		backa := ways(setup{back, len(back), input.bottom})
		backb := ways(setup{back, len(back) - 1, input.bottom})

		fronta := ways(setup{front, input.top, len(back) + 1})
		frontb := ways(setup{front, input.top, len(back)})

	}

	table[setup{bin, input.top, input.bottom}] = answer
	table[input] = answer
	return answer
}

func build(n int) (two string) {
	two = "1"
	for i := 0; i < n-1; i++ {
		two = "0" + two
	}
	return
}

func trim(a string) string {
	i := len(a) - 1
	for ; a[i:i+1] == "0"; i-- {
	}
	return a[:i+1]
}

func sum(a string) (total int) {
	for i := 0; i < len(a); i++ {
		digit, _ := strconv.Atoi(a[i : i+1])
		total += digit
	}
	return
}

func main() {
	starttime := time.Now()

	fmt.Println(ways(setup{"101", 4, 0}))

	fmt.Println("Elapsed time:", time.Since(starttime))

}
