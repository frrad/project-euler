package main

import (
	"fmt"
	"time"
)

const (
	cubes       = 500
	infinity    = 20000
	neginfinity = 0
)

var randMemo map[int]int

func S(k int) int {
	if ans, ok := randMemo[k]; ok {
		return ans
	}

	if k <= 55 {
		kay := int64(k)
		temp := (100003 - (200003 * kay) + (300007 * kay * kay * kay)) % 1000000
		if temp < 0 {
			fmt.Println(k, temp)
		}
		randMemo[k] = int(temp)
		return S(k)
	}

	randMemo[k] = (S(k-24) + S(k-55)) % 1000000
	return S(k)
}

func sort(indices []int, val map[int]int) {

	if len(indices) > 1 {

		hack := partition(indices, val)
		sort(indices[:hack], val)
		sort(indices[1+hack:], val)
	}

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

	if start > end { // empty
		return 0, 0, false
	}

	if f[list[start]] < min && max < f[list[end]] { // empty
		return 0, 0, false
	}

	return start, end, true
}

//in place restrict list between a,b
func restrict(list []int, f map[int]int, min, max int) []int {
	if len(list) == 0 {
		return list
	}

	sort(list, f)
	a, b, any := find(list, f, min, max)

	if any {
		//fmt.Println("sliced:", a,b)
		return list[a : b+1] ///ASDFASDFA?????
	}

	return []int{}

}

func size(collection []int) int64 {
	xpts, ypts, zpts := make(map[int]int), make(map[int]int), make(map[int]int)
	for _, memb := range collection {
		a, b := xstart[memb], xend[memb]
		xpts[a], xpts[b] = a, b

		a, b = ystart[memb], yend[memb]
		ypts[a], ypts[b] = a, b

		a, b = zstart[memb], zend[memb]
		zpts[a], zpts[b] = a, b
	}

	xlist, ylist, zlist := make([]int, 0), make([]int, 0), make([]int, 0)
	for key, _ := range xpts {
		xlist = append(xlist, key)
	}
	for key, _ := range ypts {
		ylist = append(ylist, key)
	}
	for key, _ := range zpts {
		zlist = append(zlist, key)
	}

	sort(xlist, xpts)
	sort(ylist, ypts)
	sort(zlist, zpts)

	state := make([][][]bool, 0)
	for i := 0; i < len(xlist)-1; i++ {
		pancake := make([][]bool, 0)
		for j := 0; j < len(ylist)-1; j++ {
			row := make([]bool, len(zlist)-1)

			pancake = append(pancake, row)
		}

		state = append(state, pancake)

	}

	//fmt.Println(state)

	volume := int64(0)

	for _, cu := range collection {

		x1, x2, _ := find(xlist, xpts, xstart[cu], xend[cu])
		y1, y2, _ := find(ylist, ypts, ystart[cu], yend[cu])
		z1, z2, _ := find(zlist, zpts, zstart[cu], zend[cu])

		for i := x1; i < x2; i++ {
			for j := y1; j < y2; j++ {
				for k := z1; k < z2; k++ {
					//fmt.Println(len(xlist), len(ylist), len(zlist), i, j, k)
					if !state[i][j][k] {
						state[i][j][k] = true
						volume += int64(xlist[i+1]-xlist[i]) * int64(ylist[j+1]-ylist[j]) * int64(zlist[k+1]-zlist[k])
					}
				}
			}
		}

	}
	return volume

}

func inits() {
	randMemo = make(map[int]int)

	xstart, xend = make(map[int]int), make(map[int]int)
	ystart, yend = make(map[int]int), make(map[int]int)
	zstart, zend = make(map[int]int), make(map[int]int)
	vol = make(map[int]int)

	for i := 1; i <= cubes; i++ {
		x, y, z := S(6*i-5)%10000, S(6*i-4)%10000, S(6*i-3)%10000
		dx, dy, dz := 1+(S(6*i-2)%399), 1+(S(6*i-1)%399), 1+(S(6*i)%399)

		//After init, everything is zero indexed
		xstart[i-1], xend[i-1] = x, x+dx
		ystart[i-1], yend[i-1] = y, y+dy
		zstart[i-1], zend[i-1] = z, z+dz
		vol[i-1] = dx * dy * dz

		//fmt.Println(x, y, z, "\t", dx, dy, dz)
	}
}

