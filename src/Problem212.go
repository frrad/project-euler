package main

import (
	"fmt"
	"time"
)

const (
	cubes       = 50000
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
		if temp < 0{ fmt.Println(k, temp)}
		randMemo[k] = int(temp)
		return S(k)
	}

	randMemo[k] = (S(k-24) + S(k-55)) % 1000000
	return S(k)
}

func sort(indices []int, val map[int]int) []int {
	if len(indices) == 1 {return indices}

	for i := 0 ; i < len(indices) ; i++{
		for j:=0 ; j< len(indices) -1 ; j++{
			if val[indices[j]] > val[indices[j+1]]{
			indices[j] , indices[j+1] =  indices[j+1], indices[j]
			}
		}

	}	

	return indices


	/* //mergesort hard on memory?
	if len(indices) == 1 {
		return indices
	}
	half := len(indices) / 2
	return merge(sort(indices[:half], val), sort(indices[half:], val), val)
		*/
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

if start > end { // empty
	return 0,0,false
}

if min < f[list[start]] && f[list[end]] < max{ // empty
	return 0,0,false
}

	return start, end, true
}

//in place restrict list between a,b
func restrict(list []int, f map[int]int, min, max int) []int {
	if len(list)==0{return list}
	list = sort(list, f)
	a, b, any := find(list, f, min, max)

	if any  {
		//fmt.Println("sliced:", a,b)
		return list[a : b+1] ///ASDFASDFA?????
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

func eat(x1, x2, y1, y2, z1, z2 int, poss []int) []int {
	temp := make(map[int]bool)
	ans := make([]int, 0)

	for _, cu := range poss {
		temp[cu] = true
	}

	//fmt.Println("eating")

	poss = enumerate(temp)
	poss = restrict(poss, xstart, x1, x2-1)
	poss = restrict(poss, yend, y1, infinity)
	poss = restrict(poss, ystart, neginfinity, y2)
	poss = restrict(poss, zend, z1, infinity)
	poss = restrict(poss, zstart, neginfinity, z2)
	for _, cu := range poss {
		temp[cu] = false
		ans = append(ans, cu)
	}

	//fmt.Println("eating1")

	poss = enumerate(temp)
	poss = restrict(poss, xend, x1+1, x2)
	poss = restrict(poss, yend, y1, infinity)
	poss = restrict(poss, ystart, neginfinity, y2)
	poss = restrict(poss, zend, z1, infinity)
	poss = restrict(poss, zstart, neginfinity, z2)
	for _, cu := range poss {
		temp[cu] = false
		ans = append(ans, cu)
	}

	//fmt.Println("eating2")

	poss = enumerate(temp)
	poss = restrict(poss, ystart, y1, y2-1)
	poss = restrict(poss, xend, x1, infinity)
	poss = restrict(poss, xstart, neginfinity, x2)
	poss = restrict(poss, zend, z1, infinity)
	poss = restrict(poss, zstart, neginfinity, z2)
	for _, cu := range poss {
		temp[cu] = false
		ans = append(ans, cu)
	}

	poss = enumerate(temp)
	poss = restrict(poss, yend, y1+1, y2)
	poss = restrict(poss, xend, x1, infinity)
	poss = restrict(poss, xstart, neginfinity, x2)
	poss = restrict(poss, zend, z1, infinity)
	poss = restrict(poss, zstart, neginfinity, z2)
	for _, cu := range poss {
		temp[cu] = false
		ans = append(ans, cu)
	}

	poss = enumerate(temp)
	poss = restrict(poss, zstart, z1, z2-1)
	poss = restrict(poss, xend, x1, infinity)
	poss = restrict(poss, xstart, neginfinity, x2)
	poss = restrict(poss, yend, y1, infinity)
	poss = restrict(poss, ystart, neginfinity, y2)
	for _, cu := range poss {
		temp[cu] = false
		ans = append(ans, cu)
	}

	poss = enumerate(temp)
	poss = restrict(poss, zend, z1+1, z2)
	poss = restrict(poss, xend, x1, infinity)
	poss = restrict(poss, xstart, neginfinity, x2)
	poss = restrict(poss, yend, y1, infinity)
	poss = restrict(poss, ystart, neginfinity, y2)
	for _, cu := range poss {
		temp[cu] = false
		ans = append(ans, cu)
	}

	fmt.Println("eaten", ans)
	return ans

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

	for start, pres := range ration {
		//If this cube has been done continue
		if !pres {
			continue
		}

		//If I'm strictly in another cube continue
		if super(start, enumerate(ration)) >= 0 {
			ration[start] = false
			continue
		}

		unit := make([]int, 1)
		unit[0] = start
		ration[start] = false

		fmt.Println("START")
		show(start)

		for {
			fmt.Println("bounding")
			x1, x2, y1, y2, z1, z2 := bound(unit)
			fmt.Println("bounded")

			add := eat(x1, x2, y1, y2, z1, z2, enumerate(ration))

			if len(add) == 0 {
				break
			}

			fmt.Println("===")
			for _, cu := range add {
				show(cu)
				ration[cu] = false
				unit = append(unit, cu)
			}
	if len(unit) > 1{panic("here")}
		}

		fmt.Println("unit:", unit)

	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
