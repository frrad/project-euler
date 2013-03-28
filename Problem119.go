package main

import (
	"fmt"
	"strconv"
	"time"
)

func seek(start int64, c chan int64) {
	for i := start; ; i++ {

		sum := int64(0)
		word := strconv.FormatInt(i, 10)
		for i := 0; i < len(word); i++ {
			digit, _ := strconv.Atoi(word[i : i+1])
			sum += int64(digit)
		}
		if sum > 1 {

			power := sum * sum
			for exp := int64(2); power <= i; exp++ {
				power = sum
				for j := int64(0); j < exp-1; j++ {
					power *= sum
				}

				if power == i {
					c <- i
				}
			}
		}
	}

}

func main() {
	starttime := time.Now()

	c, d := make(chan int64), make(chan int64)

	go seek(12, c)

	go seek(614656, d)

	for {
		select {
		case x := <-c:
			fmt.Println(x)
		case x := <-d:
			fmt.Println(x)
		}
	}

	fmt.Println("Elapsed time:", time.Since(starttime))

}
