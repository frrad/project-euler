package main

import (
	"fmt"
	"time"
)

const dim = 7
const bishes = 4

func works(a [dim]int64) bool {

	ans := -1-a[0]-a[1]-a[2]-a[3]-a[4]-a[5]-a[6] > 0 && 128+a[0]+2*a[1]+4*a[2]+8*a[3]+16*a[4]+32*a[5]+64*a[6] > 0 && -2187-a[0]-3*a[1]-9*a[2]-27*a[3]-81*a[4]-243*a[5]-729*a[6] > 0 && 16384+a[0]+4*a[1]+16*a[2]+64*a[3]+256*a[4]+1024*a[5]+4096*a[6] > 0 && -78125-a[0]-5*a[1]-25*a[2]-125*a[3]-625*a[4]-3125*a[5]-15625*a[6] > 0 && 279936+a[0]+6*a[1]+36*a[2]+216*a[3]+1296*a[4]+7776*a[5]+46656*a[6] > 0 && -823543-a[0]-7*a[1]-49*a[2]-343*a[3]-2401*a[4]-16807*a[5]-117649*a[6] > 0 && 2097152+a[0]+8*a[1]+64*a[2]+512*a[3]+4096*a[4]+32768*a[5]+262144*a[6] > 0

	return ans
}

func sum(a, b [dim]int64) [dim]int64 {
	sum := [dim]int64{}
	for i := 0; i < dim; i++ {
		sum[i] = a[i] + b[i]
	}
	return sum
}

func abs(an int64) int64 {
	if an < 0 {
		return -1 * an
	}
	return an
}

func main() {
	starttime := time.Now()

	seen := make(map[[dim]int64]bool)
	work := make(map[[dim]int64]bool)
	tocheck := make(map[[dim]int64]bool)

	start := [dim]int64{-10872, 26312, -24159, 11206, -2888, 419, -32}
	seen[start] = true
	work[start] = true
	tocheck[start] = true

	for len(tocheck) > 0 {
		for key, _ := range tocheck {

			for a := int64(-1 * bishes); a <= 1*bishes; a++ {

				for b := int64(-1 * bishes); b <= 1*bishes; b++ {

					for c := int64(-1 * bishes); c <= 1*bishes; c++ {

						for d := int64(-1 * bishes); d <= 1*bishes; d++ {

							for e := int64(-1 * bishes); e <= 1*bishes; e++ {

								for f := int64(-1 * bishes); f <= 1*bishes; f++ {

									for g := int64(-1 * bishes); g <= 1*bishes; g++ {

										tickle := sum(key, [dim]int64{a, b, c, d, e, f, g})

										if seen[tickle] {
											continue
										}

										seen[tickle] = true

										if works(tickle) {
											work[tickle] = true
											tocheck[tickle] = true
										}

									}
								}
							}
						}
					}
				}
			}

			delete(tocheck, key)
		}
		fmt.Println(work)

	}

	toto := int64(0)

	for key, _ := range work {
		for i := 0; i < dim; i++ {
			toto += abs(key[i])
		}
	}

	fmt.Println(toto)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
