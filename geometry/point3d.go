package geometry

import "math"

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

func (p point3d) DistanceTo(other point3d) float64 {
	xDistance := math.Abs(float64(other.x - p.x))
	yDistance := math.Abs(float64(other.y - p.y))
	zDistance := math.Abs(float64(other.z - p.z))
	return math.Sqrt(math.Pow(xDistance, 2) + math.Pow(yDistance, 2) + math.Pow(zDistance, 2))
}
