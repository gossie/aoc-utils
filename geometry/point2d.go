package geometry

import "math"

type Point2d struct {
	X, Y int
}

func CreatePoint2d(x, y int) Point2d {
	return Point2d{x, y}
}

func (p Point2d) Add(other Point2d) Point2d {
	return CreatePoint2d(p.X+other.X, p.Y+other.Y)
}

func (p Point2d) Multiply(n int) Point2d {
	return CreatePoint2d(p.X*n, p.Y*n)
}

func (p Point2d) DistanceTo(other Point2d) float64 {
	xDistance := math.Abs(float64(other.X - p.X))
	yDistance := math.Abs(float64(other.Y - p.Y))
	return math.Sqrt(math.Pow(xDistance, 2) + math.Pow(yDistance, 2))
}
