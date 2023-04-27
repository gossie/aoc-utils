package geometry

import (
	"math"
)

type point3d struct {
	x, y, z int
}

func CreatePoint3d(x, y, z int) point3d {
	return point3d{x, y, z}
}

func (p point3d) X() int {
	return p.x
}

func (p point3d) Y() int {
	return p.y
}

func (p point3d) Z() int {
	return p.z
}

func (p point3d) Add(other point3d) point3d {
	return CreatePoint3d(p.x+other.x, p.y+other.y, p.z+other.z)
}

func (p point3d) Multiply(n int) point3d {
	return CreatePoint3d(p.x*n, p.y*n, p.z*n)
}

func (p point3d) AirDistanceTo(other point3d) float64 {
	return calcDistance(p, other, func(d1, d2, d3 float64) float64 {
		return math.Sqrt(math.Pow(d1, 2) + math.Pow(d2, 2) + math.Pow(d3, 2))
	})
}

func (p point3d) ManhattenDistanceTo(other point3d) uint {
	return calcDistance(p, other, func(d1, d2, d3 float64) uint {
		return uint(d1 + d2 + d3)
	})
}

func calcDistance[T any](p1, p2 point3d, result func(d1, d2, d3 float64) T) T {
	xDistance := math.Abs(float64(p1.x - p2.x))
	yDistance := math.Abs(float64(p1.y - p2.y))
	zDistance := math.Abs(float64(p1.z - p2.z))
	return result(xDistance, yDistance, zDistance)
}
