package main

import (
	"euler"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	fmt.Println("Hello, World", euler.Prime(10000))

	/*
		finishes[step_] :=
		  If[step == 1,
		   N[((2^step - 1)/2^step)^32, 11],

		   N[((2^step - 1)/2^step)^32 - ((2^(step - 1) - 1)/2^(step - 1))^32, 
		    11]
		   ];
		N[Total[Table[finishes[k]*k, {k, 1, 1000}]], 11]
	*/

	fmt.Println("Elapsed time:", time.Since(starttime))
}
