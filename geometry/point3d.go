package geometry

import (
	"fmt"
	"math"
)

type Point3d struct {
	x, y, z float64
}

func NewPoint3d(x, y, z float64) Point3d {
	return Point3d{x, y, z}
}

func (p Point3d) X() float64 {
	return p.x
}

func (p Point3d) Y() float64 {
	return p.y
}

func (p Point3d) Z() float64 {
	return p.z
}

func (p Point3d) Add(other Point3d) Point3d {
	return NewPoint3d(p.x+other.x, p.y+other.y, p.z+other.z)
}

func (p Point3d) Subtract(other Point3d) Point3d {
	return NewPoint3d(p.x-other.x, p.y-other.y, p.z-other.z)
}

func (p Point3d) Multiply(n float64) Point3d {
	return NewPoint3d(p.x*n, p.y*n, p.z*n)
}

func (p Point3d) AirDistanceTo(other Point3d) float64 {
	return calcDistance3d(p, other, func(d1, d2, d3 float64) float64 {
		return math.Sqrt(math.Pow(d1, 2) + math.Pow(d2, 2) + math.Pow(d3, 2))
	})
}

func (p Point3d) ManhattenDistanceTo(other Point3d) uint {
	return calcDistance3d(p, other, func(d1, d2, d3 float64) uint {
		return uint(d1 + d2 + d3)
	})
}

func calcDistance3d[T any](p1, p2 Point3d, result func(d1, d2, d3 float64) T) T {
	xDistance := math.Abs(float64(p1.x - p2.x))
	yDistance := math.Abs(float64(p1.y - p2.y))
	zDistance := math.Abs(float64(p1.z - p2.z))
	return result(xDistance, yDistance, zDistance)
}

func (p Point3d) String() string {
	return fmt.Sprintf("x: %f, y: %f, z: %f", p.x, p.y, p.z)
}
