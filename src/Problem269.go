package main

import (
	"euler"
	"fmt"
	"time"
)

const top = 100000

func eval(f, x int) (ans int) {
	pow := 0
	for f > 0 {
		coeff := f % 10
		f /= 10
		ans += coeff * int(euler.IntExp(int64(x), int64(pow)))
		pow++
	}
	return
}

func main() {
	starttime := time.Now()

	count := 0

	for i := 0; i < top; i++ {

		if i%10 == 0 {
			//fmt.Printf("%d\n", i)
			count++
		} else {
			flag := false
			for j := -1; j > -10; j-- {
				if eval(i, j) == 0 {
					if i%(10-j) != 0 {
						fmt.Printf("ALERT\n")
					}
					fmt.Printf("%d/%d=%d\n", i, 10-j, i/(10-j))

					flag = true
				}
			}
			if flag {
				fmt.Printf("%d\t%d\n", i, euler.Factor(int64(i)))

				count++
			}
		}

	}

	fmt.Printf("%d\n", count)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
