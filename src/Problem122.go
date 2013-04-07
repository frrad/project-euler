package main

import (
	"fmt"
	"strconv"
	"time"
)

func step(first []map[int]bool) []map[int]bool {
	kees := make(map[string]bool)
	answer := make([]map[int]bool, 0)
	for _, state := range first {
		//fmt.Println(blrp, "/", len(first))
		m := max(state)
		for i := range state {
			for j := range state {
				if !state[i+j] {
					new := clone(state)
					new[i+j] = true
					if max(new) > m && !kees[keygen(new)] {
						answer = append(answer, new)
						kees[keygen(new)] = true
					}
				}
			}

		}

	}
	return answer
}

func keygen(a map[int]bool) string {
	key := ""
	for i := 0; i < searchLength; i++ {
		if a[i] {
			key += strconv.Itoa(i)
		}
	}
	return key
}

func clone(old map[int]bool) map[int]bool {
	new := make(map[int]bool)
	for i := range old {
		new[i] = true
	}
	return new
}

func max(a map[int]bool) int {
	max := 0
	for x := range a {
		if x > max {
			max = x
		}
	}
	return max
}

func clean(list []map[int]bool) []map[int]bool {
	//fmt.Println("cleaning")
	clist := make([]map[int]bool, 0)
	for i := 0; i < len(list); i++ {

		inbounds := true

		for place := range list[i] {
			if place > searchLength {
				inbounds = false
				break
			}
		}

		if inbounds {
			clist = append(clist, list[i])
		}

	}
	//fmt.Println("cleaned")
	return clist
}

func same(a, b map[int]bool) bool {
	for i := range a {
		if b[i] == false {
			return false
		}
		if i > searchLength {
			return true
		}
	}
	return true
}

const (
	searchLength = 200 + 1
)

func main() {
	starttime := time.Now()

	canhit := make([]map[int]bool, 1)
	nada := make(map[int]bool)
	nada[1] = true
	canhit[0] = nada

	var results [searchLength]int

	done := false

	for i := 0; !done; i++ {

		for _, atlas := range canhit {
			for power := range atlas {
				if power < searchLength && results[power] == 0 {
					results[power] = i
				}
			}
		}
		canhit = step(canhit)
		canhit = clean(canhit)
		//fmt.Println(canhit)
		fmt.Println("SARCHED", i)
		fmt.Println(len(canhit), "WAYS")
		fmt.Println(results)

		results[1] = 0
		done = true
		sum := 0
		for j := 1; j < searchLength; j++ {
			if j > 1 && results[j] == 0 {
				done = false
			}
			sum += results[j]
		}
		fmt.Println("Sum:", sum)
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
