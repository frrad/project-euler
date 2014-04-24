package main

import (
	"euler"
	"fmt"
	"time"
)

const size = 12

func binary(index, size int) []int {
	if size == 0 {
		return []int{}
	}
	return append(binary(index/2, size-1), index%2)

}

func main() {
	starttime := time.Now()

	count := 0
	compare := make([][][]int, size)

	for i := 1; i < int(euler.IntExp(int64(2), size)); i++ {
		set := binary(i, size)

		blip := make([]int, 0)

		drip := 0
		for a := 0; a < size; a++ {
			if set[a] == 1 {
				drip++
				blip = append(blip, a)
			}
		}

		compare[drip-1] = append(compare[drip-1], blip)

	}

	for i := 1; i < size; i++ {

		for a := 0; a < len(compare[i]); a++ {
			for b := a + 1; b < len(compare[i]); b++ {
				t1 := compare[i][a]
				t2 := compare[i][b]

				elim := false

				for x := 0; x < len(t1); x++ {

					for y := 0; y < len(t1); y++ {

						if t1[x] == t2[y] {
							elim = true
						}

					}
				}

				test := true

				for q := 0; q < len(t1); q++ {
					if t1[q] < t2[q] {
						test = false
						break
					}
				}

				elim = test || elim

				if !elim {
					//fmt.Println(t1, t2)
					count++

				}

			}
		}

	}
	fmt.Println(count)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
