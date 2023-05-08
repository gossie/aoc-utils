package geometry

type Line3d struct {
	start, direction Point3d
}

func NewLine3d(p1, p2 Point3d) Line3d {
	return Line3d{p1, p2.Subtract(p1)}
}

func (l Line3d) Contains(point Point3d) bool {
	left := point.Subtract(l.start)

	r1 := 0.0
	if l.direction.x != 0 {
		r1 = left.x / l.direction.x
	}
	r2 := 0.0
	if l.direction.y != 0 {
		r2 = left.y / l.direction.y
	}
	r3 := 0.0
	if l.direction.z != 0 {
		r3 = left.z / l.direction.z
	}

	return (r1 > 0 && ((r2 > 0 && r1 == r2) || r2 == 0) && ((r3 > 0 && r1 == r3) || r3 == 0)) || (r1 == 0 && left.x == 0) &&
		(r2 > 0 && ((r1 > 0 && r2 == r1) || r1 == 0) && ((r3 > 0 && r2 == r3) || r3 == 0)) || (r2 == 0 && left.y == 0) &&
		(r3 > 0 && ((r2 > 0 && r3 == r2) || r2 == 0) && ((r1 > 0 && r3 == r1) || r1 == 0)) || (r3 == 0 && left.z == 0)
}
