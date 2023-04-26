package geometry

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
