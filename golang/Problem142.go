package main

import (
	"euler"
	"fmt"
	"time"
)

const limit2 = 5000
const count = 200000

func main() {
	starttime := time.Now()

	C2Table := make(map[int64]map[int64]bool)

	for a := int64(1); a < limit2; a++ {
		for b := a; b < limit2; b++ {
			c2 := a*a + b*b
			if euler.IsSquare(c2) {
				if enter, ok := C2Table[c2]; ok {
					enter[a*a] = true
					enter[b*b] = true
				} else {
					C2Table[c2] = make(map[int64]bool, 0)
					C2Table[c2][a*a] = true
					C2Table[c2][b*b] = true
				}

			}

		}
	}

	fmt.Println("Table Built", time.Since(starttime))

	min := int64(3*count + 3*limit2*limit2)

	for x := int64(1); x < count; x++ {

		if x%(count/20) == 0 {
			fmt.Println("x=", x)
		}

		for c2, poss := range C2Table {

			z := x + c2

			if x+x+z < min && euler.IsSquare(x+z) { //x+x+z < x+y+z < min
				for a2, _ := range poss {
					y := x + a2
					if x+y+z < min && euler.IsSquare(x+y) && euler.IsSquare(z+y) {

						min = x + y + z
						fmt.Println(x, "+", y, "+", z, "=", x+y+z)

					}

				}
			}
		}

	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
