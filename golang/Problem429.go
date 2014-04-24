package main

import (
	"euler"
	"fmt"
	"time"
)

func main() {
	starttime := time.Now()

	fmt.Println("Hello, World", euler.Prime(10000))



	i:=100000000
	test := euler.IntPNum(i)
	fmt.Println(i,test)	
	


	fmt.Println("Elapsed time:", time.Since(starttime))


}