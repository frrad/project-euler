package main

import (
	"euler"
	"fmt"
	"time"
)

//quantity of admissable numbers \leq n with top prime p or below
//convention : 1 is not admissable
func admissable(n, p int64) int64 {
	if p == 1 && n >= 2 {
		return 2
	}else if p== 1 && n < 2 {
		return 1
	}
	answer := admissable(n, p-1) 
	if euler.Prime(p) <= n {
		answer += admissable(n/ euler.Prime(p) , p-1)
	}
	return answer
}

func main() {
	starttime := time.Now()

	a := 369000
	n := uint64(128)

	x := euler.IntPNum(a);
	

for i := uint64(0); i < (1+x.Divisors())/2; i++ {
	
	d := x.Divisor(i)
	q := euler.Quotient(x,d)
	fmt.Println(a, d.UInt64(),q.UInt64())

}


	fmt.Println(a, "mod", n, x.Mod(n))

	//Candidates can't be divisible by the square of a prime!




	fmt.Println("Elapsed time:", time.Since(starttime))


}
