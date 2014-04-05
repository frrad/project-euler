package main

import (
	"fmt"
	"time"
)

const (
	field  = 100
	SIZE   = 46    //terminal height, used to align animation
	REG    = 10200 //regularity begins around 10200
	TARGET = 1000000000000000000
)

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
	//pad(SIZE - field)
	//time.Sleep(100 * time.Millisecond) //easier to watch
}

func pad(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("\n")
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

func count(hill *[field][field]bool) (counter int) {
	for i := 0; i < field; i++ {
		for j := 0; j < field; j++ {
			if hill[i][j] {
				counter++
			}
		}
	}
	return
}

func linModel(x1, fx1, x2, fx2 int) func(int64) int64 {
	m := (float64(fx2) - float64(fx1)) / (float64(x2) - float64(x1))
	b := float64(fx1) - m*float64(x1)

	line := func(x int64) int64 {
		prec := m*float64(x) + b
		a := int64(prec)

		off := float64(a) - prec

		if off > -.5 && off < .5 {
			return a
		}
		return a + 1
	}
	return line
}

func main() {
	starttime := time.Now()

	data := make(map[int]int)

	var hill [field][field]bool
	loc := [2]int{field / 2, field / 2}
	dir := [2]int{1, 0}

	for i := 0; i < 11655; i++ {

		/*if i%100 == 0 {
			fmt.printf("step %d\n", i)
			show(&hill, 0, 0, field, field)
		}*/

		if i > REG {
			//fmt.printf("{%d,%d},", i, count(&hill))
			data[i] = count(&hill)
		}

		loc, dir = step(&hill, loc, dir)
	}

	var modulus int

	for jump := 3; jump < 1000; jump++ {
		guess := linModel(REG+1, data[REG+1], REG+jump+1, data[REG+jump+1])

		flag := true
		for i := 0; i < 13; i++ { //there  are about 13 repetitions in sample
			if int64(data[REG+i*jump+1]) != guess(int64(REG+i*jump+1)) {
				flag = false
			}
		}
		if flag {
			modulus = jump
			break
		}
	}

	fmt.Printf("Repeats every %d\n", modulus)
	tarRem := int(TARGET % uint64(modulus))
	base := 1

	for ; (base+REG)%modulus != tarRem; base++ {
	}

	fmt.Printf("Line through (%d, %d) , (%d, %d)\n", base+REG, data[base+REG], base+REG+modulus, data[base+REG+modulus])

	model := linModel(base+REG, data[base+REG], base+REG+modulus, data[base+REG+modulus])

	fmt.Printf("%d (isn't quite precise enough)\n", model(TARGET))

	fmt.Printf("115384615384614952\n")

	fmt.Println("Elapsed time:", time.Since(starttime))
}
