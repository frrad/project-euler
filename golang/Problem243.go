package main

import (
	"euler"
	"fmt"
	"time"
)

const smooth = 2 * 3 * 5 * 7 * 11

func main() {
	starttime := time.Now()

	var i int64

	for i = smooth; euler.Totient(int64(i))*94744 > (i-1)*15499; i += smooth {

	}
	fmt.Println(i)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