var xstart, xend, ystart, yend, zstart, zend, vol map[int]int

func enumerate(set map[int]bool) []int {
	ret := make([]int, 0)
	for key, val := range set {
		if val {
			ret = append(ret, key)
		}
	}
	return ret
}

func renumerate(set map[int]bool, order []int) []int {
	ret := make([]int, 0)
	for _, val := range order {
		if set[val] {
			ret = append(ret, val)
		}
	}
	return ret
}

func eat(x1, x2, y1, y2, z1, z2 int, poss []int) []int {
	poss = restrict(poss, xstart, neginfinity, x2-1)
	poss = restrict(poss, xend, x1+1, infinity)

	poss = restrict(poss, yend, y1+1, infinity)
	poss = restrict(poss, ystart, neginfinity, y2-1)

	poss = restrict(poss, zend, z1+1, infinity)
	poss = restrict(poss, zstart, neginfinity, z2-1)

	return poss

}

func bound(input []int) (x1, x2, y1, y2, z1, z2 int) {
	x1, y1, z1 = infinity, infinity, infinity
	x2, y2, z2 = neginfinity, neginfinity, neginfinity
	for _, ind := range input {
		if x1 > xstart[ind] {
			x1 = xstart[ind]
		}
		if x2 < xend[ind] {
			x2 = xend[ind]
		}

		if y1 > ystart[ind] {
			y1 = ystart[ind]
		}
		if y2 < yend[ind] {
			y2 = yend[ind]
		}

		if z1 > zstart[ind] {
			z1 = zstart[ind]
		}
		if z2 < zend[ind] {
			z2 = zend[ind]
		}

	}

	return
}

//in-place quicksort
func partition(list []int, f map[int]int) int {
	pivotValue := f[list[len(list)/2]]                                          // Pivot on the middle
	list[len(list)-1], list[len(list)/2] = list[len(list)/2], list[len(list)-1] // put pivot at end for safekeeping

	storeIndex := 0
	for i := 0; i < len(list)-1; i++ {
		if f[list[i]] < pivotValue {
			list[i], list[storeIndex] = list[storeIndex], list[i]
			storeIndex++
		}

	}

	list[storeIndex], list[len(list)-1] = list[len(list)-1], list[storeIndex]
	return storeIndex
}

//Idea: Find disjoint figures, then size those using
//[][][]bool with specified endpoints
func main() {
	starttime := time.Now()

	inits()
	fmt.Println("initialized")

	ration := make(map[int]bool)
	for i := 0; i < cubes; i++ {
		ration[i] = true

	}

	volume := int64(0)

	fmt.Println(size(enumerate(ration)))

	volsort := enumerate(ration)
	sort(volsort, vol)
	fmt.Println("built volume sort")
	xssort := enumerate(ration)
	sort(xssort, xstart)
	fmt.Println("built xs sort")

	for i, start := range volsort {
		fmt.Println(i, "/", cubes)

		//If this cube has been done continue
		if !ration[start] {
			continue
		}

		unit := make([]int, 1)
		unit[0] = start
		ration[start] = false

		for {
			x1, x2, y1, y2, z1, z2 := bound(unit)

			add := eat(x1, x2, y1, y2, z1, z2, renumerate(ration, xssort))

			if len(add) == 0 {

				break
			}

			for _, cu := range add {
				delete(ration, cu)
				unit = append(unit, cu)
			}

		}

		if len(unit) == 1 {
			volume += int64(vol[start])

		}

		if len(unit) > 1 {

			fmt.Println("\nSTART")
			show(start)
			fmt.Println("===")
			for _, cu := range unit {
				show(cu)
			}

			fmt.Println("size:", size(unit))
			volume += size(unit)

		}

	}

	fmt.Println("Volume:", volume)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
