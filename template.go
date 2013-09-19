package main

import (
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	fmt.Println("Elapsed time:", time.Since(starttime))
}
