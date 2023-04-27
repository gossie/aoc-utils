package geometry

import (
	"fmt"
	"math"
)

type Point2d struct {
	x, y int
}

func CreatePoint2d(x, y int) Point2d {
	return Point2d{x, y}
}

func (p Point2d) X() int {
	return p.x
}

func (p Point2d) Y() int {
	return p.y
}

func (p Point2d) Add(other Point2d) Point2d {
	return CreatePoint2d(p.x+other.x, p.y+other.y)
}

func (p Point2d) Multiply(n int) Point2d {
	return CreatePoint2d(p.x*n, p.y*n)
}

func (p Point2d) AirDistanceTo(other Point2d) float64 {
	return calcDistance2d(p, other, func(d1, d2 float64) float64 {
		return math.Sqrt(math.Pow(d1, 2) + math.Pow(d2, 2))
	})
}

func (p Point2d) ManhattenDistanceTo(other Point2d) uint {
	return calcDistance2d(p, other, func(d1, d2 float64) uint {
		return uint(d1 + d2)
	})
}

func calcDistance2d[T any](p1, p2 Point2d, result func(d1, d2 float64) T) T {
	xDistance := math.Abs(float64(p1.x - p2.x))
	yDistance := math.Abs(float64(p1.y - p2.y))
	return result(xDistance, yDistance)
}

func (p Point2d) String() string {
	return fmt.Sprintf("x: %d, y: %d", p.x, p.y)
}
