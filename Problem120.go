package main

import (
	"fmt"
	"time"
)

const nlid = 100

//Binomial theorem
func f(a, n int) int {

	if n%2 == 0 {
		return 2
	}
	return (2 * n * a) % (a * a)
}

func main() {
	starttime := time.Now()

	fmt.Println("Elapsed time:", time.Since(starttime))

}
