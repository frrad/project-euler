package main

import (
	"euler"
	"fmt"
	"time"
)

const top int64 = 80

func rep(x int64) [][]int64 {
	facs := euler.Factors(x)
	d1 := euler.IntExp(facs[0][0], facs[0][1])

	if len(facs) == 1 {
		return [][]int64{{1, d1}}
	}

	d2 := (x) / d1
	num1, num2 := euler.ExtendedEuclidean(d2, d1)
	up := times(rep(d2), num2)
	return append(up, []int64{num1, d1})
}

//helper function for rep
func times(fracs [][]int64, mult int64) [][]int64 {
	for i, val := range fracs {
		fracs[i][0] = val[0] * mult
	}
	return fracs
}

func compute(given []int64, pins [][2]int64, modulus int64, target int64) [][]int64 {
	concern := make([][2]int64, 0)
	for _, pin := range pins {
		ind, modifier := pin[0], pin[1]

		if given[ind] == 0 {
			concern = append(concern, pin)
		}
		if given[ind] == 1 {
			target -= modifier
		}

	}

	lambdas := make([]int64, len(concern))
	for i := 0; i < len(concern); i++ {
		lambdas[i] = concern[i][1]
	}

	results := dumbNumerate(lambdas, target, modulus)

	adjusted := make([][]int64, len(results))
	for i := 0; i < len(adjusted); i++ {
		adjusted[i] = make([]int64, len(given))
		copy(adjusted[i], given)

		for j, tuple := range concern {
			if results[i][j] {
				adjusted[i][tuple[0]] = 1
			} else {
				adjusted[i][tuple[0]] = -1
			}
		}
	}

	return adjusted
}

func dumbNumerate(vals []int64, target, modulus int64) (ans [][]bool) {
	ans = make([][]bool, 0)

	for state := 0; state < 1<<uint(len(vals)); state++ {

		crunch := int64(0)

		carrot := recover(state, len(vals))

		for i, on := range carrot {

			if on {
				crunch += vals[i]
			}

		}

		if (modulus+(crunch%modulus))%modulus == (modulus+(target%modulus))%modulus {
			ans = append(ans, carrot)
		}
	}

	return
}

func recover(ind, lent int) []bool {
	ans := make([]bool, lent)
	for i := 0; i < lent; i++ {
		if ind%2 == 1 {
			ans[i] = true
		}
		ind /= 2

	}
	return ans
}

func compatible(a, b []int64) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == 0 || b[i] == 0 {
			continue
		}

		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	starttime := time.Now()

	topPrime := euler.PrimePi(top)
	primeBasis := make([]int64, topPrime+1)

	//first compile a list of prime power common denominators
	for i := int64(2); i <= top; i++ {
		repr := rep(i * i)

		for _, frac := range repr {
			den := frac[1]
			pType := euler.Factor(den)[0]
			pDex := euler.PrimePi(pType)

			if den > primeBasis[pDex] {
				primeBasis[pDex] = den
			}
		}

	}

	canonical := make([][]int64, top+1)

	for i := int64(2); i <= top; i++ {
		standard := make([]int64, topPrime+1)

		repr := rep(i * i)

		for _, frac := range repr {
			den := frac[1]
			pType := euler.Factor(den)[0]
			pDex := euler.PrimePi(pType)
			offset := primeBasis[pDex] / den
			standard[pDex] = offset * frac[0]
		}

		canonical[i] = standard

	}

	predicates := make([][]int64, 1)
	predicates[0] = make([]int64, top+1)
	// predicates[0][2] = 1 //this is true, not much speedup though

	place := topPrime

	for place = topPrime; place >= 1; place-- {

		// fmt.Println("at", place, "have", len(predicates))

		crinkTab := make([][2]int64, 0)

		for i, rep := range canonical {
			if len(rep) >= int(place+1) && rep[place] != 0 {
				modularity := rep[place] % primeBasis[place]
				crinkTab = append(crinkTab, [2]int64{int64(i), modularity})
			}
		}

		newPreds := make([][]int64, 0)

		for _, pred := range predicates {
			target := int64(0)
			if place == 1 {
				target = primeBasis[1] / 2
			}
			newPreds = append(newPreds, compute(pred, crinkTab, primeBasis[place], target)...)
		}

		predicates = newPreds

	}

	// for _, ans := range predicates {
	// 	for i, is := range ans {
	// 		if is == 1 {
	// 			fmt.Print("1/", i, "^2+")
	// 		}
	// 	}
	// 	fmt.Print("\b \n")
	// }

	fmt.Println(len(predicates))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
