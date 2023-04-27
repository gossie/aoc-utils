package geometry

import (
	"fmt"
	"math"
)

type point2d struct {
	x, y int
}

func CreatePoint2d(x, y int) point2d {
	return point2d{x, y}
}

func (p point2d) X() int {
	return p.x
}

func (p point2d) Y() int {
	return p.y
}

func (p point2d) Add(other point2d) point2d {
	return CreatePoint2d(p.x+other.x, p.y+other.y)
}

func (p point2d) Multiply(n int) point2d {
	return CreatePoint2d(p.x*n, p.y*n)
}

func (p point2d) AirDistanceTo(other point2d) float64 {
	xDistance := math.Abs(float64(other.x - p.x))
	yDistance := math.Abs(float64(other.y - p.y))
	return math.Sqrt(math.Pow(xDistance, 2) + math.Pow(yDistance, 2))
}

func (p point2d) ManhattenDistanceTo(other point2d) int {
	xDistance := int(math.Abs(float64(other.x - p.x)))
	yDistance := int(math.Abs(float64(other.y - p.y)))
	return xDistance + yDistance
}

func (p point2d) String() string {
	return fmt.Sprintf("x: %d, y: %d", p.x, p.y)
}
