package main

import (
	"fmt"
	"time"
)

const (
	cubes       = 50000
	infinity    = 20000
	neginfinity = -10
)

var randMemo map[int]int

func S(k int) int {
	if ans, ok := randMemo[k]; ok {
		return ans
	}

	if k <= 55 {
		randMemo[k] = (100003 - (200003 * k) + (300007 * k * k * k)) % 1000000
		return S(k)
	}

	randMemo[k] = (S(k-24) + S(k-55)) % 1000000
	return S(k)
}

func sort(indices []int, val map[int]int) []int {
	if len(indices) == 1 {
		return indices
	}
	half := len(indices) / 2
	return merge(sort(indices[:half], val), sort(indices[half:], val), val)
}

//if ab sorted, merge them to return sorted slice
func merge(a, b []int, order map[int]int) []int {
	c := make([]int, len(a)+len(b))
	pa, pb, pc := 0, 0, 0
	for pa < len(a) || pb < len(b) {
		if pb == len(b) {
			c[pc] = a[pa]
			pa++
			pc++
		} else if pa == len(a) {
			c[pc] = b[pb]
			pb++
			pc++
		} else if order[a[pa]] < order[b[pb]] {
			c[pc] = a[pa]
			pa++
			pc++
		} else {
			c[pc] = b[pb]
			pb++
			pc++
		}

	}
	return c
}

func show(this int) {
	fmt.Println(xstart[this], ",", xend[this], "\t", ystart[this], ",", yend[this], "\t", zstart[this], ",", zend[this])

}

//assumes list sorted on f
func find(list []int, f map[int]int, min, max int) (start, end int, works bool) {
	if f[list[0]] > max || f[list[len(list)-1]] < min {
		return 0, 0, false
	}

	a, b := 0, len(list)-1

	for b-a > 1 {
		mid := (a + b) / 2

		if f[list[mid]] < min {
			a = mid

		} else {
			b = mid

		}
	}

	if f[list[a]] < min { ///usually
		start = b
	} else {
		start = a //at border
	}

	a, b = 0, len(list)-1

	for b-a > 1 {
		mid := (a + b) / 2

		if f[list[mid]] > max {
			b = mid

		} else {
			a = mid

		}
	}

	if f[list[b]] > max { ///usually
		end = a
	} else {
		end = b //at border
	}

	return start, end, true
}

//in place restrict list between a,b
func restrict(list []int, f map[int]int, min, max int) []int {
	list = sort(list, f)
	a, b, any := find(list, f, min, max)

	if any {
		return list[a : b+1]
	}

	return []int{}

}

func inits() {
	randMemo = make(map[int]int)

	xstart, xend = make(map[int]int), make(map[int]int)
	ystart, yend = make(map[int]int), make(map[int]int)
	zstart, zend = make(map[int]int), make(map[int]int)

	for i := 1; i <= cubes; i++ {
		x, y, z := S(6*i-5)%10000, S(6*i-4)%10000, S(6*i-3)%10000
		dx, dy, dz := 1+(S(6*i-2)%399), 1+(S(6*i-1)%399), 1+(S(6*i)%399)

		//After init, everything is zero indexed
		xstart[i-1], xend[i-1] = x, x+dx
		ystart[i-1], yend[i-1] = y, y+dy
		zstart[i-1], zend[i-1] = z, z+dz

		//fmt.Println(x, y, z, "\t", dx, dy, dz)
	}
}

var xstart, xend, ystart, yend, zstart, zend map[int]int

func super(n int, poss []int) int {
	poss = restrict(poss, xstart, neginfinity, xstart[n])
	poss = restrict(poss, xend, xend[n], infinity)

	poss = restrict(poss, ystart, neginfinity, ystart[n])
	poss = restrict(poss, yend, yend[n], infinity)

	poss = restrict(poss, zstart, neginfinity, zstart[n])
	poss = restrict(poss, zend, zend[n], infinity)

	if len(poss) == 1 {
		return -1
	}

	for _, in := range poss {
		if in != n {
			return in
		}
	}

	return -1

}

func enumerate(set map[int]bool) []int {
	ret := make([]int, 0)
	for key, val := range set {
		if val {
			ret = append(ret, key)
		}
	}
	return ret
}

//Idea: Find disjoint figures, then size those using
//[][][]bool with specified endpoints
func main() {
	starttime := time.Now()
	inits()

	ration := make(map[int]bool)
	for i := 0; i < cubes; i++ {
		ration[i] = true
	}

	for start, pres := range ration {
		//If this cube has been done continue
		if !pres {
			continue
		}

		//If I'm strictly in another cube continue
		if super(enumerate(ration), start) >= 0 {
			ration[start] = false
			continue
		}

		unit := make([]int, 1)
		unit[0] = start
		ration[start] = false

	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
