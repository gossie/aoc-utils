package geometry

import (
	"errors"
	"fmt"

	"github.com/gossie/equations"
)

type Line2d struct {
	start, direction Point2d
}

func NewLine2d(p1, p2 Point2d) Line2d {
	return Line2d{p1, p2.Subtract(p1)}
}

func (l Line2d) Contains(point Point2d) bool {
	left := point.Subtract(l.start)

	r1 := 0.0
	if l.direction.x != 0 {
		r1 = left.x / l.direction.x
	}
	r2 := 0.0
	if l.direction.y != 0 {
		r2 = left.y / l.direction.y
	}

	return (r1 > 0 && ((r2 > 0 && r1 == r2) || r2 == 0)) || (r1 == 0 && left.x == 0) &&
		(r2 > 0 && ((r1 > 0 && r2 == r1) || r1 == 0)) || (r2 == 0 && left.y == 0)
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

	lValue, _ := equations.SolveTo(&firstEq, "l")
	secondWithL := equations.Set(&secondEq, "l", *lValue)
	rValue, err := equations.SolveTo(&secondWithL, "r")
	if err != nil {
		err, ok := err.(*equations.SolveError)
		if ok {
			if err.FinalEquation.IsTrue() {
				return Point2d{}, errors.New("same line")
			} else {
				return Point2d{}, errors.New("no intersection")
			}
		}
		return Point2d{}, err

	}

	return other.Point(rValue.Number()), nil
}

func (l Line2d) String() string {
	return fmt.Sprintf("g: x = (%f, %f) + r (%f, %f)", l.start.x, l.start.y, l.direction.x, l.direction.y)
}
