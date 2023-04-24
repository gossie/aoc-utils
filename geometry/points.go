package geometry

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

type Point3d struct {
	X, Y, Z int
}

func CreatePoint3d(x, y, z int) Point3d {
	return Point3d{x, y, z}
}

func (p Point3d) Add(other Point3d) Point3d {
	return CreatePoint3d(p.X+other.X, p.Y+other.Y, p.Z+other.Z)
}

func (p Point3d) Multiply(n int) Point3d {
	return CreatePoint3d(p.X*n, p.Y*n, p.Z*n)
}
