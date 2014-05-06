package main

import (
	"fmt"
	"time"
)

const (
	cubes       = 50000
	infinity    = 20000
	delta       = 200
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

	if f[list[a]] < min {
		start = b
	} else {
		start = a
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

	if f[list[b]] > max {
		end = a
	} else {
		end = b
	}

	if start > end {
		return 0, 0, false
	}

	if f[list[start]] < min && max < f[list[end]] {
		return 0, 0, false
	}

	return start, end, true
}

func restrict(list []int, f map[int]int, min, max int) []int {
	if len(list) == 0 {
		return list
	}

	sort(list, f)
	a, b, any := find(list, f, min, max)

	if any {
		return list[a : b+1]
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
	for key := range xpts {
		xlist = append(xlist, key)
	}
	for key := range ypts {
		ylist = append(ylist, key)
	}
	for key := range zpts {
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

	volume := int64(0)

	for _, cu := range collection {

		x1, x2, _ := find(xlist, xpts, xstart[cu], xend[cu])
		y1, y2, _ := find(ylist, ypts, ystart[cu], yend[cu])
		z1, z2, _ := find(zlist, zpts, zstart[cu], zend[cu])

		for i := x1; i < x2; i++ {
			for j := y1; j < y2; j++ {
				for k := z1; k < z2; k++ {

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

		xstart[i-1], xend[i-1] = x, x+dx
		ystart[i-1], yend[i-1] = y, y+dy
		zstart[i-1], zend[i-1] = z, z+dz
		vol[i-1] = dx * dy * dz

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

func eat(x int, set map[int]bool) []int {

	list := make([]int, 1)
	list[0] = x
	delete(set, x)

	for i := 0; i < len(list); i++ {
		cu := list[i]

		poss := enumerate(set)

		poss = restrict(poss, xstart, neginfinity, xend[cu]-1)
		poss = restrict(poss, xend, xstart[cu]+1, infinity)

		poss = restrict(poss, yend, ystart[cu]+1, infinity)
		poss = restrict(poss, ystart, neginfinity, yend[cu]-1)

		poss = restrict(poss, zend, zstart[cu]+1, infinity)
		poss = restrict(poss, zstart, neginfinity, zend[cu]-1)

		for _, cube := range poss {
			list = append(list, cube)
			delete(set, cube)
		}

	}

	return list

}

func cutx(list []int, cutposition, index int) (added int) {

	temp := restrict(list, xstart, neginfinity, cutposition-1)
	temp = restrict(temp, xend, cutposition+1, infinity)

	for _, box := range temp {

		xstart[index] = xstart[box]
		xend[index] = cutposition

		ystart[index] = ystart[box]
		yend[index] = yend[box]

		zstart[index] = zstart[box]
		zend[index] = zend[box]

		vol[index] = int(size([]int{index}))

		xstart[box] = cutposition
		vol[box] = int(size([]int{box}))

		added++
		index++

	}

	return
}

func partition(list []int, f map[int]int) int {
	pivotValue := f[list[len(list)/2]]

	list[len(list)-1], list[len(list)/2] = list[len(list)/2], list[len(list)-1]

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

func findVolume(volsort []int) (volume int64) {
	ration := make(map[int]bool)

	for _, cube := range volsort {
		ration[cube] = true
	}

	for _, start := range volsort {

		if !ration[start] {
			continue
		}

		unit := eat(start, ration)

		volume += size(unit)

	}

	return volume
}

func main() {
	starttime := time.Now()

	inits()
	fmt.Println("initialized")

	ration := make(map[int]bool)
	for i := 0; i < cubes; i++ {
		ration[i] = true
	}

	volume := int64(0)

	end := cubes
	for cuton := delta; cuton < infinity; cuton += delta {
		fmt.Println(cuton)
		start := end
		end = start + cutx(enumerate(ration), cuton, start)

		piece := make([]int, 0)

		for i := start; i < end; i++ {
			piece = append(piece, i)
			ration[i] = true
		}

		temp := enumerate(ration)
		temp = restrict(temp, xend, neginfinity, cuton)

		for _, box := range temp {
			delete(ration, box)
		}

		volume += findVolume(temp)
	}

	fmt.Println(volume)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
