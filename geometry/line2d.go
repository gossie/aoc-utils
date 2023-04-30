package geometry

import (
	"fmt"
)

type Line2d struct {
	start, direction Point2d
}

func NewLine2d(p1, p2 Point2d) Line2d {
	return Line2d{p1, p2.Subtract(p1)}
}

func (l Line2d) Contains(point Point2d) bool {
	left := point.Subtract(l.start)
	return float64(left.x)/float64(l.direction.x) == float64(left.y)/float64(l.direction.y)
}

func (l Line2d) IntersectsAt(other Line2d) (Point2d, error) {
	// xLeft := equations.Add(equations.Num(l.start.x), equations.Mul(equations.Var("l"), equations.Num(l.direction.x)))
	// xRight := equations.Add(equations.Num(other.start.x), equations.Mul(equations.Var("r"), equations.Num(other.direction.x)))
	// firstEq := equations.NewEquation(xLeft, xRight)

	// yLeft := equations.Add(equations.Num(l.start.y), equations.Mul(equations.Var("l"), equations.Num(l.direction.y)))
	// yRight := equations.Add(equations.Num(other.start.y), equations.Mul(equations.Var("r"), equations.Num(other.direction.y)))
	// secondEq := equations.NewEquation(yLeft, yRight)

	// lValue, _ := firstEq.SolveTo("l")
	// rValue, _ := secondEq.Set("l", *lValue).SolveTo("r")

	// // TODO: check

	return Point2d{}, nil
}

func (l Line2d) String() string {
	return fmt.Sprintf("g: x = (%f, %f) + r (%f, %f)", l.start.x, l.start.y, l.direction.x, l.direction.y)
}
