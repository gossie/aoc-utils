package geometry

import (
	"fmt"

	"github.com/gossie/aoc-utils/equations"
)

type Line2d struct {
	start, direction Point2d
}

func NewLine2d(p1, p2 Point2d) Line2d {
	return Line2d{p1, p2.Subtract(p1)}
}

func (l Line2d) Contains(point Point2d) bool {
	left := point.Subtract(l.start)
	return left.x/l.direction.x == left.y/l.direction.y
}

func (l Line2d) Point(factor float64) Point2d {
	return NewPoint2d(l.start.x+factor*l.direction.x, l.start.y+factor*l.direction.y)
}

func (l Line2d) IntersectsAt(other Line2d) (Point2d, error) {
	xLeft := equations.Add(equations.Num(l.start.x), equations.Var(l.direction.x, "l"))
	xRight := equations.Add(equations.Num(other.start.x), equations.Var(other.direction.x, "r"))
	firstEq := equations.NewEquation(xLeft, xRight)

	yLeft := equations.Add(equations.Num(l.start.y), equations.Var(l.direction.y, "l"))
	yRight := equations.Add(equations.Num(other.start.y), equations.Var(other.direction.y, "r"))
	secondEq := equations.NewEquation(yLeft, yRight)

	// fmt.Println("firstEq ", firstEq)
	// fmt.Println("secondEq", secondEq)

	lValue, _ := firstEq.SolveTo("l")
	// fmt.Println("lValue ", lValue)
	secondWithL := secondEq.Set("l", *lValue)
	// fmt.Println("secondWithL ", secondWithL)
	rValue, _ := secondWithL.SolveTo("r")
	// fmt.Println("rValue", rValue)

	//p1 := l.Point(lValue.Number())
	p2 := other.Point(rValue.Number())
	// if p1 != p2 {
	// 	return Point2d{}, errors.New("no intersection")
	// }

	return p2, nil
}

func (l Line2d) String() string {
	return fmt.Sprintf("g: x = (%f, %f) + r (%f, %f)", l.start.x, l.start.y, l.direction.x, l.direction.y)
}
