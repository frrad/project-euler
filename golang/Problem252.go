package main

import (
	//"euler"
	"fmt"
	"time"
)

const numpoints = 500

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

//is dot strictly inside a
func (a *polygon) contains(dot [2]float64) bool {
	for _, vertex := range a.points {

		if equal(vertex, dot) {
			return false
		}
	}

	for i := 0; i < len(a.points); i++ {

		if !sameSide(a.interior, dot, a.points[i], a.points[(i+1)%len(a.points)]) {
			return false
		}

	}
	return true
}

//do a and b lie on the same side of line through x and y?
func sameSide(a, b, x, y [2]float64) bool {

	if x[0] == y[0] { //Vertical line
		ex := x[0]
		if a[0] > ex && b[0] > ex {
			return true
		}
		if a[0] < ex && b[0] < ex {
			return true
		}
		return false
	}

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

func equal(x, y [2]float64) bool {
	if x[0] == y[0] && x[1] == y[1] {
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

func (t *polygon) containsAny(list [][2]float64) bool {

	for _, point := range list {

		if t.contains(point) {
			return true
		}

	}
	return false
}

//add point after position index
func (p *polygon) add(point [2]float64, index int) {
	//p.print()
	//fmt.Println("(*adding", point, "after", index, "interior:", p.interior, "*)")
	index++
	p.points = append(p.points, [2]float64{0, 0})
	copy(p.points[index+1:], p.points[index:])
	p.points[index] = point
	//p.print()
}

func (p *polygon) grow(cans [][2]float64) {

	for i := 0; i < len(p.points); i++ {
		x := len(p.points)
		a, b, c, d := p.points[i%x], p.points[(i+1)%x], p.points[(i+2)%x], p.points[(i+3)%x]

		for _, try := range cans {

			if sameSide(p.interior, try, a, b) && sameSide(p.interior, try, c, d) {
				if !sameSide(p.interior, try, b, c) {
					temp := Make([][2]float64{b, try, c})
					if !temp.containsAny(cans) && !equal(try, b) && !equal(try, c) && !equal(try, d) && !equal(try, a) {

						//fmt.Println("(*", a, b, c, d, "*)")
						p.add(try, (i+1)%x)
						//fmt.Println(p.points)
						i--
						break

					}
				}
			}

		}

	}

}

func triangleArea(a, b, c [2]float64) float64 {
	temp := a[0] * (b[1] - c[1])
	temp += b[0] * (c[1] - a[1])
	temp += c[0] * (a[1] - b[1])
	return .5 * abs(temp)
}

func abs(a float64) float64 {
	if a < 0 {
		return -1 * a
	}
	return a
}

func (p *polygon) area() float64 {
	area := 0.
	for i := 0; i < len(p.points); i++ {
		area += triangleArea(p.points[i], p.points[(i+1)%len(p.points)], p.interior)
	}
	return area

}

func (p *polygon) print() {
	fmt.Print("ListPlot[{")
	for i := 0; i <= len(p.points); i++ {
		x := p.points[i%len(p.points)]
		fmt.Print("{", x[0], ",", x[1], "},")
	}
	fmt.Println("\b},Joined->True]")
}

func (p *polygon) subTriangles() [][3][2]float64 {
	x := len(p.points)
	ans := make([][3][2]float64, 0)
	for i := 0; i < x; i++ {
		for j := 0; j < x; j++ {
			for k := 0; k < x; k++ {
				ans = append(ans, [3][2]float64{p.points[i], p.points[j], p.points[k]})

			}
		}
	}

	return ans
}

//Algorithm does not work in general: implicitly assumes that
//it doesn't matter what order we add points to our hole
func main() {
	starttime := time.Now()

	SMemo = make(map[int]int)
	SMemo[0] = 290797
	points := make([][2]float64, 0)
	for i := 1; i < 2*numpoints; i += 2 {
		points = append(points, [2]float64{float64(T(i)), float64(T(i + 1))})
	}

	max := 0.
	seen := make(map[[3][2]float64]bool)

	for i := 0; i < numpoints; i++ {
		fmt.Println(i)
		for j := i + 1; j < numpoints; j++ {
			for k := j + 1; k < numpoints; k++ {

				if seen[[3][2]float64{points[i], points[j], points[k]}] {
					continue
				}

				triangle := Make([][2]float64{points[i], points[j], points[k]})

				if triangle.containsAny(points) {
					continue
				}

				triangle.grow(points)
				if triangle.area() > max {
					max = triangle.area()
					fmt.Println(max)
					triangle.print()

				}

			}
		}
	}

	fmt.Println(max)
	fmt.Printf("%.1f\n", max)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
