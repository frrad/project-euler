package main

import (
	"fmt"
	"time"
)

//Pattern: Look at this picture
//Table[Table[If[Mod[Binomial[n, k], 7] == 0, " ", "X"], {k, 0, n}], {n, 0, 97}] // MatrixForm
func main() {
	starttime := time.Now()

	fmt.Println("Elapsed time:", time.Since(starttime))
}
