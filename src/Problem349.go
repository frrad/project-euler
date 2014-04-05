package main

import (
	"fmt"
	"time"
)

const field = 30

func show(hill *[field][field]bool, x, y, a, b int) {
	for i := x; i < x+a; i++ {
		for j := y; j < y+b; j++ {
			if hill[i][j] {
				fmt.Printf("X")
			} else {
				fmt.Printf(" ")
			}

		}
		fmt.Printf("|\n")

	}
}

func step(hill *[field][field]bool, place, dir [2]int) ([2]int, [2]int) {
	x, y := place[0], place[1]
	color := hill[x][y]

	if color {
		hill[x][y] = !color
		dir = CCW(dir)
		x += dir[0]
		y += dir[1]
	} else {
		hill[x][y] = !color
		dir = CW(dir)
		x += dir[0]
		y += dir[1]
	}

	return [2]int{x, y}, dir

}

func CCW(dir [2]int) [2]int {
	x, y := dir[0], dir[1]
	newX := 0*x + -1*y
	newY := 1*x + 0*y
	return [2]int{newX, newY}
}

func CW(dir [2]int) [2]int {
	x, y := dir[0], dir[1]
	newX := 0*x + 1*y
	newY := -1*x + 0*y
	return [2]int{newX, newY}
}

func main() {
	starttime := time.Now()

	var hill [field][field]bool
	loc := [2]int{field / 2, field / 2}
	dir := [2]int{1, 0}

	for i := 0; i < 2000; i++ {
		fmt.Printf("step %d\n", i)

		show(&hill, 0, 0, field, field)
		loc, dir = step(&hill, loc, dir)
	}
	fmt.Println("Elapsed time:", time.Since(starttime))
}
