//Plane geometry support for Project Euler
package plane

import (
	"errors"
	"math"
)

type Circle struct {
	Center Point
	Radius float64
}

type Point struct {
	X, Y float64
}

type Triangle struct {
	A, B, C Point //Vertices
}

type Line struct {
	Slope, Intercept float64
	Vertical         bool
}

func (b *Point) Distance(a *Point) float64 {
	dx := b.X - a.X
	dy := b.Y - a.Y
	return math.Sqrt(dx*dx + dy*dy)
}

//LineFromPoints gives the line through the given points
func LineFromPoints(a, b *Point) (*Line, error) {
	if a.X == b.X && a.Y == b.Y {
		return new(Line), errors.New("Can't make line: points are the same")
	}

	if a.X == b.X {
		return &Line{0, a.Y, true}, nil
	}

	dy := b.Y - a.Y
	dx := b.X - a.X

	slope := dy / dx

	inter := a.Y - slope*a.X

	return &Line{slope, inter, false}, nil

}

//RightOf returns true if the given point is to the right of the line
//(note: More accurately, point is "above" the line. Rename to that affect?)
func (pt *Point) RightOf(line *Line) bool {

	if line.Evaluate(pt.X) < pt.Y {
		return true
	}

	return false
}

//LineFromPtSlope returns a line object having the provided point and slope.
//(note: maybe rename to NewLine or NewLineMXB)
func LineFromPtSlope(slope float64, x *Point) *Line {
	b := x.Y - slope*x.X
	return &Line{slope, b, false}
}

//IntersectLine returns the point where the given lines intersect. If they don't
//intersect, it returns an error (note: vertical line support is not yet implemented)
func (l *Line) IntersectLine(m *Line) (intersection *Point, err error) {
	if l.Slope == m.Slope {
		if l.Intercept == m.Intercept {
			return &Point{0, l.Intercept}, errors.New("Lines Coincide")
		} else {
			return new(Point), errors.New("Can't find intersection: lines don't intersect")
		}
	}

	if m.Vertical == true || l.Vertical == true {
		return new(Point), errors.New("Vertical Lines!!")
	}

	x := (m.Intercept - l.Intercept) / (l.Slope - m.Slope)
	y := (m.Intercept*l.Slope - l.Intercept*m.Slope) / (l.Slope - m.Slope)

	return &Point{x, y}, nil
}

//CircleFromPoints returns the unique circle through the provided points,
//provided such exists.
func CircleFromPoints(a, b, c *Point) (*Circle, error) {
	tri := MakeTriange(a, b, c)
	trarea := tri.Area()

	if trarea == 0 {
		return new(Circle), errors.New("Can't construct circle: points colinear")
	}

	l1, err1 := Bisect(a, b)
	l2, err2 := Bisect(b, c)

	if err1 != nil || err2 != nil {
		return new(Circle), errors.New("Can't bisect ab or bc")
	}

	center, err := l1.IntersectLine(l2)

	if err != nil {
		return new(Circle), err
	}

	radius := center.Distance(a)

	return &Circle{Point{center.X, center.Y}, radius}, nil
}

//Bisect returns the line which is the perpendicular bisector to the segment
//through ponts x, and y.
func Bisect(x, y *Point) (*Line, error) {
	through, err := LineFromPoints(x, y)
	if err != nil {
		return new(Line), err
	}

	if through.Slope == 0 {
		return &Line{0, x.Midpoint(y).X, true}, nil
	}

	slope := -1 / through.Slope

	return LineFromPtSlope(slope, x.Midpoint(y)), nil

}

func (a *Point) Midpoint(b *Point) *Point {
	return &Point{.5 * (a.X + b.X), .5 * (a.Y + b.Y)}
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (l *Line) Evaluate(x float64) float64 {
	return l.Slope*x + l.Intercept
}

func MakeTriange(a, b, c *Point) *Triangle {
	A := Point{a.X, a.Y}
	B := Point{b.X, b.Y}
	C := Point{c.X, c.Y}

	return &Triangle{A, B, C}
}

func (t *Triangle) Area() float64 {
	return .5 * math.Abs(t.A.X*(t.B.Y-t.C.Y)+t.B.X*(t.C.Y-t.A.Y)+t.C.X*(t.A.Y-t.B.Y))
}
