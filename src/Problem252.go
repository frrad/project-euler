package main

import (
	//"euler"
	"fmt"
	"time"
)

const numpoints = 20

var SMemo map[int]int

type polygon struct {
	points   [][2]float64
	interior [2]float64
}

func S(n int) int {
	if answer, ok := SMemo[n]; ok {
		return answer
	}
	SMemo[n] = (S(n-1) * S(n-1)) % 50515093
	return S(n)
}

func T(n int) int {
	return (S(n) % 2000) - 1000
}

func (a *polygon) contains(dot [2]float64) bool {
	for i := 0; i < len(a.points); i++ {

		if !sameSide(a.interior, dot, a.points[i], a.points[(i+1)%len(a.points)]) {
			return false
		}

	}
	return true
}

//do a and b lie on the same side of line through x and y?
func sameSide(a, b, x, y [2]float64) bool {
	m := (x[1] - y[1]) / (x[0] - y[0])
	B := x[1] - (m * x[0])

	//fmt.Println("y=", m, "x+", B, "is line through", x, y)

	if a[1] < (m*a[0])+B && b[1] < (m*b[0])+B {
		return true
	}

	if a[1] > (m*a[0])+B && b[1] > (m*b[0])+B {
		return true
	}

	return false

}

func Make(pts [][2]float64) *polygon {
	poly := new(polygon)
	poly.points = make([][2]float64, len(pts))
	copy(poly.points, pts)
	poly.interior = [2]float64{(pts[0][0] + pts[1][0] + pts[2][0]) / 3,
		(pts[0][1] + pts[1][1] + pts[2][1]) / 3}
	return poly
}

func main() {
	starttime := time.Now()

	SMemo = make(map[int]int)
	SMemo[0] = 290797
	points := make([][2]float64, 0)
	for i := 1; i < 2*numpoints; i += 2 {
		points = append(points, [2]float64{float64(T(i)), float64(T(i + 1))})
	}

	seen := make(map[[3][2]float64]bool)

	for i := 0; i < numpoints; i++ {
		for j := i + 1; j < numpoints; j++ {
			for k := j + 1; k < numpoints; k++ {

				if seen[[3][2]float64{points[i], points[j], points[k]}] {
					continue
				}

				triangle := Make([][2]float64{points[i], points[j], points[k]})
				fmt.Println(triangle)

				broke := false

				for l := 0; l < numpoints; l++ {
					for l == i || l == j || l == k {
						l++
					}
					if l == numpoints {
						continue
					}

					if triangle.contains(points[l]) {
						broke = true
						break
					}

				}

				if broke {
					continue
				}

				fmt.Println(triangle)

				//this is a valid triangle expand via greedy then measure area

			}
		}
	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
